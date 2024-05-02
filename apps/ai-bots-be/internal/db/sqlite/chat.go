package sqlite

import (
	"context"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"
)

func (db *Sqlite) CreateChat(ctx context.Context, actors []domain.Actor) (domain.Chat, error) {
	chat := domain.Chat{
		ID:     uuid.New(),
		Actors: actors,
	}

	_, err := db.db.ExecContext(ctx, "INSERT INTO chat (id) VALUES (?)", chat.ID)
	if err != nil {
		return domain.Chat{}, err
	}

	for _, actor := range actors {
		_, err := db.db.ExecContext(
			ctx,
			"INSERT INTO chat_actor (chat_id, actor_id) VALUES (?, ?)",
			chat.ID,
			actor.ID,
		)
		if err != nil {
			return domain.Chat{}, err
		}
	}

	return chat, nil
}

func (db *Sqlite) AllChats(ctx context.Context) ([]domain.Chat, error) {
	rows, err := db.db.QueryContext(ctx, "SELECT id FROM chat")
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

	return chats, nil
}

func (db *Sqlite) FindChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	var chat domain.Chat
	err := db.db.QueryRowContext(ctx, "SELECT id FROM chat WHERE id = ?", chatID).Scan(&chat.ID)
	if err != nil {
		return domain.Chat{}, err
	}

	return chat, nil
}

func (db *Sqlite) CreateActor(ctx context.Context, actorType domain.ActorType) (domain.Actor, error) {
	actor := domain.Actor{
		ID:   uuid.New(),
		Type: actorType,
	}

	_, err := db.db.ExecContext(ctx, "INSERT INTO actor (id, type) VALUES (?, ?)", actor.ID, actor.Type)
	if err != nil {
		return domain.Actor{}, err
	}

	return actor, nil
}

func (db *Sqlite) FindActor(ctx context.Context, actorID uuid.UUID) (domain.Actor, error) {
	var actor domain.Actor
	err := db.db.QueryRowContext(ctx, "SELECT id, type FROM actor WHERE id = ?", actorID).Scan(&actor.ID, &actor.Type)
	if err != nil {
		return domain.Actor{}, err
	}

	return actor, nil
}

func (db *Sqlite) FindActorByType(ctx context.Context, actorType domain.ActorType) (domain.Actor, error) {
	var actor domain.Actor
	err := db.db.QueryRowContext(ctx, "SELECT id, type FROM actor WHERE type = ?", actorType).Scan(&actor.ID, &actor.Type)
	if err != nil {
		return domain.Actor{}, err
	}

	return actor, nil
}

func (db *Sqlite) AddMessage(ctx context.Context, chat domain.Chat, actor domain.Actor, message string) error {
	messageID := uuid.New()

	_, err := db.db.ExecContext(
		ctx,
		"INSERT INTO message (id, chat_id, actor_id, text) VALUES (?, ?, ?, ?)",
		messageID, chat.ID, actor.ID, message,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *Sqlite) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	rows, err := db.db.QueryContext(
		ctx,
		`SELECT message.id, actor_id, actor.type, text
		FROM message
		JOIN actor ON message.actor_id = actor.id
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
		if err := rows.Scan(&message.ID, &message.Actor.ID, &message.Actor.Type, &message.Text); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
