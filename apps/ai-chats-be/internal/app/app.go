package app

import (
	"context"

	"ai-chats/internal/domain"
)

// Tx is a transaction interface.
type Tx interface {
	Tx(context.Context, func(context.Context) error) error
}

type App struct {
	*Auth
	users         domain.Users
	chats         domain.Chats
	modelsLibrary domain.ModelsLibrary
	ollamaClient  OllamaClient
	ollamaModels  domain.OllamaModels
	pubsub        PubSub
	tx            Tx
}

func New(
	auth *Auth,
	chats domain.Chats,
	users domain.Users,
	modelsLibrary domain.ModelsLibrary,
	ollamaClient OllamaClient,
	ollamaModels domain.OllamaModels,
	pubsub PubSub,
	tx Tx,
) *App {
	return &App{
		Auth:          auth,
		chats:         chats,
		users:         users,
		modelsLibrary: modelsLibrary,
		ollamaClient:  ollamaClient,
		ollamaModels:  ollamaModels,
		pubsub:        pubsub,
		tx:            tx,
	}
}
