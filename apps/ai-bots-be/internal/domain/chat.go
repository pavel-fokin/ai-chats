package domain

import "github.com/google/uuid"

type SenderType string

const (
	AI    SenderType = "AI"
	Human SenderType = "human"
)

type Chat struct {
	ID        uuid.UUID `json:"id"`
	CreatedBy User      `json:"created_by"`
}

func NewChat(createdBy User) Chat {
	return Chat{
		ID:        uuid.New(),
		CreatedBy: createdBy,
	}
}
