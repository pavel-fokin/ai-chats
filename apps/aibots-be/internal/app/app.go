package app

import (
	"context"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type MessageSender interface {
	SendMessage(ctx context.Context, chatID uuid.UUID, message domain.Message) error
}
type Events interface {
	Subscribe(context.Context, string) (chan []byte, error)
	Unsubscribe(context.Context, string, chan []byte) error
	Publish(context.Context, string, []byte) error
}

type App struct {
	users    domain.Users
	chats    domain.Chats
	messages domain.Messages
	chatting MessageSender
	events   Events
}

func New(
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
	events Events,
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
