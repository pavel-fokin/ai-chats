package domain

import (
	"ai-chats/internal/pkg/types"
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
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`

	Events []types.Message
}

func NewChat(user User, modelID ModelID) Chat {
	return Chat{
		ID:           NewChatID(),
		Title:        "New chat",
		User:         user,
		DefaultModel: modelID,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
}

func (c *Chat) AddMessage(message Message) {
	c.Messages = append(c.Messages, message)
	c.UpdatedAt = time.Now().UTC()
	c.Events = append(c.Events, NewMessageAdded(c.ID, message))
}

func (c *Chat) UpdateTitle(title string) {
	c.Title = title
	c.UpdatedAt = time.Now().UTC()
	c.Events = append(c.Events, NewChatTitleUpdated(c.ID, title))
}
