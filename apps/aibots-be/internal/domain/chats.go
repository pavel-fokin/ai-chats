package domain

import (
	"context"

	"github.com/google/uuid"
)

type Chats interface {
	Add(ctx context.Context, chat Chat) error
	AllChats(ctx context.Context, userID uuid.UUID) ([]Chat, error)
	FindChat(ctx context.Context, chatID uuid.UUID) (Chat, error)
	Exists(ctx context.Context, chatID uuid.UUID) (bool, error)
}
