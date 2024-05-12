package domain

import "github.com/google/uuid"

type SenderType string

const (
	AI    SenderType = "AI"
	Human SenderType = "human"
)

type Chat struct {
	ID uuid.UUID `json:"id"`
}

func NewChat() Chat {
	return Chat{
		ID: uuid.New(),
	}
}
