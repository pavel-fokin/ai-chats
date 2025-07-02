package domain

import (
	"github.com/google/uuid"

	"ai-chats/internal/pkg/types"
)

type Event interface {
	types.Message
	Channel() string
}

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

func (m MessageAdded) Channel() string {
	return m.ChatID.String()
}

// ChatTitleUpdated represents a title updated event.
type ChatTitleUpdated struct {
	ID     uuid.UUID `json:"id"`
	ChatID ChatID    `json:"chatId"`
	UserID UserID    `json:"userId"`
	Title  string    `json:"title"`
}

// NewChatTitleUpdated creates a new title updated event.
func NewChatTitleUpdated(chatID ChatID, userID UserID, title string) ChatTitleUpdated {
	return ChatTitleUpdated{
		ID:     uuid.New(),
		ChatID: chatID,
		UserID: userID,
		Title:  title,
	}
}

func (t ChatTitleUpdated) Type() types.MessageType {
	return ChatTitleUpdatedType
}

func (t ChatTitleUpdated) Channel() string {
	return t.UserID.String()
}

// ModelStreamedMessageEvent is an event that can be notified to the app when a model stream response is generated.
type ModelStreamedMessageEvent struct {
	ChatID ChatID `json:"chatId"`
	Text   string `json:"text"`
	Sender string `json:"sender"`
}

func (m ModelStreamedMessageEvent) Channel() string {
	return m.ChatID.String()
}

func (m ModelStreamedMessageEvent) Type() types.MessageType {
	return types.MessageType("ModelStreamMessageNotification")
}
