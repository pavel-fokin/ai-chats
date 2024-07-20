package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type Topic = string

// PubSub is a publish/subscribe interface.
type PubSub interface {
	Subscribe(context.Context, Topic) (chan []byte, error)
	Unsubscribe(context.Context, Topic, chan []byte) error
	Publish(context.Context, Topic, []byte) error
}

// Tx is a transaction interface.
type Tx interface {
	Tx(context.Context, func(context.Context) error) error
}

type App struct {
	users  domain.Users
	chats  domain.Chats
	models domain.Models
	ollama domain.Ollama
	pubsub PubSub
	tx     Tx
}

func New(
	chats domain.Chats,
	users domain.Users,
	models domain.Models,
	ollama domain.Ollama,
	pubsub PubSub,
	tx Tx,
) *App {
	return &App{
		chats:  chats,
		users:  users,
		models: models,
		ollama: ollama,
		pubsub: pubsub,
		tx:     tx,
	}
}
