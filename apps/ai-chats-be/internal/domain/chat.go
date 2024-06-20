package domain

import (
	"time"

	"github.com/google/uuid"
)

type SenderType string

const (
	AI    SenderType = "AI"
	Human SenderType = "human"
)

type ChatID = uuid.UUID

func NewChatID() ChatID {
	return uuid.New()
}

type Chat struct {
	ID        ChatID    `json:"id"`
	Title     string    `json:"title"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewChat(user User) Chat {
	return Chat{
		ID:        NewChatID(),
		Title:     "New chat",
		User:      user,
		CreatedAt: time.Now().UTC(),
	}
}
