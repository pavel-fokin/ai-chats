package domain

import (
	"context"
)

type Chats interface {
	// Add adds a chat to the database.
	Add(ctx context.Context, chat Chat) error
	// Update updates a chat in the database.
	Update(ctx context.Context, chat Chat) error
	// AllChats returns all chats from the database.
	AllChats(ctx context.Context, userID UserID) ([]Chat, error)
	Delete(ctx context.Context, chatID ChatID) error
	Exists(ctx context.Context, chatID ChatID) (bool, error)
	FindByID(ctx context.Context, chatID ChatID) (Chat, error)
	FindByIDWithMessages(ctx context.Context, chatID ChatID) (Chat, error)
}
