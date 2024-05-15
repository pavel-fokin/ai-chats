package events

import (
	"context"
	"fmt"
)

type Topic = string
type Subscriber = string

type Events[T any] struct {
	topics map[Topic]map[Subscriber]chan T
}

func New[T any]() *Events[T] {
	return &Events[T]{
		topics: make(map[Topic]map[Subscriber]chan T),
	}
}

func (e *Events[T]) Subscribe(ctx context.Context, topic Topic, subscriber Subscriber) (<-chan T, error) {
	if e.topics[topic] == nil {
		e.topics[topic] = make(map[Subscriber]chan T)
	}

	ch := make(chan T, 1)
	e.topics[topic][subscriber] = ch

	return ch, nil
}

func (e *Events[T]) Unsubscribe(ctx context.Context, topic Topic, subscriber Subscriber) error {
	if e.topics[topic] == nil {
		return nil
	}
	close(e.topics[topic][subscriber])
	for range e.topics[topic][subscriber] {
		// drain the channel
	}

	delete(e.topics[topic], subscriber)

	return nil
}

func (e *Events[T]) Publish(ctx context.Context, topic Topic, event T) error {
	if e.topics[topic] == nil {
		return fmt.Errorf("topic %s not found", topic)
	}

	for _, ch := range e.topics[topic] {
		ch <- event
	}

	return nil
}
