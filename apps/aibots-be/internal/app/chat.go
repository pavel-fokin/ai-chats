package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/llm"

	"github.com/google/uuid"
)

// CreateChat creates a chat for the user.
func (a *App) CreateChat(ctx context.Context, userID uuid.UUID) (domain.Chat, error) {
	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find a user: %w", err)
	}

	chat := domain.NewChat(user)

	if err := a.chats.Add(ctx, chat); err != nil {
		return domain.Chat{}, fmt.Errorf("failed to add a chat: %w", err)
	}

	return chat, nil
}

// AllChats returns all chats for the user.
func (a *App) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	return a.chats.AllChats(ctx, userID)
}

// SendMessage sends a message to the chat.
func (a *App) SendMessage(ctx context.Context, chatID uuid.UUID, text string) (domain.Message, error) {
	message := domain.NewMessage("User", text)

	_, err := a.chatting.SendMessage(ctx, chatID, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("failed to send a message: %w", err)
	}

	messageSent := domain.NewMessageSent(chatID, message)
	if err := a.events.Publish(ctx, chatID.String(), messageSent.AsBytes()); err != nil {
		return domain.Message{}, fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	if err := a.events.Publish(ctx, "worker", messageSent.AsBytes()); err != nil {
		return domain.Message{}, fmt.Errorf("failed to generate respone event: %w", err)
	}

	return domain.Message{}, nil
}

// GenerateResponse generates a LLM response for the chat.
func (a *App) GenerateResponse(ctx context.Context, chatID uuid.UUID) error {
	messages, err := a.AllMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to get messages: %w", err)
	}

	llm, err := llm.NewChatModel("llama3")
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	llmMessage, err := llm.GenerateResponse(ctx, messages)
	if err != nil {
		return fmt.Errorf("failed to generate a response: %w", err)
	}

	_, err = a.chatting.SendMessage(ctx, chatID, llmMessage)
	if err != nil {
		return fmt.Errorf("failed to send a message: %w", err)
	}

	messageSent := domain.NewMessageSent(chatID, llmMessage)
	if err := a.events.Publish(ctx, chatID.String(), messageSent.AsBytes()); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	return nil
}

// AllMessages returns all messages in the chat.
func (a *App) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	return a.messages.AllMessages(ctx, chatID)
}
