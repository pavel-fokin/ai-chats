package app

import (
	"context"
	"log"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"

	"github.com/google/uuid"
)

type ChatBot interface {
	SingleMessage(ctx context.Context, message string) (Message, error)
	ChatMessage(ctx context.Context, history []domain.Message, message string) (Message, error)
}

type ChatDB interface {
	CreateChat(ctx context.Context, actors []domain.Actor) (domain.Chat, error)
	AllChats(ctx context.Context) ([]domain.Chat, error)
	FindChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error)
	AddMessage(ctx context.Context, chat domain.Chat, actor domain.Actor, message string) error
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error)
	CreateActor(ctx context.Context, actorType domain.ActorType) (domain.Actor, error)
	FindActor(ctx context.Context, actorID uuid.UUID) (domain.Actor, error)
}

type AuthDB interface {
	CreateUser(ctx context.Context, username, password string) (User, error)
	FindUser(ctx context.Context, username string) (User, error)
}

type App struct {
	chatbot ChatBot
	authDB  AuthDB
	chatDB  ChatDB
	chat    domain.Chat
}

func New(chatbot ChatBot, chatDB ChatDB) *App {
	aiActor, err := chatDB.CreateActor(context.Background(), "ai")
	if err != nil {
		log.Fatal(err)
	}

	userActor, err := chatDB.CreateActor(context.Background(), "user")
	if err != nil {
		log.Fatal(err)
	}

	chat, err := chatDB.CreateChat(context.Background(), []domain.Actor{aiActor, userActor})
	if err != nil {
		log.Fatal(err)
	}

	return &App{chatbot: chatbot, chatDB: chatDB, chat: chat}
}

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (a *App) SendMessage(ctx context.Context, userID uuid.UUID, chatID uuid.UUID, message string) (Message, error) {
	chat, err := a.chatDB.FindChat(ctx, a.chat.ID)
	if err != nil {
		return Message{}, err
	}

	userActor, err := a.chatDB.FindActor(ctx, a.chat.Actors[1].ID)
	if err != nil {
		return Message{}, err
	}

	history, err := a.chatDB.AllMessages(ctx, chat.ID)
	if err != nil {
		return Message{}, err
	}

	if err := a.chatDB.AddMessage(ctx, chat, userActor, message); err != nil {
		return Message{}, err
	}

	aiMessage, err := a.chatbot.ChatMessage(ctx, history, message)
	if err != nil {
		return Message{}, err
	}

	aiActor := a.chat.Actors[0]
	if err := a.chatDB.AddMessage(ctx, chat, aiActor, aiMessage.Text); err != nil {
		return Message{}, err
	}

	return aiMessage, nil
}
