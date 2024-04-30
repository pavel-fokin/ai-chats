package app

import (
	"context"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func (a *App) SendMessage(ctx context.Context, chatID, message string) (Message, error) {
	return a.chatbot.SingleMessage(ctx, message)
}
