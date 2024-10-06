package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

func (a *App) notifyInChat(ctx context.Context, chatID string, event types.Message) error {
	if err := a.pubsub.Publish(ctx, chatID, event); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}
	return nil
}

func (a *App) notifyApp(ctx context.Context, userID domain.UserID, event types.Message) error {
	if err := a.pubsub.Publish(ctx, userID.String(), event); err != nil {
		return fmt.Errorf("failed to publish an event: %w", err)
	}
	return nil
}

func (a *App) notifyOllamaModelPulling(ctx context.Context, model string, event types.Message) error {
	if err := a.pubsub.Publish(ctx, model, event); err != nil {
		return fmt.Errorf("failed to publish an event: %w", err)
	}
	return nil
}
