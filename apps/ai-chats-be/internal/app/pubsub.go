package app

import (
	"ai-chats/internal/pkg/types"
	"context"
)

const (
	GenerateChatTitleTopic TopicName = "generate-chat-title"
	MessageAddedTopic      TopicName = "message-added"
	PullOllamaModelTopic   TopicName = "pull-ollama-model"
)

type TopicName = string

// PubSub is a publish/subscribe interface.
type PubSub interface {
	Subscribe(context.Context, TopicName) (chan types.Message, error)
	Unsubscribe(context.Context, TopicName, chan types.Message) error
	Publish(context.Context, TopicName, types.Message) error
}

func (a *App) PublishEvents(ctx context.Context, events []types.Message) error {
	return nil
}
