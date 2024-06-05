package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
	_, err := c.db.ExecContext(
		ctx,
		`INSERT INTO chat
		(id, title, created_at, user_id)
		VALUES (?, ?, ?, ?)`,
		chat.ID,
		chat.Title,
		chat.CreatedAt.Format(time.RFC3339Nano),
		chat.User.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert chat: %w", err)
	}

	return nil
}

func (c *Chats) Delete(ctx context.Context, chatID uuid.UUID) error {
	result, err := c.db.ExecContext(
		ctx,
		`UPDATE chat
		SET deleted_at = ?
		WHERE id = ?`,
		time.Now().Format(time.RFC3339Nano),
		chatID,
	)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrChatNotFound
	}

	return nil
}

func (c *Chats) UpdateTitle(ctx context.Context, chatID uuid.UUID, title string) error {
	_, err := c.db.ExecContext(
		ctx,
		`UPDATE chat
		SET title = ?
		WHERE id = ? AND deleted_at IS NULL`,
		title,
		chatID,
	)
	if err != nil {
		return fmt.Errorf("failed to update chat title: %w", err)
	}

	return nil
}

func (c *Chats) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	rows, err := c.db.QueryContext(
		ctx,
		`SELECT
		id, title, created_at
		FROM chat
		WHERE user_id = ? AND deleted_at IS NULL`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to select chats: %w", err)
	}
	defer rows.Close()

	var chats []domain.Chat
	for rows.Next() {
		var chat domain.Chat
		var createdAt string
		if err := rows.Scan(&chat.ID, &chat.Title, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}

		chat.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse created_at: %w", err)
		}

		chats = append(chats, chat)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to select chats: %w", rows.Err())
	}

	return chats, nil
}

func (c *Chats) FindByID(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	var chat domain.Chat
	var createdAt string
	err := c.db.QueryRowContext(
		ctx,
		`SELECT
		id, title, created_at
		FROM chat
		WHERE id = ? AND deleted_at IS NULL`,
		chatID,
	).Scan(&chat.ID, &chat.Title, &createdAt)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find chat by id: %w", err)
	}

	chat.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to parse created_at: %w", err)
	}

	return chat, nil
}

func (c *Chats) Exists(ctx context.Context, chatID uuid.UUID) (bool, error) {
	var exists bool
	err := c.db.QueryRowContext(
		ctx,
		`SELECT EXISTS(
		SELECT 1 FROM chat WHERE id = ? AND deleted_at IS NULL
		)`,
		chatID,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check chat existence: %w", err)
	}

	return exists, nil
}
