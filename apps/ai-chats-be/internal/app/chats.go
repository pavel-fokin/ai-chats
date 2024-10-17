package app

import (
	"context"
	"fmt"
	"strings"

	"ai-chats/internal/domain"
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
func (a *App) CreateChat(ctx context.Context, userID domain.UserID, model, messageText string) (domain.Chat, error) {
	messageText = strings.TrimSpace(messageText)
	if messageText == "" {
		return domain.Chat{}, fmt.Errorf("message text is empty")
	}

	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("error finding user: %w", err)
	}

	var chat domain.Chat
	if err := a.tx.Tx(ctx, func(ctx context.Context) error {
		chat = domain.NewChat(user, domain.NewModelID(model))
		chat.AddMessage(domain.NewUserMessage(user, messageText))

		if err := a.chats.Add(ctx, chat); err != nil {
			return fmt.Errorf("error adding a chat: %w", err)
		}

		return nil
	}); err != nil {
		return domain.Chat{}, fmt.Errorf("error creating chat: %w", err)
	}

	for _, event := range chat.Events {
		if err := a.pubsub.Publish(ctx, MessageAddedTopic, event); err != nil {
			return domain.Chat{}, fmt.Errorf("error publishing chat events: %w", err)
		}
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
func (a *App) ProcessAddedMessage(ctx context.Context, event domain.MessageAdded) error {
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
	ctx context.Context, userID domain.UserID, chatID domain.ChatID, messageText string,
) (domain.Message, error) {
	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("failed to find a user: %w", err)
	}

	var (
		chat    domain.Chat
		message domain.Message
	)
	if err := a.tx.Tx(ctx, func(ctx context.Context) error {
		chat, err = a.chats.FindByID(ctx, chatID)
		if err != nil {
			return fmt.Errorf("error finding chat: %w", err)
		}

		chat.AddMessage(domain.NewUserMessage(user, messageText))
		if err := a.chats.Update(ctx, chat); err != nil {
			return fmt.Errorf("error updating chat: %w", err)
		}

		return nil
	}); err != nil {
		return domain.Message{}, fmt.Errorf("error sending message: %w", err)
	}

	for _, event := range chat.Events {
		if err := a.pubsub.Publish(ctx, MessageAddedTopic, event); err != nil {
			return domain.Message{}, fmt.Errorf("error publishing events: %w", err)
		}
	}

	return message, nil
}
