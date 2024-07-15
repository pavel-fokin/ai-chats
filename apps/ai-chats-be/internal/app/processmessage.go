package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/json"
)

// ProcessAddedMessage processes a message added event.
func (a *App) ProcessAddedMessage(ctx context.Context, event events.MessageAdded) error {

	messageAdded := events.NewMessageAdded(event.ChatID, event.Message)
	if err := a.pubsub.Publish(ctx, event.ChatID.String(), json.MustMarshal(ctx, messageAdded)); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	switch {
	case event.Message.IsFromModel():
		// Ignore messages from models.
	case event.Message.IsFromUser():
		a.GenerateResponse(ctx, event.ChatID)
	default:
		return fmt.Errorf("unknown message type: %s", event.Message)
	}

	messages, err := a.chats.AllMessages(ctx, event.ChatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	if len(messages) == 2 {
		return a.GenerateTitle(ctx, event.ChatID)
	}

	return nil
}
