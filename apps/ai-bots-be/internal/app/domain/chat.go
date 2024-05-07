package domain

import "github.com/google/uuid"

type ActorType string

const (
	AI    ActorType = "ai"
	Human ActorType = "human"
)

type Actor struct {
	ID   uuid.UUID `json:"id"`
	Type ActorType `json:"type"`
}

type Message struct {
	ID    uuid.UUID `json:"id"`
	Actor Actor     `json:"actor"`
	Text  string    `json:"text"`
}

type Chat struct {
	ID       uuid.UUID `json:"id"`
	Actors   []Actor   `json:"actors"`
	Messages []Message `json:"messages"`
}
