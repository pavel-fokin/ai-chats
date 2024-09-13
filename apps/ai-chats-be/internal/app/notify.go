package app

import (
	"ai-chats/internal/domain"
	"context"
	"fmt"
)

func (a *App) notifyInChat(ctx context.Context, chatID string, event []byte) error {
	if err := a.pubsub.Publish(ctx, chatID, event); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}
	return nil
}

func (a *App) notifyApp(ctx context.Context, userID domain.UserID, event []byte) error {
	fmt.Println("notifyApp", userID.String(), string(event))
	if err := a.pubsub.Publish(ctx, userID.String(), event); err != nil {
		return fmt.Errorf("failed to publish an event: %w", err)
	}
	return nil
}
