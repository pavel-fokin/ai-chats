package domain

import (
	"context"

	"github.com/google/uuid"
)

// Messages represents a repository of messages.
type Messages interface {
	AddMessage(ctx context.Context, chat Chat, sender, message string) error
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]Message, error)
}
