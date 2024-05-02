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

type Chat struct {
	ID       uuid.UUID
	Actors   []Actor
	Messages []Message
}
