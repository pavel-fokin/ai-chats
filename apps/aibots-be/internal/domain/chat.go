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
	CreatedAt time.Time `json:"created_at"`
	CreatedBy User      `json:"created_by"`
}

func NewChat(createdBy User) Chat {
	today := time.Now().Format("02 Jan 15:04:05")
	title := "Chat - " + today

	return Chat{
		ID:        uuid.New(),
		Title:     title,
		CreatedAt: time.Now().UTC(),
		CreatedBy: createdBy,
	}
}
