package app

import (
	"context"
	"fmt"
	"strings"

	"ai-chats/internal/domain"
)

// Chats is an service for managing chats.
type Chats struct {
	chats  domain.Chats
	users  domain.Users
	pubsub PubSub
	tx     Tx
}

// NewChats creates a new chats service.
func NewChats(chats domain.Chats, users domain.Users, pubsub PubSub, tx Tx) *Chats {
	return &Chats{chats: chats, users: users, pubsub: pubsub, tx: tx}
}

// CreateChat creates a chat for the user.
func (c *Chats) CreateChat(ctx context.Context, userID domain.UserID, model, messageText string) (domain.Chat, error) {
	messageText = strings.TrimSpace(messageText)
	if messageText == "" {
		return domain.Chat{}, fmt.Errorf("message text is empty")
	}

	user, err := c.users.FindByID(ctx, userID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("error finding user: %w", err)
	}

	var chat domain.Chat
	if err := c.tx.Tx(ctx, func(ctx context.Context) error {
		chat = domain.NewChat(user, domain.NewModelID(model))
		chat.AddMessage(domain.NewUserMessage(user, messageText))

		if err := c.chats.Add(ctx, chat); err != nil {
			return fmt.Errorf("error adding a chat: %w", err)
		}

		return nil
	}); err != nil {
		return domain.Chat{}, fmt.Errorf("error creating chat: %w", err)
	}

	for _, event := range chat.Events {
		if err := c.pubsub.Publish(ctx, MessageAddedTopic, event); err != nil {
			return domain.Chat{}, fmt.Errorf("error publishing chat events: %w", err)
		}
	}

	return chat, nil
}

// FindChatsByUserID returns all chats for the user.
func (c *Chats) FindChatsByUserID(ctx context.Context, userID domain.UserID) ([]domain.Chat, error) {
	return c.chats.FindByUserID(ctx, userID)
}

// FindChatByIDWithMessages returns a chat with messages.
func (c *Chats) FindChatByIDWithMessages(ctx context.Context, userID domain.UserID, chatID domain.ChatID) (domain.Chat, error) {
	chat, err := c.chats.FindByIDWithMessages(ctx, chatID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("error finding chat: %w", err)
	}

	if err := chat.CanUserAccess(userID); err != nil {
		return domain.Chat{}, fmt.Errorf("error checking chat access: %w", err)
	}

	return chat, nil
}

// FindChatByID finds a chat by ID.
func (c *Chats) FindChatByID(ctx context.Context, userID domain.UserID, chatID domain.ChatID) (domain.Chat, error) {
	chat, err := c.chats.FindByID(ctx, chatID)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("error finding chat: %w", err)
	}

	if err := chat.CanUserAccess(userID); err != nil {
		return domain.Chat{}, fmt.Errorf("error checking chat access: %w", err)
	}

	return chat, nil
}

// DeleteChat deletes the chat.
func (c *Chats) DeleteChat(ctx context.Context, userID domain.UserID, chatID domain.ChatID) error {
	chat, err := c.chats.FindByID(ctx, chatID)
	if err != nil {
		return fmt.Errorf("error finding chat: %w", err)
	}

	if err := chat.CanUserAccess(userID); err != nil {
		return fmt.Errorf("error checking chat access: %w", err)
	}

	return c.chats.Delete(ctx, chatID)
}

// SendMessage sends a message to the chat.
func (c *Chats) SendMessage(
	ctx context.Context, userID domain.UserID, chatID domain.ChatID, messageText string,
) error {
	user, err := c.users.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to find a user: %w", err)
	}

	var chat domain.Chat
	if err := c.tx.Tx(ctx, func(ctx context.Context) error {
		chat, err = c.chats.FindByID(ctx, chatID)
		if err != nil {
			return fmt.Errorf("error finding chat: %w", err)
		}

		if err := chat.CanUserAccess(userID); err != nil {
			return fmt.Errorf("error checking chat access: %w", err)
		}

		chat.AddMessage(domain.NewUserMessage(user, messageText))
		if err := c.chats.Update(ctx, chat); err != nil {
			return fmt.Errorf("error updating chat: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	for _, event := range chat.Events {
		if err := c.pubsub.Publish(ctx, MessageAddedTopic, event); err != nil {
			return fmt.Errorf("error publishing events: %w", err)
		}
	}

	return nil
}
