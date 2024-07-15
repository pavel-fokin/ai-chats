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

func NewUserMessage(user User, text string) Message {
	return NewMessage("User", text)
}

func NewModelMessage(text string) Message {
	return NewMessage("AI", text)
}

func (m Message) IsFromUser() bool {
	return m.Sender == "User"
}

func (m Message) IsFromModel() bool {
	return m.Sender == "AI"
}
