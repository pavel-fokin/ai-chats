package app

import (
	"context"

	"ai-chats/internal/domain"
)

// Tx is a transaction interface.
type Tx interface {
	Tx(context.Context, func(context.Context) error) error
}

type Config struct {
	HashCost int
}

type App struct {
	config        Config
	users         domain.Users
	chats         domain.Chats
	modelsLibrary domain.ModelsLibrary
	ollamaClient  OllamaClient
	ollamaModels  domain.OllamaModels
	pubsub        PubSub
	tx            Tx
}

func New(
	chats domain.Chats,
	users domain.Users,
	modelsLibrary domain.ModelsLibrary,
	ollamaClient OllamaClient,
	ollamaModels domain.OllamaModels,
	pubsub PubSub,
	tx Tx,
) *App {
	return &App{
		config: Config{
			HashCost: 14,
		},
		chats:         chats,
		users:         users,
		modelsLibrary: modelsLibrary,
		ollamaClient:  ollamaClient,
		ollamaModels:  ollamaModels,
		pubsub:        pubsub,
		tx:            tx,
	}
}
