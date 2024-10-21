package domain

import (
	"github.com/google/uuid"

	"ai-chats/internal/pkg/types"
)

const (
	MessageAddedType     types.MessageType = "messageAdded"
	ChatTitleUpdatedType types.MessageType = "chatTitleUpdated"
)

// MessageAdded represents a message added event.
type MessageAdded struct {
	ID      uuid.UUID `json:"id"`
	ChatID  ChatID    `json:"chatId"`
	Message Message   `json:"message"`
}

func NewMessageAdded(chatID ChatID, message Message) MessageAdded {
	return MessageAdded{
		ID:      uuid.New(),
		ChatID:  chatID,
		Message: message,
	}
}

func (m MessageAdded) Type() types.MessageType {
	return MessageAddedType
}

// ChatTitleUpdated represents a title updated event.
type ChatTitleUpdated struct {
	ID     uuid.UUID `json:"id"`
	ChatID ChatID    `json:"chatId"`
	Title  string    `json:"title"`
}

// NewChatTitleUpdated creates a new title updated event.
func NewChatTitleUpdated(chatID ChatID, title string) ChatTitleUpdated {
	return ChatTitleUpdated{
		ID:     uuid.New(),
		ChatID: chatID,
		Title:  title,
	}
}

func (t ChatTitleUpdated) Type() types.MessageType {
	return ChatTitleUpdatedType
}
