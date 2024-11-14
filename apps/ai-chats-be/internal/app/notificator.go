package app

import (
	"context"

	"ai-chats/internal/app/notifications"
)

type notificator struct {
	pubsub PubSub
}

func NewNotificator(pubsub PubSub) *notificator {
	return &notificator{pubsub: pubsub}
}

func (n *notificator) Notify(ctx context.Context, notification notifications.Notification) error {
	return n.pubsub.Publish(ctx, notification.Channel(), notification)
}
