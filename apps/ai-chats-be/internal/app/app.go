package app

import (
	"context"

	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
)

type Topic = string

// PubSub is a publish/subscribe interface.
type PubSub interface {
	Subscribe(context.Context, Topic) (chan events.Event, error)
	Unsubscribe(context.Context, Topic, chan events.Event) error
	Publish(context.Context, Topic, events.Event) error
}

// Tx is a transaction interface.
type Tx interface {
	Tx(context.Context, func(context.Context) error) error
}

type Config struct {
	HashCost int
}

type App struct {
	config       Config
	users        domain.Users
	chats        domain.Chats
	models       domain.Models
	ollamaClient domain.OllamaClient
	ollamaModels domain.OllamaModels
	pubsub       PubSub
	tx           Tx
}

func New(
	chats domain.Chats,
	users domain.Users,
	models domain.Models,
	ollamaClient domain.OllamaClient,
	ollamaModels domain.OllamaModels,
	pubsub PubSub,
	tx Tx,
) *App {
	return &App{
		config: Config{
			HashCost: 14,
		},
		chats:        chats,
		users:        users,
		models:       models,
		ollamaClient: ollamaClient,
		ollamaModels: ollamaModels,
		pubsub:       pubsub,
		tx:           tx,
	}
}
