package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

// Chats implements a repository for chats.
type Chats struct {
	db *sql.DB
}

func NewChats(db *sql.DB) *Chats {
	return &Chats{db: db}
}

func (c *Chats) Add(ctx context.Context, chat domain.Chat) error {
	_, err := c.db.ExecContext(ctx, "INSERT INTO chat (id, created_by) VALUES (?, ?)", chat.ID, chat.CreatedBy.ID)
	if err != nil {
		return fmt.Errorf("failed to insert chat: %w", err)
	}

	return nil
}

func (c *Chats) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	rows, err := c.db.QueryContext(ctx, "SELECT id FROM chat WHERE created_by = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to select chats: %w", err)
	}
	defer rows.Close()

	var chats []domain.Chat
	for rows.Next() {
		var chat domain.Chat
		if err := rows.Scan(&chat.ID); err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}
		chats = append(chats, chat)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to select chats: %w", rows.Err())
	}

	return chats, nil
}

func (c *Chats) FindChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	var chat domain.Chat
	err := c.db.QueryRowContext(ctx, "SELECT id FROM chat WHERE id = ?", chatID).Scan(&chat.ID)
	if err != nil {
		return domain.Chat{}, err
	}

	return chat, nil
}
