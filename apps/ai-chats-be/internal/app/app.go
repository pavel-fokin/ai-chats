package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

// PubSub is a publish/subscribe interface.
type PubSub interface {
	Subscribe(context.Context, string) (chan []byte, error)
	Unsubscribe(context.Context, string, chan []byte) error
	Publish(context.Context, string, []byte) error
}

// Tx is a transaction interface.
type Tx interface {
	Tx(context.Context, func(context.Context) error) error
}

type App struct {
	users    domain.Users
	chats    domain.Chats
	messages domain.Messages
	ollama   domain.Ollama
	pubsub   PubSub
	tx       Tx
}

func New(
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
	ollama domain.Ollama,
	pubsub PubSub,
	tx Tx,
) *App {
	return &App{
		chats:    chats,
		users:    users,
		messages: messages,
		ollama:   ollama,
		pubsub:   pubsub,
		tx:       tx,
	}
}
