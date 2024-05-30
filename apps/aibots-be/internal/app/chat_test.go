//go:build !race

package app

import (
	"context"
	"errors"
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

func (m *MockChats) UpdateTitle(ctx context.Context, chatID uuid.UUID, title string) error {
	args := m.Called(ctx, chatID, title)
	return args.Error(0)
}

func (m *MockChats) FindByID(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *MockChats) AddMessage(ctx context.Context, chatID uuid.UUID, sender, message string) error {
	args := m.Called(ctx, chatID, sender, message)
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
	mockUsers.On("FindByID", mock.Anything, user.ID).Return(user, nil)

	app := &App{chats: mockChats, users: mockUsers}

	chat, err := app.CreateChat(ctx, user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, chat)

	mockChats.AssertExpectations(t)
}

func TestChat_FindById(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username")
		chat := domain.NewChat(user)

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)

		app := &App{chats: mockChats}

		foundChat, err := app.chats.FindByID(ctx, chat.ID)
		assert.NoError(err)

		assert.Equal(chat.ID, foundChat.ID)
		assert.Equal(chat.Title, foundChat.Title)
		assert.Equal(chat.CreatedAt, foundChat.CreatedAt)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		chatID := uuid.New()
		expectedErr := errors.New("chat not found")

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chatID).Return(domain.Chat{}, expectedErr)

		app := &App{chats: mockChats}

		_, err := app.chats.FindByID(ctx, chatID)
		assert.Error(err)
		assert.Equal(expectedErr, err)
	})
}
