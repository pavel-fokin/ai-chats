package domain

import "github.com/google/uuid"

type ActorType string

const (
	AI    ActorType = "ai"
	Human ActorType = "human"
)

type Actor struct {
	ID   uuid.UUID
	Type ActorType
}

type Message struct {
	ID    uuid.UUID
	Actor Actor
	Text  string
}

func NewMessage(id uuid.UUID, actor Actor, text string) Message {
	return Message{
		ID:    id,
		Actor: actor,
		Text:  text,
	}
}

type Chat struct {
	ID       uuid.UUID
	Actors   []Actor
	Messages []Message
}

func NewChat(id uuid.UUID, actors []Actor) *Chat {
	return &Chat{
		ID:     id,
		Actors: actors,
	}
}

func (c *Chat) AddMessage(message Message) {
	c.Messages = append(c.Messages, message)
}
