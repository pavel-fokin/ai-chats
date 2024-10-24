package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

type ChatMessage struct {
	domain.Message
}

func (m ChatMessage) Type() types.MessageType {
	return types.MessageType("chatMessage")
}

// GenerateResponse generates a LLM response for the chat.
func (a *App) GenerateResponse(ctx context.Context, chatID domain.ChatID) error {
	chat, err := a.chats.FindByIDWithMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	model, err := a.ollamaClient.NewModel(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	chatResponseFunc := func(message domain.Message) error {
		if err := a.notifyChat(ctx, chatID.String(), ChatMessage{Message: message}); err != nil {
			return fmt.Errorf("failed to notify in chat: %w", err)
		}
		return nil
	}

	llmMessage, err := model.Chat(ctx, chat.Messages, chatResponseFunc)
	if err != nil {
		return fmt.Errorf("failed to generate a response: %w", err)
	}

	chat.AddMessage(llmMessage)
	if err := a.chats.Update(ctx, chat); err != nil {
		return fmt.Errorf("error adding a message to chat %s: %w", chatID, err)
	}

	messageAdded := domain.NewMessageAdded(chatID, llmMessage)
	if err := a.pubsub.Publish(ctx, MessageAddedTopic, messageAdded); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	return nil
}

// GenerateChatTitle generates a chat title.
func (a *App) GenerateChatTitleAsync(ctx context.Context, chatID domain.ChatID) error {
	generateChatTitleCommand := NewGenerateChatTitle(chatID.String())
	if err := a.pubsub.Publish(
		ctx,
		GenerateChatTitleTopic,
		generateChatTitleCommand,
	); err != nil {
		return fmt.Errorf("failed to publish a generate chat title command: %w", err)
	}

	return nil
}

// GenerateTitle generates a LLM title for the chat.
func (a *App) GenerateTitle(ctx context.Context, chatID domain.ChatID) error {
	chat, err := a.chats.FindByIDWithMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("error finding chat %s: %w", chatID, err)
	}

	model, err := a.ollamaClient.NewModel(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("error creating a chat model for chat %s: %w", chatID, err)
	}

	messages := append(
		chat.Messages,
		domain.NewSystemMessage(
			`Provide a one-sentence, short title of this conversation.
Use less than 100 characters. Don't use quotes or special characters.`,
		),
	)

	generatedTitle, err := model.Chat(ctx, messages, func(message domain.Message) error {
		return nil
	})
	if err != nil {
		return fmt.Errorf("error generating title for chat %s: %w", chatID, err)
	}

	if err := a.tx.Tx(ctx, func(ctx context.Context) error {
		chat.UpdateTitle(generatedTitle.Text)
		if err := a.chats.Update(ctx, chat); err != nil {
			return fmt.Errorf("error updating title for chat %s: %w", chatID, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("error in transaction while updating title for chat %s: %w", chatID, err)
	}

	titleUpdated := domain.NewChatTitleUpdated(chatID, generatedTitle.Text)
	if err := a.notifyApp(ctx, chat.User.ID, titleUpdated); err != nil {
		return fmt.Errorf("error notifying app about title update for chat %s: %w", chatID, err)
	}

	return nil
}
