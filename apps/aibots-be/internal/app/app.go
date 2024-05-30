package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type PubSub interface {
	Subscribe(context.Context, string) (chan []byte, error)
	Unsubscribe(context.Context, string, chan []byte) error
	Publish(context.Context, string, []byte) error
}

type App struct {
	users    domain.Users
	chats    domain.Chats
	messages domain.Messages
	pubsub   PubSub
}

func New(
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
	pubsub PubSub,
) *App {
	return &App{
		chats:    chats,
		users:    users,
		messages: messages,
		pubsub:   pubsub,
	}
}
