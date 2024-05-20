package app

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChats struct {
	mock.Mock
}

func (m *MockChats) Add(ctx context.Context, chat domain.Chat) error {
	args := m.Called(ctx, chat)
	return args.Error(0)
}

func (m *MockChats) FindChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChats) AddMessage(ctx context.Context, chat domain.Chat, sender, message string) error {
	args := m.Called(ctx, chat, sender, message)
	return args.Error(0)
}

func (m *MockChats) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *MockChats) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func (m *MockChats) Exists(ctx context.Context, chatID uuid.UUID) (bool, error) {
	args := m.Called(ctx, chatID)
	return args.Bool(0), args.Error(1)
}

func TestCreateChat(t *testing.T) {
	ctx := context.Background()

	user := domain.NewUser("username")

	mockChats := &MockChats{}
	mockChats.On("Add", ctx, mock.AnythingOfType("domain.Chat")).Return(nil)

	mockUsers := &MockUsers{}
	mockUsers.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)

	app := &App{chats: mockChats, users: mockUsers}

	chat, err := app.CreateChat(ctx, user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, chat)

	mockChats.AssertExpectations(t)
}
