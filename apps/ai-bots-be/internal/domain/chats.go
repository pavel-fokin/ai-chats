package domain

import (
	"context"

	"github.com/google/uuid"
)

type Chats interface {
	CreateChat(ctx context.Context, userID uuid.UUID) (Chat, error)
	AllChats(ctx context.Context, userID uuid.UUID) ([]Chat, error)
	FindChat(ctx context.Context, chatID uuid.UUID) (Chat, error)
	AddMessage(ctx context.Context, chat Chat, sender, message string) error
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]Message, error)
}
