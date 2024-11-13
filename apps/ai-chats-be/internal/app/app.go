package app

import (
	"context"
)

// Tx is a transaction interface.
type Tx interface {
	Tx(context.Context, func(context.Context) error) error
}

// App is the application.
type App struct {
	*Auth
	*Chats
	*LLM
	*Ollama
}

// New creates a new app.
func New(
	auth *Auth,
	chats *Chats,
	llm *LLM,
	ollama *Ollama,
) *App {
	return &App{
		Auth:   auth,
		Ollama: ollama,
		LLM:    llm,
		Chats:  chats,
	}
}
