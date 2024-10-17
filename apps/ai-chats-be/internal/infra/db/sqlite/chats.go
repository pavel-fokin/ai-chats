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

var _ domain.Chats = (*Chats)(nil)

func NewChats(db *sql.DB) *Chats {
	return &Chats{DB{db: db}}
}

// Add adds a chat to the database.
func (c *Chats) Add(ctx context.Context, chat domain.Chat) error {
	_, err := c.DBTX(ctx).Exec(
		`INSERT INTO chat
		(id, title, user_id, default_model_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		chat.ID,
		chat.Title,
		chat.User.ID,
		chat.DefaultModel.String(),
		chat.CreatedAt.Format(time.RFC3339Nano),
		chat.UpdatedAt.Format(time.RFC3339Nano),
	)
	if err != nil {
		return fmt.Errorf("failed to insert chat: %w", err)
	}

	for _, message := range chat.Messages {
		if err := c.AddMessage(ctx, chat.ID, message); err != nil {
			return fmt.Errorf("failed to add message: %w", err)
		}
	}

	return nil
}

func (c *Chats) Update(ctx context.Context, chat domain.Chat) error {
	result, err := c.DBTX(ctx).ExecContext(
		ctx,
		`UPDATE chat
			SET updated_at = ?
			WHERE id = ? AND deleted_at IS NULL`,
		chat.UpdatedAt.Format(time.RFC3339Nano),
		chat.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update chat: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return domain.ErrChatNotFound
	}

	for _, event := range chat.Events {
		switch event.Type() {
		case domain.MessageAddedType:
			messageAdded := event.(domain.MessageAdded)
			if err := c.AddMessage(ctx, chat.ID, messageAdded.Message); err != nil {
				return fmt.Errorf("failed to add message: %w", err)
			}
		case domain.ChatTitleUpdatedType:
			chatTitleUpdated := event.(domain.ChatTitleUpdated)
			if err := c.UpdateTitle(ctx, chat.ID, chatTitleUpdated.Title); err != nil {
				return fmt.Errorf("failed to update chat title: %w", err)
			}
		default:
			return fmt.Errorf("unknown event type: %s", event.Type())
		}
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

func (c *Chats) AllChats(ctx context.Context, userID domain.UserID) ([]domain.Chat, error) {
	rows, err := c.DBTX(ctx).QueryContext(
		ctx,
		`SELECT
		chat.id, chat.title, chat.default_model_id, chat.created_at, user.id, user.username
		FROM chat
		LEFT JOIN user ON chat.user_id = user.id
		WHERE chat.user_id = ? AND chat.deleted_at IS NULL`,
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
			modelID   string
		)
		if err := rows.Scan(
			&chat.ID,
			&chat.Title,
			&modelID,
			&createdAt,
			&chat.User.ID,
			&chat.User.Username,
		); err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}

		chat.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse chat.created_at: %w", err)
		}
		chat.DefaultModel = domain.NewModelID(modelID)

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
		modelID   string
	)
	err := c.DBTX(ctx).QueryRowContext(
		ctx,
		`SELECT
		chat.id, chat.title, chat.default_model_id, chat.created_at, user.id, user.username
		FROM chat
		LEFT JOIN user ON chat.user_id = user.id
		WHERE chat.id = ? AND chat.deleted_at IS NULL`,
		chatID,
	).Scan(&chat.ID, &chat.Title, &modelID, &createdAt, &chat.User.ID, &chat.User.Username)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find chat by id: %w", err)
	}

	chat.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAt)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to parse chat.created_at: %w", err)
	}
	chat.DefaultModel = domain.NewModelID(modelID)

	return chat, nil
}

func (c *Chats) FindByIDWithMessages(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	chat, err := c.FindByID(ctx, chatID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to find chat by id: %w", err)
	}

	messages, err := c.AllMessages(ctx, chatID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to get messages: %w", err)
	}

	chat.Messages = messages

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
