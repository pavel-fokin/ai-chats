package api

import (
	"context"

	"github.com/stretchr/testify/mock"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

type MockChat struct {
	mock.Mock
}

func (m *MockChat) FindChatsByUserID(ctx context.Context, userID domain.UserID) ([]domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func (m *MockChat) CreateChat(ctx context.Context, userID domain.UserID, defaultModel, message string) (domain.Chat, error) {
	args := m.Called(ctx, userID, defaultModel, message)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChat) DeleteChat(ctx context.Context, chatID domain.ChatID) error {
	args := m.Called(ctx, chatID)
	return args.Error(0)
}

func (m *MockChat) FindChatByID(ctx context.Context, chatID domain.ChatID) (domain.Chat, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChat) SendMessage(ctx context.Context, userID domain.UserID, chatID domain.ChatID, message string) (domain.Message, error) {
	args := m.Called(ctx, chatID, message)
	return args.Get(0).(domain.Message), args.Error(1)
}

func (m *MockChat) ChatMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *MockChat) ChatExists(ctx context.Context, chatID domain.ChatID) (bool, error) {
	args := m.Called(ctx, chatID)
	return args.Bool(0), args.Error(1)
}

func (m *MockChat) GenerateChatTitleAsync(ctx context.Context, chatID domain.ChatID) error {
	args := m.Called(ctx, chatID)
	return args.Error(0)
}

type EventsMock struct {
	mock.Mock
}

func (m *EventsMock) Subscribe(ctx context.Context, topic string) (chan types.Message, error) {
	args := m.Called(ctx, topic)
	return args.Get(0).(chan types.Message), args.Error(1)
}

func (m *EventsMock) Unsubscribe(ctx context.Context, topic string, channel chan types.Message) error {
	args := m.Called(ctx, topic, channel)
	return args.Error(0)
}
