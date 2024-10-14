package domain

import "context"

type ChatResponseFunc func(message Message) error

type Model interface {
	ID() ModelID
	Chat(ctx context.Context, messages []Message, fn ChatResponseFunc) (Message, error)
}
