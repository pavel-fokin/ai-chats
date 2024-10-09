package app

import (
	"ai-chats/internal/pkg/types"
	"context"
)

const (
	GenerateChatTitleTopic Topic = "generate-chat-title"
	MessageAddedTopic      Topic = "message-added"
	PullOllamaModelTopic   Topic = "pull-ollama-model"
)

type Topic = string

// PubSub is a publish/subscribe interface.
type PubSub interface {
	Subscribe(context.Context, Topic) (chan types.Message, error)
	Unsubscribe(context.Context, Topic, chan types.Message) error
	Publish(context.Context, Topic, types.Message) error
}
