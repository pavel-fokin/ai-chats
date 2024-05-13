package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/llm"

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

	llm, err := llm.NewChatModel("llama3")
	if err != nil {
		return domain.Message{}, fmt.Errorf("failed to create a chat model: %w", err)
	}

	return a.chatting.SendMessage(ctx, llm, chatID, message)
}

func (a *App) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	return a.messages.AllMessages(ctx, chatID)
}
