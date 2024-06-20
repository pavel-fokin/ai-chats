package domain

import (
	"context"

	"github.com/google/uuid"
)

type Chats interface {
	Add(ctx context.Context, chat Chat) error
	Delete(ctx context.Context, chatID ChatID) error
	UpdateTitle(ctx context.Context, chatID ChatID, title string) error
	AllChats(ctx context.Context, userID uuid.UUID) ([]Chat, error)
	FindByID(ctx context.Context, chatID ChatID) (Chat, error)
	Exists(ctx context.Context, chatID ChatID) (bool, error)
}
