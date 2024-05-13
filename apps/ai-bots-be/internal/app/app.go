package app

import (
	"context"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type MessageSender interface {
	SendMessage(ctx context.Context, llm domain.LLM, chatID uuid.UUID, message domain.Message) (domain.Message, error)
}

type App struct {
	users    domain.Users
	chats    domain.Chats
	messages domain.Messages
	chatting MessageSender
}

func New(
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
) *App {
	chatting := domain.NewChatting(chats, messages)

	return &App{
		chats:    chats,
		users:    users,
		messages: messages,
		chatting: chatting,
	}
}
