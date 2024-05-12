package app

import (
	"context"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type ChatBot interface {
	SingleMessage(ctx context.Context, message string) (domain.Message, error)
	ChatMessage(ctx context.Context, history []domain.Message, message string) (domain.Message, error)
}

type ChatDB interface {
	CreateChat(ctx context.Context, userID uuid.UUID, actors []domain.Actor) (domain.Chat, error)
	AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error)
	FindChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error)
	AddMessage(ctx context.Context, chat domain.Chat, actor domain.Actor, message string) error
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error)
	CreateActor(ctx context.Context, actorType domain.ActorType) (domain.Actor, error)
	FindActor(ctx context.Context, actorID uuid.UUID) (domain.Actor, error)
	FindActorByType(ctx context.Context, actorType domain.ActorType) (domain.Actor, error)
}

type UserDB interface {
	CreateUser(ctx context.Context, username, password string) (User, error)
	FindUser(ctx context.Context, username string) (User, error)
}

type App struct {
	chatbot ChatBot
	userDB  UserDB
	chatDB  ChatDB
}

func New(chatbot ChatBot, chatDB ChatDB, userDB UserDB) *App {
	return &App{chatbot: chatbot, chatDB: chatDB, userDB: userDB}
}
