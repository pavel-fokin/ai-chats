package domain

import (
	"context"

	"github.com/google/uuid"
)

type LLM interface {
	GenerateResponse(ctx context.Context, messages []Message) (Message, error)
}

type Chatting struct {
	chats    Chats
	messages Messages
}

func NewChatting(chats Chats, messages Messages) *Chatting {
	return &Chatting{
		chats:    chats,
		messages: messages,
	}
}

func (c *Chatting) SendMessage(ctx context.Context, chatID uuid.UUID, message Message) (MessageSent, error) {
	chat, err := c.chats.FindChat(ctx, chatID)
	if err != nil {
		return MessageSent{}, err
	}

	if err := c.messages.Add(ctx, chat, message); err != nil {
		return MessageSent{}, err
	}

	return NewMessageSent(chatID, message), nil
}
