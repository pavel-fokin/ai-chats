package events

import (
	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

// MessageSent represents a message sent event.
type MessageAdded struct {
	ID      uuid.UUID      `json:"id"`
	ChatID  uuid.UUID      `json:"chat_id"`
	Message domain.Message `json:"message"`
}

func NewMessageAdded(chatID uuid.UUID, m domain.Message) MessageAdded {
	return MessageAdded{
		ID:      uuid.New(),
		ChatID:  chatID,
		Message: m,
	}
}

// TitleGenerated represents a title updated event.
type TitleUpdated struct {
	ID     uuid.UUID `json:"id"`
	ChatID uuid.UUID `json:"chat_id"`
	Title  string    `json:"title"`
}

// NewTitleGenerated creates a new title generated event.
func NewTitleUpdated(chatID uuid.UUID, title string) TitleUpdated {
	return TitleUpdated{
		ID:     uuid.New(),
		ChatID: chatID,
		Title:  title,
	}
}
