package domain

import (
	"context"

	"github.com/google/uuid"
)

type Chats interface {
	Add(ctx context.Context, chat Chat) error
	AddMessage(ctx context.Context, chatID ChatID, message Message) error
	AllChats(ctx context.Context, userID uuid.UUID) ([]Chat, error)
	AllMessages(ctx context.Context, chatID ChatID) ([]Message, error)
	Delete(ctx context.Context, chatID ChatID) error
	Exists(ctx context.Context, chatID ChatID) (bool, error)
	FindByID(ctx context.Context, chatID ChatID) (Chat, error)
	UpdateTitle(ctx context.Context, chatID ChatID, title string) error
}
