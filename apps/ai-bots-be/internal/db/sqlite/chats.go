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

func (c *Chats) CreateChat(ctx context.Context, userId uuid.UUID) (domain.Chat, error) {
	chat := domain.NewChat()

	_, err := c.db.ExecContext(ctx, "INSERT INTO chat (id) VALUES (?)", chat.ID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to insert chat: %w", err)
	}

	_, err = c.db.ExecContext(ctx, "INSERT INTO chat_user (chat_id, user_id) VALUES (?, ?)", chat.ID, userId)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("failed to insert chat_user: %w", err)
	}

	return chat, nil
}

func (c *Chats) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	rows, err := c.db.QueryContext(ctx, "SELECT chat_id FROM chat_user WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatIDs []uuid.UUID
	for rows.Next() {
		var chatID uuid.UUID
		err := rows.Scan(&chatID)
		if err != nil {
			return nil, err
		}
		chatIDs = append(chatIDs, chatID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(chatIDs) == 0 {
		return nil, nil
	}

	query, args := QueryIn("SELECT id FROM chat WHERE chat.id", chatIDs)
	rows, err = c.db.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []domain.Chat
	for rows.Next() {
		var chat domain.Chat
		err := rows.Scan(&chat.ID)
		if err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
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
