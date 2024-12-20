package worker

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ai-chats/internal/pkg/types"
)

type EventsMock struct {
	mock.Mock
}

func (m *EventsMock) Subscribe(ctx context.Context, topic string) (chan types.Message, error) {
	args := m.Called(ctx, topic)
	return args.Get(0).(chan types.Message), args.Error(1)
}

func (m *EventsMock) Unsubscribe(ctx context.Context, topic string, channel chan types.Message) error {
	args := m.Called(ctx, topic, channel)
	return args.Error(0)
}

func TestWorker(t *testing.T) {
	t.Run("start and shutdown", func(t *testing.T) {
		eventsMock := &EventsMock{}
		eventsMock.On("Subscribe", mock.Anything, mock.Anything).Return(make(chan types.Message), nil)
		eventsMock.On("Unsubscribe", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		w := New(eventsMock)
		w.RegisterHandler("topic", 1, func(ctx context.Context, e types.Message) error {
			return nil
		})

		// Get the number of goroutines before starting the worker.
		numGoroutins := runtime.NumGoroutine()
		w.Start()
		time.Sleep(time.Millisecond)

		// Assert that the number of goroutines has increased by one.
		assert.Equal(t, numGoroutins+1, runtime.NumGoroutine())

		w.Shutdown()
		time.Sleep(time.Millisecond)

		// Assert that the number of goroutines is the same as before the start of the worker.
		assert.Equal(t, numGoroutins, runtime.NumGoroutine())
	})
}
