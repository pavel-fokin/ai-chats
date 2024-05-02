package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"

	"github.com/google/uuid"
)

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (a *App) CreateChat(ctx context.Context) (domain.Chat, error) {
	ai, err := a.chatDB.FindActorByType(ctx, domain.AI)
	if err != nil {
		return domain.Chat{}, err
	}

	human, err := a.chatDB.FindActorByType(ctx, domain.Human)
	if err != nil {
		return domain.Chat{}, err
	}

	return a.chatDB.CreateChat(ctx, []domain.Actor{ai, human})
}

func (a *App) SendMessage(ctx context.Context, chatID uuid.UUID, message string) (Message, error) {
	chat, err := a.chatDB.FindChat(ctx, chatID)
	if err != nil {
		return Message{}, err
	}

	human, err := a.chatDB.FindActorByType(ctx, domain.Human)
	if err != nil {
		return Message{}, err
	}

	history, err := a.chatDB.AllMessages(ctx, chat.ID)
	if err != nil {
		return Message{}, err
	}

	if err := a.chatDB.AddMessage(ctx, chat, human, message); err != nil {
		return Message{}, err
	}

	aiMessage, err := a.chatbot.ChatMessage(ctx, history, message)
	if err != nil {
		return Message{}, err
	}

	ai, err := a.chatDB.FindActorByType(ctx, domain.AI)
	if err != nil {
		return Message{}, err
	}

	if err := a.chatDB.AddMessage(ctx, chat, ai, aiMessage.Text); err != nil {
		return Message{}, err
	}

	return aiMessage, nil
}
