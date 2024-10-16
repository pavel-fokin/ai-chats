package app

import (
	"context"
	"fmt"

	"ai-chats/internal/app/commands"
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

	streamFunc := func(message domain.Message) error {
		if err := a.notifyChat(ctx, chatID.String(), ChatMessage{Message: message}); err != nil {
			return fmt.Errorf("failed to notify in chat: %w", err)
		}
		return nil
	}

	llmMessage, err := model.Chat(ctx, chat.Messages, streamFunc)
	if err != nil {
		return fmt.Errorf("failed to generate a response: %w", err)
	}

	if err := a.chats.AddMessage(ctx, chatID, llmMessage); err != nil {
		return fmt.Errorf("failed to add a message: %w", err)
	}

	messageAdded := domain.NewMessageAdded(chatID, llmMessage)
	if err := a.pubsub.Publish(ctx, MessageAddedTopic, messageAdded); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	return nil
}

// GenerateChatTitle generates a chat title.
func (a *App) GenerateChatTitleAsync(ctx context.Context, chatID domain.ChatID) error {
	generateChatTitleCommand := commands.GenerateChatTitle{ChatID: chatID.String()}
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
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	model, err := a.ollamaClient.NewModel(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	messages := make([]domain.Message, 0, len(chat.Messages))
	messages = append(messages, chat.Messages...)
	messages = append(
		messages,
		domain.NewSystemMessage(
			`Provide a one-sentence, short title of this following conversation.
Use less than 100 characters. Don't use quotes or special characters.`,
		),
	)

	generatedTitle, err := model.Chat(ctx, messages, func(message domain.Message) error {
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to generate a title: %w", err)
	}

	if err := a.chats.UpdateTitle(ctx, chatID, generatedTitle.Text); err != nil {
		return fmt.Errorf("failed to update chat title: %w", err)
	}

	titleUpdated := domain.NewChatTitleUpdated(chatID, generatedTitle.Text)
	if err := a.notifyApp(ctx, chat.User.ID, titleUpdated); err != nil {
		return fmt.Errorf("failed to publish a title updated event: %w", err)
	}

	return nil
}
