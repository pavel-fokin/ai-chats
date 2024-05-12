package domain

import (
	"context"

	"github.com/google/uuid"
)

type Chats interface {
	CreateChat(ctx context.Context, userID uuid.UUID, actors []Actor) (Chat, error)
	AllChats(ctx context.Context, userID uuid.UUID) ([]Chat, error)
	FindChat(ctx context.Context, chatID uuid.UUID) (Chat, error)
	AddMessage(ctx context.Context, chat Chat, actor Actor, message string) error
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]Message, error)
	CreateActor(ctx context.Context, actorType ActorType) (Actor, error)
	FindActor(ctx context.Context, actorID uuid.UUID) (Actor, error)
	FindActorByType(ctx context.Context, actorType ActorType) (Actor, error)
}
