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
	CreateActor(ctx context.Context, actorType string) (domain.Actor, error)
	FindActor(ctx context.Context, actorID uuid.UUID) (domain.Actor, error)
}

type App struct {
	chatbot ChatBot
	db      ChatDB
	chat    domain.Chat
}

func New(chatbot ChatBot, db ChatDB) *App {
	aiActor, err := db.CreateActor(context.Background(), "ai")
	if err != nil {
		log.Fatal(err)
	}

	userActor, err := db.CreateActor(context.Background(), "user")
	if err != nil {
		log.Fatal(err)
	}

	chat, err := db.CreateChat(context.Background(), []domain.Actor{aiActor, userActor})
	if err != nil {
		log.Fatal(err)
	}

	return &App{chatbot: chatbot, db: db, chat: chat}
}

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (a *App) SendMessage(ctx context.Context, userID uuid.UUID, chatID uuid.UUID, message string) (Message, error) {
	chat, err := a.db.FindChat(ctx, a.chat.ID)
	if err != nil {
		return Message{}, err
	}

	userActor, err := a.db.FindActor(ctx, a.chat.Actors[1].ID)
	if err != nil {
		return Message{}, err
	}

	if err := a.db.AddMessage(ctx, chat, userActor, message); err != nil {
		return Message{}, err
	}

	history, err := a.db.AllMessages(ctx, chat.ID)
	if err != nil {
		return Message{}, err
	}

	aiMessage, err := a.chatbot.ChatMessage(ctx, history, message)
	if err != nil {
		return Message{}, err
	}

	aiActor := a.chat.Actors[0]
	if err := a.db.AddMessage(ctx, chat, aiActor, aiMessage.Text); err != nil {
		return Message{}, err
	}

	return aiMessage, nil
}
