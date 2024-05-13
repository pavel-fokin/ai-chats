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

func (c *Chatting) SendMessage(ctx context.Context, llm LLM, chatID uuid.UUID, message Message) (Message, error) {
	chat, err := c.chats.FindChat(ctx, chatID)
	if err != nil {
		return Message{}, err
	}

	if err := c.messages.Add(ctx, chat, message); err != nil {
		return Message{}, err
	}

	allMessages, err := c.messages.AllMessages(ctx, chat.ID)
	if err != nil {
		return Message{}, err
	}

	llmResponse, err := llm.GenerateResponse(ctx, allMessages)
	if err != nil {
		return Message{}, err
	}

	if err := c.messages.Add(ctx, chat, llmResponse); err != nil {
		return Message{}, err
	}

	return llmResponse, nil
}
