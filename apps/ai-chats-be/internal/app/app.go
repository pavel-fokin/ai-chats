package app

import (
	"context"

	"ai-chats/internal/domain"
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
	users        domain.Users
	chats        domain.Chats
	models       domain.Models
	ollamaClient domain.OllamaClient
	pubsub       PubSub
	tx           Tx
}

func New(
	chats domain.Chats,
	users domain.Users,
	models domain.Models,
	ollamaClient domain.OllamaClient,
	pubsub PubSub,
	tx Tx,
) *App {
	return &App{
		chats:        chats,
		users:        users,
		models:       models,
		ollamaClient: ollamaClient,
		pubsub:       pubsub,
		tx:           tx,
	}
}
