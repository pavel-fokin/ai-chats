package sqlite

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"pavel-fokin/ai/apps/ai-bot/internal/app/domain"
)

type Actor struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid"`
	Type string
}

func (a *Actor) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()
	return nil
}

type Chat struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Actors   []Actor   `gorm:"many2many:chat_to_actor"`
	Messages []Message
}

func (c *Chat) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}

type ChatToActor struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid"`
	ChatID  string
	ActorID string
}

func (c *ChatToActor) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}

type Message struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	// Chat    Chat
	Actor   Actor
	Text    string
	ChatID  uuid.UUID
	ActorID uuid.UUID
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	return nil
}

func (s *Sqlite) CreateChat(ctx context.Context, actors []domain.Actor) (*domain.Chat, error) {
	dbActors := make([]Actor, len(actors))
	for i, actor := range actors {
		dbActors[i] = Actor{
			ID:   actor.ID,
			Type: actor.Type,
		}
	}

	chat := Chat{
		Actors: dbActors,
	}
	tx := s.db.Create(&chat)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return domain.NewChat(chat.ID, actors), nil
}

func (s *Sqlite) CreateActor(ctx context.Context, actorType string) (*domain.Actor, error) {
	actor := Actor{
		Type: actorType,
	}
	tx := s.db.Create(&actor)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &domain.Actor{
		ID:   actor.ID,
		Type: actor.Type,
	}, nil
}

func (s *Sqlite) FindChat(ctx context.Context, chatID uuid.UUID) (*domain.Chat, error) {
	var chat Chat
	tx := s.db.First(&chat, "id = ?", chatID).Preload("Actors")
	if tx.Error != nil {
		return nil, tx.Error
	}

	var actors []domain.Actor
	for _, actor := range chat.Actors {
		actors = append(actors, domain.Actor{
			ID:   actor.ID,
			Type: actor.Type,
		})
	}

	return domain.NewChat(chat.ID, actors), nil
}

func (s *Sqlite) AddMessage(ctx context.Context, chat *domain.Chat, actor *domain.Actor, message string) error {
	dbActor := Actor{
		ID:   actor.ID,
		Type: actor.Type,
	}
	dbChat := Chat{
		ID: chat.ID,
	}
	dbChat.Messages = append(dbChat.Messages, Message{
		Actor: dbActor,
		Text:  message,
	})
	tx := s.db.Save(&dbChat)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (s *Sqlite) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	var dbMessages []Message
	tx := s.db.Preload("Actor").Find(&dbMessages, "chat_id = ?", chatID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var messages []domain.Message
	for _, message := range dbMessages {
		messages = append(messages, domain.NewMessage(message.ID, &domain.Actor{ID: message.ActorID, Type: message.Actor.Type}, message.Text))
	}

	return messages, nil
}

func (s *Sqlite) FindActor(ctx context.Context, actorID uuid.UUID) (*domain.Actor, error) {
	var actor Actor
	tx := s.db.First(&actor, "id = ?", actorID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &domain.Actor{
		ID:   actor.ID,
		Type: actor.Type,
	}, nil
}
