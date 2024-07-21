package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"ai-chats/internal/domain"
)

// Chats implements a repository for chats.
type Chats struct {
	DB
}

func NewChats(db *sql.DB) *Chats {
	return &Chats{DB{db: db}}
}

func (c *Chats) Add(ctx context.Context, chat domain.Chat) error {
	_, err := c.DBTX(ctx).Exec(
		`INSERT INTO chat
		(id, title, user_id, default_model, created_at)
		VALUES (?, ?, ?, ?, ?)`,
		chat.ID,
		chat.Title,
		chat.User.ID,
		chat.DefaultModel.String(),
		chat.CreatedAt.Format(time.RFC3339Nano),
	)
	if err != nil {
		return fmt.Errorf("failed to insert chat: %w", err)
	}

	return nil
}

func (m *Chats) AddMessage(ctx context.Context, chatID domain.ChatID, message domain.Message) error {
	_, err := m.DBTX(ctx).ExecContext(
		ctx,
		`INSERT INTO message
		(id, chat_id, sender, text, created_at)
		VALUES (?, ?, ?, ?, ?)`,
		message.ID,
		chatID,
		message.Sender.String(),
		message.Text,
		message.CreatedAt.Format(time.RFC3339Nano),
	)
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}

	return nil
}

func (c *Chats) AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error) {
	chatExists, err := c.exists(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to check chat existence: %w", err)
	}
	if !chatExists {
		return nil, domain.ErrChatNotFound
	}

	rows, err := c.DBTX(ctx).QueryContext(
		ctx,
		`SELECT id, sender, text, created_at
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
			createdAt string
			err       error
			message   domain.Message
			sender    string
		)
		if err := rows.Scan(&message.ID, &sender, &message.Text, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}

		message.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse message.created_at: %w", err)
		}

		message.Sender = domain.NewSender(sender)

		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to select messages: %w", err)
	}

	return messages, nil
}

func (c *Chats) Delete(ctx context.Context, chatID uuid.UUID) error {
	result, err := c.DBTX(ctx).ExecContext(
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
	_, err := c.DBTX(ctx).ExecContext(
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
	rows, err := c.DBTX(ctx).QueryContext(
		ctx,
		`SELECT
		id, title, default_model, created_at
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
		var (
			chat      domain.Chat
			createdAt string
			model     string
		)
		if err := rows.Scan(&chat.ID, &chat.Title, &model, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}

		chat.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse chat.created_at: %w", err)
		}
		chat.DefaultModel = domain.NewModel(model)

		chats = append(chats, chat)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to select chats: %w", rows.Err())
	}

	return chats, nil
}

func (c *Chats) FindByID(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	var (
		chat      domain.Chat
		createdAt string
		model     string
	)
	err := c.DBTX(ctx).QueryRowContext(
		ctx,
		`SELECT
		id, title, default_model, created_at
		FROM chat
		WHERE id = ? AND deleted_at IS NULL`,
		chatID,
	).Scan(&chat.ID, &chat.Title, &model, &createdAt)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find chat by id: %w", err)
	}

	chat.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to parse chat.created_at: %w", err)
	}
	chat.DefaultModel = domain.NewModel(model)

	return chat, nil
}

func (c *Chats) exists(ctx context.Context, chatID uuid.UUID) (bool, error) {
	var exists bool
	err := c.DBTX(ctx).QueryRowContext(
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

func (c *Chats) Exists(ctx context.Context, chatID uuid.UUID) (bool, error) {
	return c.exists(ctx, chatID)
}
