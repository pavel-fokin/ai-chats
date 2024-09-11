package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Sender    Sender    `json:"sender"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewMessage(sender Sender, text string) Message {
	return Message{
		ID:        uuid.New(),
		Sender:    sender,
		Text:      text,
		CreatedAt: time.Now().UTC(),
	}
}

func NewUserMessage(user User, text string) Message {
	return NewMessage(
		NewUserSender(user.ID), text,
	)
}

func NewModelMessage(modelID ModelID, text string) Message {
	return NewMessage(
		NewModelSender(modelID), text,
	)
}

func (m Message) IsFromUser() bool {
	return m.Sender.IsUser()
}

func (m Message) IsFromModel() bool {
	return m.Sender.IsModel()
}
