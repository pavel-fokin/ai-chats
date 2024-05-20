package events

import (
	"context"
	"fmt"
)

type Topic = string

type Events struct {
	topics map[Topic]map[chan []byte]struct{}
}

func New() *Events {
	return &Events{
		topics: make(map[Topic]map[chan []byte]struct{}),
	}
}

func (e *Events) Subscribe(ctx context.Context, topic string) (chan []byte, error) {
	ch := make(chan []byte, 1)
	if _, ok := e.topics[topic]; !ok {
		e.topics[topic] = make(map[chan []byte]struct{})
	}
	e.topics[topic][ch] = struct{}{}
	return ch, nil
}

func (e *Events) Unsubscribe(ctx context.Context, topic string, ch chan []byte) error {
	close(ch)
	for range ch {
		// drain channel
	}
	if _, ok := e.topics[topic]; ok {
		delete(e.topics[topic], ch)
	}
	return nil
}

func (e *Events) Publish(ctx context.Context, topic string, data []byte) error {
	if _, ok := e.topics[topic]; !ok {
		return fmt.Errorf("topic %s not found", topic)
	}
	for ch := range e.topics[topic] {
		ch <- data
	}
	return nil
}