package sqlite

import (
	"context"
	"database/sql"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"github.com/google/uuid"
)

// Messages implements a repository of messages.
type Messages struct {
	db *sql.DB
}

// NewMessages creates a new messages repository.
func NewMessages(db *sql.DB) *Messages {
	return &Messages{db: db}
}

func (m *Messages) AddMessage(ctx context.Context, chat domain.Chat, sender, message string) error {
	messageID := uuid.New()

	_, err := m.db.ExecContext(
		ctx,
		"INSERT INTO message (id, chat_id, sender, text) VALUES (?, ?, ?, ?)",
		messageID, chat.ID, sender, message,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Messages) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	rows, err := c.db.QueryContext(
		ctx,
		`SELECT message.id, text
		FROM message
		WHERE chat_id = ?`,
		chatID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []domain.Message
	for rows.Next() {
		var message domain.Message
		if err := rows.Scan(&message.ID, &message.Text); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
