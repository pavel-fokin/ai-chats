package domain

import (
	"context"
)

// Messages represents a repository of messages.
type Messages interface {
	Add(ctx context.Context, chatID ChatID, message Message) error
	AllMessages(ctx context.Context, chatID ChatID) ([]Message, error)
}
