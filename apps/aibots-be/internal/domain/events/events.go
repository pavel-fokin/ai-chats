package events

import (
	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type EventType string

const (
	MessageAddedType         EventType = "message_added"
	ChatTitleUpdatedType     EventType = "chat_title_updated"
	MessageChunkReceivedType EventType = "message_chunk_received"
)

// MessageSent represents a message sent event.
type MessageAdded struct {
	ID      uuid.UUID      `json:"id"`
	Type    EventType      `json:"type"`
	ChatID  uuid.UUID      `json:"chat_id"`
	Message domain.Message `json:"message"`
}

func NewMessageAdded(chatID uuid.UUID, m domain.Message) MessageAdded {
	return MessageAdded{
		ID:      uuid.New(),
		Type:    "message_added",
		ChatID:  chatID,
		Message: m,
	}
}

// TitleGenerated represents a title updated event.
type ChatTitleUpdated struct {
	ID     uuid.UUID `json:"id"`
	Type   EventType `json:"type"`
	ChatID uuid.UUID `json:"chat_id"`
	Title  string    `json:"title"`
}

// NewTitleGenerated creates a new title generated event.
func NewChatTitleUpdated(chatID uuid.UUID, title string) ChatTitleUpdated {
	return ChatTitleUpdated{
		ID:     uuid.New(),
		Type:   "chat_title_updated",
		ChatID: chatID,
		Title:  title,
	}
}

// MessageChunkReceived represents a message chunk received event.
type MessageChunkReceived struct {
	MessageID uuid.UUID `json:"message_id"`
	Type      EventType `json:"type"`
	Sender    string    `json:"sender"`
	Text      string    `json:"text"`
	Final     bool      `json:"done"`
}

// NewMessageChunkReceived creates a new message chunk received event.
func NewMessageChunkReceived(messageID uuid.UUID, sender, text string, final bool) MessageChunkReceived {
	return MessageChunkReceived{
		MessageID: messageID,
		Type:      "message_chunk_received",
		Sender:    sender,
		Text:      text,
		Final:     final,
	}
}
