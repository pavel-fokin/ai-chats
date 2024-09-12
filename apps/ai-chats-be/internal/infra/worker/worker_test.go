package worker

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type EventsMock struct {
	mock.Mock
}

func (m *EventsMock) Subscribe(ctx context.Context, topic string) (chan []byte, error) {
	args := m.Called(ctx, topic)
	return args.Get(0).(chan []byte), args.Error(1)
}

func (m *EventsMock) Unsubscribe(ctx context.Context, topic string, channel chan []byte) error {
	args := m.Called(ctx, topic, channel)
	return args.Error(0)
}

func TestWorker(t *testing.T) {
	t.Run("TestWorker", func(t *testing.T) {
		eventsMock := &EventsMock{}
		eventsMock.On("Subscribe", mock.Anything, mock.Anything).Return(make(chan []byte), nil)
		eventsMock.On("Unsubscribe", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		w := New(eventsMock)
		w.RegisterHandler("topic", 1, func(ctx context.Context, e []byte) error {
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
