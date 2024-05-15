package worker

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type App interface {
	GenerateResponse(ctx context.Context, chatID uuid.UUID) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string, subsctiber string) (<-chan GenerateResponse, error)
	Unsubscribe(ctx context.Context, topic string, subsctiber string) error
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
	events, err := w.events.Subscribe(context.Background(), "generate-response", "worker")
	if err != nil {
		slog.ErrorContext(w.ctx, "failed to subscribe to events", "err", err)
		return
	}
	defer w.events.Unsubscribe(context.Background(), "generate-response", "worker")

	for {
		select {
		case <-w.ctx.Done():
			return
		case e := <-events:
			err := w.app.GenerateResponse(w.ctx, e.ChatID)
			if err != nil {
				slog.ErrorContext(w.ctx, "failed to generate a response", "err", err)
			}
		}
	}
}

func (w *Worker) Stop() {
	w.ctx.Done()
}
