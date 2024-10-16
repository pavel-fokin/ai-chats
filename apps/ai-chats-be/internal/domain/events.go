package domain

import (
	"github.com/google/uuid"

	"ai-chats/internal/pkg/types"
)

const (
	MessageAddedType         types.MessageType = "messageAdded"
	ChatTitleUpdatedType     types.MessageType = "chatTitleUpdated"
	MessageChunkReceivedType types.MessageType = "messageChunkReceived"
)

// MessageAdded represents a message sent event.
type MessageAdded struct {
	ID      uuid.UUID `json:"id"`
	ChatID  ChatID    `json:"chatId"`
	Message Message   `json:"message"`
}

func NewMessageAdded(chatID ChatID, m Message) MessageAdded {
	return MessageAdded{
		ID:      uuid.New(),
		ChatID:  chatID,
		Message: m,
	}
}

func (m MessageAdded) Type() types.MessageType {
	return MessageAddedType
}

// ChatTitleUpdated represents a title updated event.
type ChatTitleUpdated struct {
	ID     uuid.UUID `json:"id"`
	ChatID uuid.UUID `json:"chatId"`
	Title  string    `json:"title"`
}

// NewChatTitleUpdated creates a new title generated event.
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
