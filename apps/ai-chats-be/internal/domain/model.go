package domain

import "context"

// ChatResponseFunc is a function that streams the response of a model.
type ChatResponseFunc func(message Message) error

// Model is an interface for a model.
type Model interface {
	ID() ModelID
	Chat(ctx context.Context, messages []Message, fn ChatResponseFunc) (Message, error)
}
