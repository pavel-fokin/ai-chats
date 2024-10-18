package domain

import (
	"context"
)

type Chats interface {
	// Add adds a chat to the database.
	Add(ctx context.Context, chat Chat) error
	// Delete deletes a chat from the database.
	Delete(ctx context.Context, chatID ChatID) error
	// Exists checks if a chat exists in the database.
	Exists(ctx context.Context, chatID ChatID) (bool, error)
	// FindByID returns a chat from the database.
	FindByID(ctx context.Context, chatID ChatID) (Chat, error)
	// FindByIDWithMessages returns a chat from the database with messages.
	FindByIDWithMessages(ctx context.Context, chatID ChatID) (Chat, error)
	// FindByUserID returns all chats from the database.
	FindByUserID(ctx context.Context, userID UserID) ([]Chat, error)
	// Update updates a chat in the database.
	Update(ctx context.Context, chat Chat) error
}
