package app

import (
	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockChats struct {
	mock.Mock
}

func (m *MockChats) Add(ctx context.Context, chat domain.Chat) error {
	args := m.Called(ctx, chat)
	return args.Error(0)
}

func (m *MockChats) AddMessage(ctx context.Context, chatID domain.ChatID, message domain.Message) error {
	args := m.Called(ctx, chatID, message)
	return args.Error(0)
}

func (m *MockChats) AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *MockChats) Delete(ctx context.Context, chatID domain.ChatID) error {
	args := m.Called(ctx, chatID)
	return args.Error(0)
}

func (m *MockChats) UpdateTitle(ctx context.Context, chatID domain.ChatID, title string) error {
	args := m.Called(ctx, chatID, title)
	return args.Error(0)
}

func (m *MockChats) FindByID(ctx context.Context, chatID domain.ChatID) (domain.Chat, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChats) FindByIDWithMessages(ctx context.Context, chatID domain.ChatID) (domain.Chat, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChats) AllChats(ctx context.Context, userID domain.UserID) ([]domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func (m *MockChats) Exists(ctx context.Context, chatID domain.ChatID) (bool, error) {
	args := m.Called(ctx, chatID)
	return args.Bool(0), args.Error(1)
}

type MockPubSub struct {
	mock.Mock
}

func (m *MockPubSub) Subscribe(ctx context.Context, topic string) (chan events.Event, error) {
	args := m.Called(ctx, topic)
	return args.Get(0).(chan events.Event), args.Error(1)
}

func (m *MockPubSub) Unsubscribe(ctx context.Context, topic string, channel chan events.Event) error {
	args := m.Called(ctx, topic, channel)
	return args.Error(0)
}

func (m *MockPubSub) Publish(ctx context.Context, topic string, data events.Event) error {
	args := m.Called(ctx, topic, data)
	return args.Error(0)
}

type MockMessages struct {
	mock.Mock
}

func (m *MockMessages) Add(ctx context.Context, chatID domain.ChatID, message domain.Message) error {
	args := m.Called(ctx, chatID, message)
	return args.Error(0)
}

func (m *MockMessages) AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

type MockTx struct{}

func (m *MockTx) Tx(ctx context.Context, f func(context.Context) error) error {
	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}
	return nil
}
