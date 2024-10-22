package api

import (
	"context"

	"github.com/stretchr/testify/mock"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

type MockChats struct {
	mock.Mock
}

func (m *MockChats) FindChatsByUserID(ctx context.Context, userID domain.UserID) ([]domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func (m *MockChats) CreateChat(ctx context.Context, userID domain.UserID, defaultModel, message string) (domain.Chat, error) {
	args := m.Called(ctx, userID, defaultModel, message)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChats) DeleteChat(ctx context.Context, userID domain.UserID, chatID domain.ChatID) error {
	args := m.Called(ctx, userID, chatID)
	return args.Error(0)
}

func (m *MockChats) FindChatByID(ctx context.Context, userID domain.UserID, chatID domain.ChatID) (domain.Chat, error) {
	args := m.Called(ctx, userID, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChats) FindChatByIDWithMessages(ctx context.Context, userID domain.UserID, chatID domain.ChatID) (domain.Chat, error) {
	args := m.Called(ctx, userID, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChats) SendMessage(ctx context.Context, userID domain.UserID, chatID domain.ChatID, message string) error {
	args := m.Called(ctx, userID, chatID, message)
	return args.Error(0)
}

func (m *MockChats) GenerateChatTitleAsync(ctx context.Context, chatID domain.ChatID) error {
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
