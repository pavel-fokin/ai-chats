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

func (m *MockChats) AddMessage(ctx context.Context, chatID domain.ChatID, sender, message string) error {
	args := m.Called(ctx, chatID, sender, message)
	return args.Error(0)
}

func (m *MockChats) AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error) {
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

type MockPubSub struct {
	mock.Mock
}

func (m *MockPubSub) Subscribe(ctx context.Context, topic string) (chan []byte, error) {
	args := m.Called(ctx, topic)
	return args.Get(0).(chan []byte), args.Error(1)
}

func (m *MockPubSub) Unsubscribe(ctx context.Context, topic string, channel chan []byte) error {
	args := m.Called(ctx, topic, channel)
	return args.Error(0)
}

func (m *MockPubSub) Publish(ctx context.Context, topic string, data []byte) error {
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

func TestCreateChat(t *testing.T) {
	ctx := context.Background()

	user := domain.NewUser("username")
	mockUsers := &MockUsers{}
	mockUsers.On("FindByID", mock.Anything, user.ID).Return(user, nil)

	mockPubSub := &MockPubSub{}
	mockPubSub.On(
		"Publish",
		ctx,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("[]uint8"),
	).Return(nil)

	mockMessages := &MockMessages{}
	mockMessages.On(
		"Add",
		ctx,
		mock.AnythingOfType("uuid.UUID"),
		mock.AnythingOfType("domain.Message"),
	).Return(nil)

	t.Run("with empty message", func(t *testing.T) {
		mockChats := &MockChats{}
		mockChats.On("Add", ctx, mock.AnythingOfType("domain.Chat")).Return(nil)

		app := &App{chats: mockChats, users: mockUsers}

		chat, err := app.CreateChat(ctx, user.ID, "")
		assert.NoError(t, err)
		assert.NotNil(t, chat)

		mockChats.AssertExpectations(t)
	})

	t.Run("with message", func(t *testing.T) {
		mockChats := &MockChats{}
		mockChats.On("Add", ctx, mock.AnythingOfType("domain.Chat")).Return(nil)

		app := &App{chats: mockChats, users: mockUsers, messages: mockMessages, pubsub: mockPubSub}

		chat, err := app.CreateChat(ctx, user.ID, "message")
		assert.NoError(t, err)
		assert.NotNil(t, chat)

		mockChats.AssertExpectations(t)
	})
}

func TestChat_Delete(t *testing.T) {
	t.Run("chat exists", func(t *testing.T) {
		ctx := context.Background()
		assert := assert.New(t)

		chat := domain.NewChat(domain.NewUser("username"))

		mockChats := &MockChats{}
		mockChats.On("Delete", ctx, chat.ID).Return(nil)

		app := &App{chats: mockChats}

		err := app.DeleteChat(ctx, chat.ID)
		assert.NoError(err)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		ctx := context.Background()
		assert := assert.New(t)

		chatID := uuid.New()
		expectedErr := errors.New("chat not found")

		mockChats := &MockChats{}
		mockChats.On("Delete", ctx, chatID).Return(expectedErr)

		app := &App{chats: mockChats}

		err := app.DeleteChat(ctx, chatID)
		assert.Error(err)
		assert.Equal(expectedErr, err)
	})

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
