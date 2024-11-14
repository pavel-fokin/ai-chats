package domain

import (
	"context"
)

// ModelStreamMessage is a struct that streams the response of a model.
type ModelStreamMessage struct {
	Sender Sender `json:"sender"`
	Text   string `json:"text"`
}

// NewModelStreamMessage creates a new model stream message.
func NewModelStreamMessage(sender Sender, text string) ModelStreamMessage {
	return ModelStreamMessage{Sender: sender, Text: text}
}

// ModelResponseFunc is a function that streams the response of a model.
type ModelResponseFunc func(message ModelStreamMessage) error

// Model is an interface for a model.
type Model interface {
	Chat(ctx context.Context, messages []Message, fn ModelResponseFunc) (Message, error)
}
