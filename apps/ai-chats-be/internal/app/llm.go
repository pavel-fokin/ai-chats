package app

import (
	"context"
	"fmt"

	"ai-chats/internal/app/notifications"
	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

// ModelStreamResponse is a message that represents streamed response of a model.
type ModelStreamResponse struct {
	Text   string `json:"text"`
	Sender string `json:"sender"`
}

func (m ModelStreamResponse) Type() types.MessageType {
	return types.MessageType("ModelStreamResponse")
}

type LLM struct {
	chats        domain.Chats
	ollamaClient domain.OllamaClient
	pubsub       PubSub
	tx           Tx
	notificator  Notificator
}

func NewLLM(
	chats domain.Chats,
	ollamaClient domain.OllamaClient,
	pubsub PubSub,
	tx Tx,
	notificator Notificator,
) *LLM {
	return &LLM{chats: chats, ollamaClient: ollamaClient, pubsub: pubsub, tx: tx, notificator: notificator}
}

// GenerateResponse generates a LLM response for the chat.
func (l *LLM) GenerateResponse(ctx context.Context, chatID domain.ChatID) error {
	chat, err := l.chats.FindByIDWithMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	model, err := l.ollamaClient.NewModel(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	chatResponseFunc := func(modelStreamMessage domain.ModelStreamMessage) error {
		if err := l.notificator.Notify(ctx, notifications.NewModelStreamMessage(
			chatID,
			modelStreamMessage.Text,
			modelStreamMessage.Sender.Format(),
		)); err != nil {
			return fmt.Errorf("failed to notify in chat: %w", err)
		}
		return nil
	}

	llmMessage, err := model.Chat(ctx, chat.Messages, chatResponseFunc)
	if err != nil {
		return fmt.Errorf("failed to generate a response: %w", err)
	}

	chat.AddMessage(llmMessage)
	if err := l.chats.Update(ctx, chat); err != nil {
		return fmt.Errorf("error adding a message to chat %s: %w", chatID, err)
	}

	messageAdded := domain.NewMessageAdded(chatID, llmMessage)
	if err := l.pubsub.Publish(ctx, MessageAddedTopic, messageAdded); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	return nil
}

// GenerateChatTitle generates a chat title.
func (l *LLM) GenerateChatTitleAsync(ctx context.Context, chatID domain.ChatID) error {
	generateChatTitleCommand := NewGenerateChatTitle(chatID.String())
	if err := l.pubsub.Publish(
		ctx,
		GenerateChatTitleTopic,
		generateChatTitleCommand,
	); err != nil {
		return fmt.Errorf("failed to publish a generate chat title command: %w", err)
	}

	return nil
}

// GenerateTitle generates a LLM title for the chat.
func (l *LLM) GenerateTitle(ctx context.Context, chatID domain.ChatID) error {
	chat, err := l.chats.FindByIDWithMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("error finding chat %s: %w", chatID, err)
	}

	model, err := l.ollamaClient.NewModel(chat.DefaultModel.AsOllamaModel())
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

	generatedTitle, err := model.Chat(
		ctx,
		messages,
		func(message domain.ModelStreamMessage) error {
			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("error generating title for chat %s: %w", chatID, err)
	}

	if err := l.tx.Tx(ctx, func(ctx context.Context) error {
		chat.UpdateTitle(generatedTitle.Text)
		if err := l.chats.Update(ctx, chat); err != nil {
			return fmt.Errorf("error updating title for chat %s: %w", chatID, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("error in transaction while updating title for chat %s: %w", chatID, err)
	}

	chatTitleUpdated := notifications.NewChatTitleUpdated(chatID, chat.User.ID)
	if err := l.notificator.Notify(ctx, chatTitleUpdated); err != nil {
		return fmt.Errorf("error notifying app about title update for chat %s: %w", chatID, err)
	}

	return nil
}

// ProcessAddedMessage processes a message added event.
func (l *LLM) ProcessAddedMessage(ctx context.Context, event domain.MessageAdded) error {
	messageAdded := notifications.NewMessageAdded(event.ChatID)
	if err := l.notificator.Notify(ctx, messageAdded); err != nil {
		return fmt.Errorf("failed to notify in chat: %w", err)
	}

	switch {
	case event.Message.IsFromUser():
		l.GenerateResponse(ctx, event.ChatID)
	case event.Message.IsFromModel():
		// Ignore messages from models.
	default:
		return fmt.Errorf("unknown message type: %s", event.Message)
	}

	chat, err := l.chats.FindByIDWithMessages(ctx, event.ChatID)
	if err != nil {
		return fmt.Errorf("error finding chat: %w", err)
	}

	if len(chat.Messages) == 2 {
		return l.GenerateTitle(ctx, event.ChatID)
	}

	return nil
}
