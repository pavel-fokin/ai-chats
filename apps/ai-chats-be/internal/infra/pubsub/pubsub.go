package pubsub

import (
	"context"
	"sync"

	"ai-chats/internal/pkg/types"
)

type TopicName = string

type Topic struct {
	messages    []types.Message
	subscribers map[chan types.Message]struct{}
}

type PubSub struct {
	mu     sync.RWMutex
	topics map[TopicName]*Topic
}

func New() *PubSub {
	return &PubSub{
		topics: make(map[TopicName]*Topic),
	}
}

func (ps *PubSub) Subscribe(ctx context.Context, topicName TopicName) (chan types.Message, error) {
	ch := make(chan types.Message, 1)
	ps.mu.Lock()
	topic, exists := ps.topics[topicName]
	if !exists {
		topic = &Topic{
			subscribers: make(map[chan types.Message]struct{}),
			messages:    []types.Message{},
		}
		ps.topics[topicName] = topic
	}
	topic.subscribers[ch] = struct{}{}
	messages := make([]types.Message, len(topic.messages))
	copy(messages, topic.messages)
	ps.mu.Unlock()

	// Deliver stored messages to the new subscriber.
	go func() {
		for _, msg := range messages {
			select {
			case ch <- msg:
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch, nil
}

func (ps *PubSub) Publish(ctx context.Context, topicName TopicName, message types.Message) error {
	ps.mu.Lock()
	topic, exists := ps.topics[topicName]
	if !exists {
		topic = &Topic{
			subscribers: make(map[chan types.Message]struct{}),
			messages:    []types.Message{},
		}
		ps.topics[topicName] = topic
	}
	topic.messages = append(topic.messages, message)
	subscribers := make([]chan types.Message, 0, len(topic.subscribers))
	for ch := range topic.subscribers {
		subscribers = append(subscribers, ch)
	}
	ps.mu.Unlock()

	for _, ch := range subscribers {
		select {
		case ch <- message:
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Optionally handle full channels.
		}
	}

	return nil
}

func (ps *PubSub) Unsubscribe(ctx context.Context, topicName TopicName, ch chan types.Message) error {
	ps.mu.Lock()
	if topic, ok := ps.topics[topicName]; ok {
		delete(topic.subscribers, ch)
		if len(topic.subscribers) == 0 && len(topic.messages) == 0 {
			delete(ps.topics, topicName)
		}
	}
	ps.mu.Unlock()

	close(ch)
	return nil
}

func (ps *PubSub) Close() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for topicName, topic := range ps.topics {
		for ch := range topic.subscribers {
			close(ch)
		}
		delete(ps.topics, topicName)
	}

	return nil
}
