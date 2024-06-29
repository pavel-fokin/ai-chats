package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/commands"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/json"
)

// CreateChat creates a chat for the user.
func (a *App) CreateChat(ctx context.Context, userID uuid.UUID, text string) (domain.Chat, error) {
	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find a user: %w", err)
	}

	chat := domain.NewChat(user)
	if err := a.tx.Tx(ctx, func(ctx context.Context) error {
		if err := a.chats.Add(ctx, chat); err != nil {
			return fmt.Errorf("failed to add a chat: %w", err)
		}

		// If the message is empty, we don't need to send it.
		if text == "" {
			return nil
		}

		message := domain.NewMessage("User", text)

		if err := a.messages.Add(ctx, chat.ID, message); err != nil {
			return fmt.Errorf("failed to add a message: %w", err)
		}

		generateResponse := commands.NewGenerateResponse(chat.ID)
		if err := a.pubsub.Publish(ctx, "worker", json.MustMarshal(ctx, generateResponse)); err != nil {
			return fmt.Errorf("failed to publish a generate response command: %w", err)
		}

		return nil
	}); err != nil {
		return domain.Chat{}, fmt.Errorf("failed to create a chat: %w", err)
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

	messageSent := events.NewMessageAdded(chatID, message)
	if err := a.pubsub.Publish(ctx, chatID.String(), json.MustMarshal(ctx, messageSent)); err != nil {
		return domain.Message{}, fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	generateResponse := commands.NewGenerateResponse(chatID)
	if err := a.pubsub.Publish(ctx, "worker", json.MustMarshal(ctx, generateResponse)); err != nil {
		return domain.Message{}, fmt.Errorf("failed to publish a generate response command: %w", err)
	}

	return domain.Message{}, nil
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