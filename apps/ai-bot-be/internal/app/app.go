package app

import "context"

type ChatBot interface {
	SingleMessage(ctx context.Context, message string) (Message, error)
	ChatMessage(ctx context.Context, history []string, message string) (Message, error)
}

type App struct {
	chatbot ChatBot
}

func New(chatbot ChatBot) *App {
	return &App{chatbot: chatbot}
}
