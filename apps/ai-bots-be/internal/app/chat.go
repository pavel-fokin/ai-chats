package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"github.com/google/uuid"
)

func (a *App) CreateChat(ctx context.Context, userID uuid.UUID) (domain.Chat, error) {
	ai, err := a.chats.FindActorByType(ctx, domain.AI)
	if err != nil {
		return domain.Chat{}, err
	}

	human, err := a.chats.FindActorByType(ctx, domain.Human)
	if err != nil {
		return domain.Chat{}, err
	}

	return a.chats.CreateChat(ctx, userID, []domain.Actor{ai, human})
}

func (a *App) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	return a.chats.AllChats(ctx, userID)
}

func (a *App) SendMessage(ctx context.Context, chatID uuid.UUID, message string) (domain.Message, error) {
	chat, err := a.chats.FindChat(ctx, chatID)
	if err != nil {
		return domain.Message{}, err
	}

	human, err := a.chats.FindActorByType(ctx, domain.Human)
	if err != nil {
		return domain.Message{}, err
	}

	history, err := a.chats.AllMessages(ctx, chat.ID)
	if err != nil {
		return domain.Message{}, err
	}

	if err := a.chats.AddMessage(ctx, chat, human, message); err != nil {
		return domain.Message{}, err
	}

	aiMessage, err := a.chatbot.ChatMessage(ctx, history, message)
	if err != nil {
		return domain.Message{}, err
	}

	ai, err := a.chats.FindActorByType(ctx, domain.AI)
	if err != nil {
		return domain.Message{}, err
	}

	if err := a.chats.AddMessage(ctx, chat, ai, aiMessage.Text); err != nil {
		return domain.Message{}, err
	}

	return aiMessage, nil
}

func (a *App) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	return a.chats.AllMessages(ctx, chatID)
}
