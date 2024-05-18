package worker

import (
	"context"
	"encoding/json"
	"log/slog"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"github.com/google/uuid"
)

type App interface {
	GenerateResponse(ctx context.Context, chatID uuid.UUID) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (chan []byte, error)
	Unsubscribe(ctx context.Context, topic string, channel chan []byte) error
}

type Worker struct {
	app    App
	events Subscriber
	ctx    context.Context
}

func New(app App, events Subscriber) *Worker {
	return &Worker{
		app:    app,
		events: events,
		ctx:    context.Background(),
	}
}

func (w *Worker) Start() {
	events, err := w.events.Subscribe(context.Background(), "worker")
	if err != nil {
		slog.ErrorContext(w.ctx, "failed to subscribe to events", "err", err)
		return
	}
	defer w.events.Unsubscribe(context.Background(), "worker", events)

	for {
		select {
		case <-w.ctx.Done():
			return
		case e := <-events:
			var messageSent domain.MessageSent
			if err := json.Unmarshal(e, &messageSent); err != nil {
				slog.ErrorContext(w.ctx, "failed to unmarshal event", "err", err)
				continue
			}

			err := w.app.GenerateResponse(w.ctx, messageSent.ChatID)
			if err != nil {
				slog.ErrorContext(w.ctx, "failed to generate a response", "err", err)
			}
		}
	}
}

func (w *Worker) Stop() {
	w.ctx.Done()
}
