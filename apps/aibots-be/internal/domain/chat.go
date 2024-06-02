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

type Chat struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedBy User      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewChat(createdBy User) Chat {
	return Chat{
		ID:        uuid.New(),
		Title:     "New chat",
		CreatedAt: time.Now().UTC(),
		CreatedBy: createdBy,
	}
}
