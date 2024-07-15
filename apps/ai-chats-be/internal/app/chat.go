package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/json"
	"pavel-fokin/ai/apps/ai-bots-be/internal/worker"
)

// CreateChat creates a chat for the user.
func (a *App) CreateChat(ctx context.Context, userID uuid.UUID, model, text string) (domain.Chat, error) {
	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find a user: %w", err)
	}

	var message domain.Message
	chat := domain.NewChat(user, model)
	if err := a.tx.Tx(ctx, func(ctx context.Context) error {
		if err := a.chats.Add(ctx, chat); err != nil {
			return fmt.Errorf("failed to add a chat: %w", err)
		}

		// If the message is empty, we don't need to send it.
		if text == "" {
			return nil
		}

		message = domain.NewMessage("User", text)
		if err := a.messages.Add(ctx, chat.ID, message); err != nil {
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
	if err := a.pubsub.Publish(ctx, worker.MessageAddedTopic, json.MustMarshal(ctx, messageAdded)); err != nil {
		return domain.Chat{}, fmt.Errorf("failed to publish a message added event: %w", err)
	}

	return chat, nil
}

// DeleteChat deletes the chat.
func (a *App) DeleteChat(ctx context.Context, chatID domain.ChatID) error {
	return a.chats.Delete(ctx, chatID)
}

// AllChats returns all chats for the user.
func (a *App) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	return a.chats.AllChats(ctx, userID)
}

// SendMessage sends a message to the chat.
func (a *App) SendMessage(ctx context.Context, chatID domain.ChatID, text string) (domain.Message, error) {
	message := domain.NewMessage("User", text)

	if err := a.messages.Add(ctx, chatID, message); err != nil {
		return domain.Message{}, fmt.Errorf("failed to add a message: %w", err)
	}

	messageAdded := events.NewMessageAdded(chatID, message)
	if err := a.pubsub.Publish(ctx, worker.MessageAddedTopic, json.MustMarshal(ctx, messageAdded)); err != nil {
		return domain.Message{}, fmt.Errorf("failed to publish a message added event: %w", err)
	}

	return message, nil
}

// AllMessages returns all messages in the chat.
func (a *App) AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error) {
	return a.messages.AllMessages(ctx, chatID)
}

// ChatExists checks if the chat exists.
func (a *App) ChatExists(ctx context.Context, chatID domain.ChatID) (bool, error) {
	return a.chats.Exists(ctx, chatID)
}

// FindChatByID finds a chat by ID.
func (a *App) FindChatByID(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	return a.chats.FindByID(ctx, chatID)
}
