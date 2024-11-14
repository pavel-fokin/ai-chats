package notifications

import (
	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

// Notification is an event that can be notified to the app.
type Notification interface {
	types.Message
	Channel() string
}

// MessageAdded is an event that can be notified to the app when a message is added to a chat.
type MessageAdded struct {
	ChatID domain.ChatID `json:"chatId"`
}

func NewMessageAdded(chatID domain.ChatID) MessageAdded {
	return MessageAdded{ChatID: chatID}
}

func (m MessageAdded) Channel() string {
	return m.ChatID.String()
}

func (m MessageAdded) Type() types.MessageType {
	return types.MessageType("MessageAddedNotification")
}

// ModelStreamMessage is an event that can be notified to the app when a model stream response is generated.
type ModelStreamMessage struct {
	ChatID domain.ChatID `json:"chatId"`
	Text   string        `json:"text"`
	Sender string        `json:"sender"`
}

func NewModelStreamMessage(chatID domain.ChatID, text, sender string) ModelStreamMessage {
	return ModelStreamMessage{ChatID: chatID, Text: text, Sender: sender}
}

func (m ModelStreamMessage) Channel() string {
	return m.ChatID.String()
}

func (m ModelStreamMessage) Type() types.MessageType {
	return types.MessageType("ModelStreamMessageNotification")
}

// ChatTitleUpdated is an event that can be notified to the app when a chat title is updated.
type ChatTitleUpdated struct {
	ChatID domain.ChatID `json:"chatId"`
	UserID domain.UserID `json:"userId"`
}

func NewChatTitleUpdated(chatID domain.ChatID, userID domain.UserID) ChatTitleUpdated {
	return ChatTitleUpdated{ChatID: chatID, UserID: userID}
}

func (c ChatTitleUpdated) Channel() string {
	return c.UserID.String()
}

func (c ChatTitleUpdated) Type() types.MessageType {
	return types.MessageType("ChatTitleUpdatedNotification")
}
