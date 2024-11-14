package app

import (
	"ai-chats/internal/app/notifications"
	"context"
)

// Tx is a transaction interface.
type Tx interface {
	Tx(context.Context, func(context.Context) error) error
}

// Notificator is an interface that notifies about an event.
type Notificator interface {
	Notify(ctx context.Context, notification notifications.Notification) error
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
		Chats:  chats,
		LLM:    llm,
		Ollama: ollama,
	}
}
