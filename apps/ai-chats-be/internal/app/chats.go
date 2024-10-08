package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
)

// AllChats returns all chats for the user.
func (a *App) AllChats(ctx context.Context, userID domain.UserID) ([]domain.Chat, error) {
	return a.chats.AllChats(ctx, userID)
}

// AllMessages returns all messages in the chat.
func (a *App) AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error) {
	return a.chats.AllMessages(ctx, chatID)
}

// ChatExists checks if the chat exists.
func (a *App) ChatExists(ctx context.Context, chatID domain.ChatID) (bool, error) {
	return a.chats.Exists(ctx, chatID)
}

// CreateChat creates a chat for the user.
func (a *App) CreateChat(ctx context.Context, userID domain.UserID, model, text string) (domain.Chat, error) {
	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find a user: %w", err)
	}

	var message domain.Message
	chat := domain.NewChat(user, domain.NewModelID(model))
	if err := a.tx.Tx(ctx, func(ctx context.Context) error {
		if err := a.chats.Add(ctx, chat); err != nil {
			return fmt.Errorf("failed to add a chat: %w", err)
		}

		// If the message is empty, we don't need to send it.
		if text == "" {
			return nil
		}

		message = domain.NewUserMessage(user, text)
		if err := a.chats.AddMessage(ctx, chat.ID, message); err != nil {
			return fmt.Errorf("failed to add a message: %w", err)
		}

		return nil
	}); err != nil {
		return domain.Chat{}, fmt.Errorf("failed to create a chat: %w", err)
	}

	// If the message is empty, we don't need to send it.
	if text == "" {
		return chat, nil
	}

	messageAdded := events.NewMessageAdded(chat.ID, message)
	if err := a.pubsub.Publish(ctx, MessageAddedTopic, messageAdded); err != nil {
		return domain.Chat{}, fmt.Errorf("failed to publish a message added event: %w", err)
	}

	return chat, nil
}

// FindChatByID finds a chat by ID.
func (a *App) FindChatByID(ctx context.Context, chatID domain.ChatID) (domain.Chat, error) {
	return a.chats.FindByID(ctx, chatID)
}

// DeleteChat deletes the chat.
func (a *App) DeleteChat(ctx context.Context, chatID domain.ChatID) error {
	return a.chats.Delete(ctx, chatID)
}

// ProcessAddedMessage processes a message added event.
func (a *App) ProcessAddedMessage(ctx context.Context, event events.MessageAdded) error {
	if err := a.notifyChat(ctx, event.ChatID.String(), event); err != nil {
		return fmt.Errorf("failed to notify in chat: %w", err)
	}

	switch {
	case event.Message.IsFromUser():
		a.GenerateResponse(ctx, event.ChatID)
	case event.Message.IsFromModel():
		// Ignore messages from models.
	default:
		return fmt.Errorf("unknown message type: %s", event.Message)
	}

	messages, err := a.chats.AllMessages(ctx, event.ChatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	if len(messages) == 2 {
		return a.GenerateTitle(ctx, event.ChatID)
	}

	return nil
}

// SendMessage sends a message to the chat.
func (a *App) SendMessage(
	ctx context.Context, userID domain.UserID, chatID domain.ChatID, text string,
) (domain.Message, error) {
	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("failed to find a user: %w", err)
	}

	message := domain.NewUserMessage(user, text)

	if err := a.chats.AddMessage(ctx, chatID, message); err != nil {
		return domain.Message{}, fmt.Errorf("failed to add a message: %w", err)
	}

	messageAdded := events.NewMessageAdded(chatID, message)
	if err := a.pubsub.Publish(ctx, MessageAddedTopic, messageAdded); err != nil {
		return domain.Message{}, fmt.Errorf("failed to publish a message added event: %w", err)
	}

	return message, nil
}
