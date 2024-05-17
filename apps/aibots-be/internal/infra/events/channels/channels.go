package channels

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/events"
)

type TopicName = string
type ChannelID = string

type Channel struct {
	id        string
	topicName TopicName
	c         chan []byte
}

func NewChannel(topic TopicName) Channel {
	return Channel{
		id:        uuid.New().String(),
		topicName: topic,
		c:         make(chan []byte, 1),
	}
}

func (c *Channel) ID() ChannelID {
	return c.id
}

func (c *Channel) Topic() TopicName {
	return c.topicName
}

func (c *Channel) C() chan []byte {
	return c.c
}

type Topic struct {
	name     TopicName
	channels map[ChannelID]Channel
}

func NewTopic(name string) *Topic {
	return &Topic{
		name:     name,
		channels: make(map[ChannelID]Channel),
	}
}

func (t *Topic) Subscribe() (events.Channel, error) {
	channel := NewChannel(t.name)
	t.channels[channel.ID()] = channel
	return &channel, nil
}

func (t *Topic) Publish(event []byte) error {
	for _, channel := range t.channels {
		channel.c <- event
	}
	return nil
}

func (t *Topic) Unsubscribe(channel events.Channel) error {
	close(channel.C())
	delete(t.channels, channel.ID())
	return nil
}

type Events struct {
	sync.RWMutex
	topics map[TopicName]Topic
}

func New() *Events {
	return &Events{
		topics: make(map[TopicName]Topic),
	}
}

func (e *Events) Subscribe(ctx context.Context, topicName TopicName) (events.Channel, error) {
	e.Lock()
	defer e.Unlock()

	topic, ok := e.topics[topicName]
	if !ok {
		topic = *NewTopic(topicName)
		e.topics[topicName] = topic
	}

	return topic.Subscribe()
}

func (e *Events) Publish(ctx context.Context, topicName TopicName, event []byte) error {
	e.RLock()
	defer e.RUnlock()

	topic, ok := e.topics[topicName]
	if !ok {
		return fmt.Errorf("topic %s not found", topicName)
	}

	return topic.Publish(event)
}

func (e *Events) Unsubscribe(ctx context.Context, channel events.Channel) error {
	e.Lock()
	defer e.Unlock()

	topic, ok := e.topics[(channel.Topic())]
	if !ok {
		return nil
	}

	return topic.Unsubscribe(channel)
}
