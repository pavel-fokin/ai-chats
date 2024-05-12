package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type ChatBot interface {
	SingleMessage(ctx context.Context, message string) (domain.Message, error)
	ChatMessage(ctx context.Context, history []domain.Message, message string) (domain.Message, error)
}

type App struct {
	chatbot  ChatBot
	users    domain.Users
	chats    domain.Chats
	messages domain.Messages
}

func New(
	chatbot ChatBot,
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
) *App {
	return &App{
		chatbot:  chatbot,
		chats:    chats,
		users:    users,
		messages: messages,
	}
}
