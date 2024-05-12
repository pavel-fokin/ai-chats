package domain

import "github.com/google/uuid"

type SenderType string

const (
	AI    SenderType = "AI"
	Human SenderType = "human"
)

type Message struct {
	ID     uuid.UUID `json:"id"`
	Sender string    `json:"sender"`
	Text   string    `json:"text"`
}

func NewMessage(sender string, text string) Message {
	return Message{
		ID:     uuid.New(),
		Sender: sender,
		Text:   text,
	}
}

type Chat struct {
	ID uuid.UUID `json:"id"`
}

func NewChat() Chat {
	return Chat{
		ID: uuid.New(),
	}
}
