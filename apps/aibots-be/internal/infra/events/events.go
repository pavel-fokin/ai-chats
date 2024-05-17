package events

import "context"

type Topic = string

type Channel interface {
	ID() string
	Topic() Topic
	C() chan []byte
}

type Events interface {
	Subscribe(context.Context, Topic) (Channel, error)
	Unsubscribe(context.Context, Channel) error
	Publish(context.Context, Topic, []byte) error
}
