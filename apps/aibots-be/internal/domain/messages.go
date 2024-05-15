package domain

import (
	"context"

	"github.com/google/uuid"
)

// Messages represents a repository of messages.
type Messages interface {
	Add(ctx context.Context, chat Chat, message Message) error
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]Message, error)
}
