package app

import (
	"context"
)

// Subscribe subscribes to the chat events.
func (a *App) Subscribe(ctx context.Context, topic string) (chan []byte, error) {
	return a.events.Subscribe(ctx, topic)
}

// Unsubscribe unsubscribes from the chat events.
func (a *App) Unsubscribe(ctx context.Context, topic string, channel chan []byte) error {
	return a.events.Unsubscribe(ctx, topic, channel)
}
