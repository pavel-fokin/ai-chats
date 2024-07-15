package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

// Messages implements a repository of messages.
type Messages struct {
	DB
}

// NewMessages creates a new messages repository.
func NewMessages(db *sql.DB) *Messages {
	return &Messages{DB{db: db}}
}

func (m *Messages) Add(ctx context.Context, chatID uuid.UUID, message domain.Message) error {
	_, err := m.DBTX(ctx).ExecContext(
		ctx,
		`INSERT INTO message
		(id, chat_id, sender, text, created_at)
		VALUES (?, ?, ?, ?, ?)`,
		message.ID,
		chatID,
		message.Sender,
		message.Text,
		message.CreatedAt.Format(time.RFC3339Nano),
	)
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}

	return nil
}

func (c *Messages) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	rows, err := c.DBTX(ctx).QueryContext(
		ctx,
		`SELECT message.id, sender, text, created_at
		FROM message
		WHERE chat_id = ?`,
		chatID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to select messages: %w", err)
	}
	defer rows.Close()

	var messages []domain.Message
	for rows.Next() {
		var (
			message   domain.Message
			createdAt string
			err       error
		)
		if err := rows.Scan(&message.ID, &message.Sender, &message.Text, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}

		message.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse message.created_at: %w", err)
		}

		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to select messages: %w", err)
	}

	return messages, nil
}
