package app

import (
	"context"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/events"
	"pavel-fokin/ai/apps/ai-bots-be/internal/worker"
)

type MessageSender interface {
	SendMessage(ctx context.Context, chatID uuid.UUID, message domain.Message) (domain.MessageSent, error)
}

type EventsPublisher[Event any] interface {
	Publish(context.Context, events.Topic, Event) error
}

type EventsSubscriber[T any] interface {
	Subscribe(context.Context, events.Topic, events.Subscriber) (<-chan T, error)
	Unsubscribe(context.Context, events.Topic, events.Subscriber) error
}

type Events[T any] interface {
	EventsPublisher[T]
	EventsSubscriber[T]
}

type App struct {
	users                  domain.Users
	chats                  domain.Chats
	messages               domain.Messages
	chatting               MessageSender
	messageSentEvents      Events[domain.MessageSent]
	generateResponseEvents EventsPublisher[worker.GenerateResponse]
}

func New(
	chats domain.Chats,
	users domain.Users,
	messages domain.Messages,
	generateResponseEvents Events[worker.GenerateResponse],
) *App {
	chatting := domain.NewChatting(chats, messages)
	messageSentEvents := events.New[domain.MessageSent]()
	// genResponseEvents := events.New[worker.GenerateResponse]()

	return &App{
		chats:                  chats,
		users:                  users,
		messages:               messages,
		chatting:               chatting,
		messageSentEvents:      messageSentEvents,
		generateResponseEvents: generateResponseEvents,
	}
}
