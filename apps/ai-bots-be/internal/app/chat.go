package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/llm"
	"pavel-fokin/ai/apps/ai-bots-be/internal/worker"

	"github.com/google/uuid"
)

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

func (a *App) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	return a.chats.AllChats(ctx, userID)
}

func (a *App) SendMessage(ctx context.Context, chatID uuid.UUID, text string) (domain.Message, error) {
	message := domain.NewMessage("User", text)

	messageSent, err := a.chatting.SendMessage(ctx, chatID, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("failed to send a message: %w", err)
	}

	if err := a.messageSentEvents.Publish(ctx, chatID.String(), messageSent); err != nil {
		return domain.Message{}, fmt.Errorf("failed to publish a message: %w", err)
	}

	if err := a.generateResponseEvents.Publish(
		ctx,
		"generate-response",
		worker.GenerateResponse{
			ChatID: chatID,
		},
	); err != nil {
		return domain.Message{}, fmt.Errorf("failed to generate a response: %w", err)
	}

	return domain.Message{}, nil
}

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

	messageSent, err := a.chatting.SendMessage(ctx, chatID, llmMessage)
	if err != nil {
		return fmt.Errorf("failed to send a message: %w", err)
	}

	if err := a.messageSentEvents.Publish(ctx, chatID.String(), messageSent); err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	return nil
}

func (a *App) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	return a.messages.AllMessages(ctx, chatID)
}

func (a *App) Subscribe(ctx context.Context, chatID string, subscriber string) (<-chan domain.MessageSent, error) {
	return a.messageSentEvents.Subscribe(ctx, chatID, subscriber)
}

func (a *App) Unsubscribe(ctx context.Context, chatID string, subscriber string) error {
	return a.messageSentEvents.Unsubscribe(ctx, chatID, subscriber)
}
