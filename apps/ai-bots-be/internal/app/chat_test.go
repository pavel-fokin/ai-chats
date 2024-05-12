package app

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockChatDB struct {
	mock.Mock
}

func (m *mockChatDB) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockChatDB) AddUser(ctx context.Context, username, password string) error {
	args := m.Called(ctx, username, password)
	return args.Error(0)
}

func (m *mockChatDB) CreateChat(ctx context.Context, userID uuid.UUID) (domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *mockChatDB) FindChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *mockChatDB) AddMessage(ctx context.Context, chat domain.Chat, sender, message string) error {
	args := m.Called(ctx, chat, sender, message)
	return args.Error(0)
}

func (m *mockChatDB) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *mockChatDB) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func TestCreateChat(t *testing.T) {
	user := domain.NewUser("username")

	mockDB := &mockChatDB{}
	mockDB.On("CreateChat", mock.Anything, user.ID).Return(domain.Chat{}, nil)

	mockUsersDB := &mockAuthDB{}
	mockUsersDB.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)

	app := &App{chats: mockDB, users: mockUsersDB}

	chat, err := app.CreateChat(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, chat)

	// mockDB.AssertExpectations(t)
}
