package pubsub

import (
	"ai-chats/internal/pkg/types"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Event1 struct{}

func (e Event1) Type() types.MessageType {
	return types.MessageType("test")
}

func TestPubSub(t *testing.T) {
	t.Run("subscribe and publish", func(t *testing.T) {
		ctx := context.Background()
		ps := New()

		ch, err := ps.Subscribe(ctx, "test")
		defer ps.Unsubscribe(ctx, "test", ch)
		assert.NoError(t, err)
		assert.NotNil(t, ch)

		err = ps.Publish(ctx, "test", Event1{})
		assert.NoError(t, err)

		msg := <-ch
		assert.Equal(t, Event1{}, msg)
	})

	t.Run("publish and subscribe", func(t *testing.T) {
		ctx := context.Background()
		ps := New()

		err := ps.Publish(ctx, "test", Event1{})
		assert.NoError(t, err)

		ch, err := ps.Subscribe(ctx, "test")
		defer ps.Unsubscribe(ctx, "test", ch)
		assert.NoError(t, err)

		msg := <-ch
		assert.Equal(t, Event1{}, msg)
	})
}
