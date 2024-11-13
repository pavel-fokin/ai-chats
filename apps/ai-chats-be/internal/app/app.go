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
	*Ollama
	*LLM
	chats  domain.Chats
	pubsub PubSub
	tx     Tx
}

func New(
	auth *Auth,
	ollama *Ollama,
	llm *LLM,
	chats domain.Chats,
	pubsub PubSub,
	tx Tx,
) *App {
	return &App{
		Auth:   auth,
		Ollama: ollama,
		LLM:    llm,
		chats:  chats,
		pubsub: pubsub,
		tx:     tx,
	}
}
