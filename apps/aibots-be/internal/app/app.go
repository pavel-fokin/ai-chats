package app

import (
	"context"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/events"
)

type MessageSender interface {
	SendMessage(ctx context.Context, chatID uuid.UUID, message domain.Message) (domain.MessageSent, error)
}

type App struct {
	users    domain.Users
	chats    domain.Chats
	messages domain.Messages
	chatting MessageSender
	events   events.Events
}

func New(
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
	events events.Events,
) *App {
	chatting := domain.NewChatting(chats, messages)

	return &App{
		chats:    chats,
		users:    users,
		messages: messages,
		chatting: chatting,
		events:   events,
	}
}
