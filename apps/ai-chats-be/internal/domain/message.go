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

func NewUserMessage(user User, messageText string) Message {
	return NewMessage(
		NewUserSender(user.ID), messageText,
	)
}

func NewModelMessage(modelID ModelID, messageText string) Message {
	return NewMessage(
		NewModelSender(modelID), messageText,
	)
}

func NewSystemMessage(messageText string) Message {
	return NewMessage(NewSystemSender(), messageText)
}

func (m Message) IsFromUser() bool {
	return m.Sender.IsUser()
}

func (m Message) IsFromModel() bool {
	return m.Sender.IsModel()
}

func (m Message) IsFromSystem() bool {
	return m.Sender.IsSystem()
}
