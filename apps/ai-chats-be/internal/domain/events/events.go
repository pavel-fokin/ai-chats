package events

import (
	"github.com/google/uuid"

	"ai-chats/internal/domain"
)

type EventType string

const (
	MessageAddedType         EventType = "messageAdded"
	ChatTitleUpdatedType     EventType = "chatTitleUpdated"
	MessageChunkReceivedType EventType = "messageChunkReceived"
)

// MessageSent represents a message sent event.
type MessageAdded struct {
	ID      uuid.UUID      `json:"id"`
	ChatID  domain.ChatID  `json:"chatId"`
	Message domain.Message `json:"message"`
}

func NewMessageAdded(chatID uuid.UUID, m domain.Message) MessageAdded {
	return MessageAdded{
		ID:      uuid.New(),
		ChatID:  chatID,
		Message: m,
	}
}

func (m MessageAdded) Type() EventType {
	return MessageAddedType
}

// ChatTitleUpdated represents a title updated event.
type ChatTitleUpdated struct {
	ID uuid.UUID `json:"id"`
	// Type   EventType `json:"type"`
	ChatID uuid.UUID `json:"chatId"`
	Title  string    `json:"title"`
}

// NewTitleGenerated creates a new title generated event.
func NewChatTitleUpdated(chatID domain.ChatID, title string) ChatTitleUpdated {
	return ChatTitleUpdated{
		ID: uuid.New(),
		// Type:   ChatTitleUpdatedType,
		ChatID: chatID,
		Title:  title,
	}
}

func (t ChatTitleUpdated) Type() EventType {
	return ChatTitleUpdatedType
}

// MessageChunkReceived represents a message chunk received event.
type MessageChunkReceived struct {
	MessageID uuid.UUID `json:"messageId"`
	// Type      EventType `json:"type"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Final  bool   `json:"done"`
}

// NewMessageChunkReceived creates a new message chunk received event.
func NewMessageChunkReceived(messageID uuid.UUID, sender, text string, final bool) MessageChunkReceived {
	return MessageChunkReceived{
		MessageID: messageID,
		// Type:      MessageChunkReceivedType,
		Sender: sender,
		Text:   text,
		Final:  final,
	}
}

func (m MessageChunkReceived) Type() EventType {
	return MessageChunkReceivedType
}
