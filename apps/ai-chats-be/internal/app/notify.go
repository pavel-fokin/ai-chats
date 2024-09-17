package app

import (
	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
	"context"
	"fmt"
)

func (a *App) notifyInChat(ctx context.Context, chatID string, event events.Event) error {
	if err := a.pubsub.Publish(ctx, chatID, event); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}
	return nil
}

func (a *App) notifyApp(ctx context.Context, userID domain.UserID, event events.Event) error {
	if err := a.pubsub.Publish(ctx, userID.String(), event); err != nil {
		return fmt.Errorf("failed to publish an event: %w", err)
	}
	return nil
}
