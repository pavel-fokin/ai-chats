package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type Events interface {
	Subscribe(context.Context, string) (chan []byte, error)
	Unsubscribe(context.Context, string, chan []byte) error
	Publish(context.Context, string, []byte) error
}

type App struct {
	users    domain.Users
	chats    domain.Chats
	messages domain.Messages
	events   Events
}

func New(
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
	events Events,
) *App {
	return &App{
		chats:    chats,
		users:    users,
		messages: messages,
		events:   events,
	}
}
