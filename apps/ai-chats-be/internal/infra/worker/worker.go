package worker

import (
	"context"
	"fmt"
	"log/slog"

	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
)

type App interface {
	GenerateTitle(ctx context.Context, chatID domain.ChatID) error
	ProcessAddedMessage(ctx context.Context, event events.MessageAdded) error
	PullOllamaModel(ctx context.Context, model string) error
}

type PubSub interface {
	Subscribe(ctx context.Context, topic string) (chan events.Event, error)
	Unsubscribe(ctx context.Context, topic string, channel chan events.Event) error
}

// type Handler interface {
// 	Handle(ctx context.Context, events Events, topic string, concurrency int, handler HandlerFunc) error
// }

type HandlerFunc func(ctx context.Context, event events.Event) error

func (hf HandlerFunc) Handle(ctx context.Context, pubsub PubSub, topic string, concurrency int) error {
	channel, err := pubsub.Subscribe(ctx, topic)
	if err != nil {
		slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
		return fmt.Errorf("failed to subscribe to events: %w", err)
	}
	defer pubsub.Unsubscribe(ctx, topic, channel)

	for {
		select {
		case <-ctx.Done():
			return nil
		case e := <-channel:
			if err := hf(ctx, e); err != nil {
				slog.ErrorContext(ctx, "failed to handle event", "err", err)
			}
		}
	}
}

type handler struct {
	concurrency int
	fn          HandlerFunc
}

type Topic = string

type Worker struct {
	pubsub   PubSub
	ctx      context.Context
	stop     context.CancelFunc
	handlers map[Topic]handler
}

func New(pubsub PubSub) *Worker {
	ctx, stop := context.WithCancel(context.Background())

	return &Worker{
		pubsub:   pubsub,
		ctx:      ctx,
		stop:     stop,
		handlers: make(map[Topic]handler),
	}
}

func (w *Worker) RegisterHandler(topic string, concurrency int, fn HandlerFunc) {
	w.handlers[topic] = handler{concurrency: concurrency, fn: fn}
}

func (w *Worker) Start() {
	for topic, handler := range w.handlers {
		go handler.fn.Handle(w.ctx, w.pubsub, topic, handler.concurrency)
	}
}

func (w *Worker) Shutdown() {
	w.stop()
}
