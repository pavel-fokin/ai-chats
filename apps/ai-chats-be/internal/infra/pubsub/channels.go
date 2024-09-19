package pubsub

import (
	"context"

	"ai-chats/internal/pkg/types"
)

type Topic = string

type Events struct {
	topics map[Topic]map[chan types.Message]struct{}
}

func New() *Events {
	return &Events{
		topics: make(map[Topic]map[chan types.Message]struct{}),
	}
}

func (e *Events) Subscribe(ctx context.Context, topic string) (chan types.Message, error) {
	ch := make(chan types.Message, 1)
	if _, ok := e.topics[topic]; !ok {
		e.topics[topic] = make(map[chan types.Message]struct{})
	}
	e.topics[topic][ch] = struct{}{}
	return ch, nil
}

func (e *Events) Unsubscribe(ctx context.Context, topic string, ch chan types.Message) error {
	close(ch)
	for range ch {
		// drain channel
	}
	if _, ok := e.topics[topic]; ok {
		delete(e.topics[topic], ch)
	}
	return nil
}

func (e *Events) Publish(ctx context.Context, topic string, event types.Message) error {
	if _, ok := e.topics[topic]; !ok {
		e.topics[topic] = make(map[chan types.Message]struct{})
		// return fmt.Errorf("topic %s not found", topic)
	}
	for ch := range e.topics[topic] {
		ch <- event
	}
	return nil
}

func (e *Events) CloseAll() {
	for topic, chs := range e.topics {
		for ch := range chs {
			close(ch)
			for range ch {
				// drain channel
			}
		}
		delete(e.topics, topic)
	}
}
