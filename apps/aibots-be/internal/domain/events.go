package domain

import (
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

// TitleGenerated represents a title updated event.
type TitleGenerated struct {
	ID     uuid.UUID `json:"id"`
	ChatID uuid.UUID `json:"chat_id"`
	Title  string    `json:"title"`
}

// NewTitleGenerated creates a new title generated event.
func NewTitleGenerated(chatID uuid.UUID, title string) TitleGenerated {
	return TitleGenerated{
		ID:     uuid.New(),
		ChatID: chatID,
		Title:  title,
	}
}
