package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Sender    string    `json:"sender"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(sender string, text string) Message {
	return Message{
		ID:        uuid.New(),
		Sender:    sender,
		Text:      text,
		CreatedAt: time.Now().UTC(),
	}
}

func (m Message) IsFromUser() bool {
	return m.Sender == "User"
}

func (m Message) IsFromBot() bool {
	return m.Sender == "AI"
}
