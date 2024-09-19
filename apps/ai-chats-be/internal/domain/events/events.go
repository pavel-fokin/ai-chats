package events

import (
	"github.com/google/uuid"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

const (
	MessageAddedType         types.MessageType = "messageAdded"
	ChatTitleUpdatedType     types.MessageType = "chatTitleUpdated"
	MessageChunkReceivedType types.MessageType = "messageChunkReceived"
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

func (m MessageAdded) Type() types.MessageType {
	return MessageAddedType
}

// ChatTitleUpdated represents a title updated event.
type ChatTitleUpdated struct {
	ID     uuid.UUID `json:"id"`
	ChatID uuid.UUID `json:"chatId"`
	Title  string    `json:"title"`
}

// NewTitleGenerated creates a new title generated event.
func NewChatTitleUpdated(chatID domain.ChatID, title string) ChatTitleUpdated {
	return ChatTitleUpdated{
		ID:     uuid.New(),
		ChatID: chatID,
		Title:  title,
	}
}

func (t ChatTitleUpdated) Type() types.MessageType {
	return ChatTitleUpdatedType
}

// MessageChunkReceived represents a message chunk received event.
type MessageChunkReceived struct {
	MessageID uuid.UUID `json:"messageId"`
	Sender    string    `json:"sender"`
	Text      string    `json:"text"`
	Final     bool      `json:"done"`
}

// NewMessageChunkReceived creates a new message chunk received event.
func NewMessageChunkReceived(messageID uuid.UUID, sender, text string, final bool) MessageChunkReceived {
	return MessageChunkReceived{
		MessageID: messageID,
		Sender:    sender,
		Text:      text,
		Final:     final,
	}
}

func (m MessageChunkReceived) Type() types.MessageType {
	return MessageChunkReceivedType
}
