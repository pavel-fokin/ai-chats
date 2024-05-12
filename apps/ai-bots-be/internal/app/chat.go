package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"github.com/google/uuid"
)

func (a *App) CreateChat(ctx context.Context, userID uuid.UUID) (domain.Chat, error) {
	user, err := a.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to create a chat: %w", err)
	}

	return a.chats.CreateChat(ctx, user.ID)
}

func (a *App) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	return a.chats.AllChats(ctx, userID)
}

func (a *App) SendMessage(ctx context.Context, chatID uuid.UUID, message string) (domain.Message, error) {
	chat, err := a.chats.FindChat(ctx, chatID)
	if err != nil {
		return domain.Message{}, err
	}

	history, err := a.messages.AllMessages(ctx, chat.ID)
	if err != nil {
		return domain.Message{}, err
	}

	if err := a.messages.AddMessage(ctx, chat, "User", message); err != nil {
		return domain.Message{}, err
	}

	aiMessage, err := a.chatbot.ChatMessage(ctx, history, message)
	if err != nil {
		return domain.Message{}, err
	}

	if err := a.messages.AddMessage(ctx, chat, "AI", aiMessage.Text); err != nil {
		return domain.Message{}, err
	}

	return aiMessage, nil
}

func (a *App) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	return a.messages.AllMessages(ctx, chatID)
}
