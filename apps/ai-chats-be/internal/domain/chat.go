package domain

import (
	"time"

	"github.com/google/uuid"
)

type ChatID = uuid.UUID

func NewChatID() ChatID {
	return uuid.New()
}

type Chat struct {
	ID           ChatID    `json:"id"`
	Title        string    `json:"title"`
	User         User      `json:"user"`
	Messages     []Message `json:"messages"`
	DefaultModel ModelID   `json:"defaultModel"`
	CreatedAt    time.Time `json:"createdAt"`
	DeletedAt    time.Time `json:"deletedAt"`
}

func NewChat(user User, modelID ModelID) Chat {
	return Chat{
		ID:           NewChatID(),
		Title:        "New chat",
		User:         user,
		DefaultModel: modelID,
		CreatedAt:    time.Now().UTC(),
	}
}

func (c *Chat) AddMessage(message Message) {
	c.Messages = append(c.Messages, message)
}
