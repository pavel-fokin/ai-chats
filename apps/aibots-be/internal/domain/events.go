package domain

import (
	"encoding/json"

	"github.com/google/uuid"
)

// MessageSent represents a message sent event.
type MessageSent struct {
	ID      uuid.UUID `json:"id"`
	ChatID  uuid.UUID `json:"chat_id"`
	Message Message   `json:"message"`
}

func NewMessageSent(chatID uuid.UUID, m Message) MessageSent {
	return MessageSent{
		ID:      uuid.New(),
		ChatID:  chatID,
		Message: m,
	}
}

// AsBytes returns the byte representation of the event.
func (m MessageSent) AsBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
