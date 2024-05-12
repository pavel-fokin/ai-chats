package domain

import "github.com/google/uuid"

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
