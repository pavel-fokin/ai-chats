package worker

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type App interface {
	GenerateResponse(ctx context.Context, chatID uuid.UUID) error
}

type Events interface {
	Subscribe(ctx context.Context, topic string) (chan []byte, error)
	Unsubscribe(ctx context.Context, topic string, channel chan []byte) error
}

// type Handler interface {
// 	Handle(ctx context.Context, events Events, topic string, concurrency int, handler HandlerFunc) error
// }

type HandlerFunc func(ctx context.Context, event []byte) error

func (hf HandlerFunc) Handle(ctx context.Context, events Events, topic string, concurrency int) error {
	channel, err := events.Subscribe(ctx, topic)
	if err != nil {
		slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
		return fmt.Errorf("failed to subscribe to events: %w", err)
	}
	defer events.Unsubscribe(ctx, topic, channel)

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
	events   Events
	ctx      context.Context
	stop     context.CancelFunc
	handlers map[Topic]handler
}

func New(events Events) *Worker {
	ctx, stop := context.WithCancel(context.Background())

	return &Worker{
		events:   events,
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
		go handler.fn.Handle(w.ctx, w.events, topic, handler.concurrency)
	}
}

func (w *Worker) Stop() {
	w.stop()
}
