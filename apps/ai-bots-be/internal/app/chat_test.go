package app

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockChatDB struct {
	mock.Mock
}

func (m *mockChatDB) CreateChat(ctx context.Context, actors []domain.Actor) (domain.Chat, error) {
	args := m.Called(ctx, actors)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *mockChatDB) FindActorByType(ctx context.Context, actorType domain.ActorType) (domain.Actor, error) {
	args := m.Called(ctx, actorType)
	return args.Get(0).(domain.Actor), args.Error(1)
}

func (m *mockChatDB) FindChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *mockChatDB) AddMessage(ctx context.Context, chat domain.Chat, actor domain.Actor, message string) error {
	args := m.Called(ctx, chat, actor, message)
	return args.Error(0)
}

func (m *mockChatDB) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *mockChatDB) CreateActor(ctx context.Context, actorType domain.ActorType) (domain.Actor, error) {
	args := m.Called(ctx, actorType)
	return args.Get(0).(domain.Actor), args.Error(1)
}

func (m *mockChatDB) FindActor(ctx context.Context, actorID uuid.UUID) (domain.Actor, error) {
	args := m.Called(ctx, actorID)
	return args.Get(0).(domain.Actor), args.Error(1)
}

func (m *mockChatDB) AllChats(ctx context.Context) ([]domain.Chat, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func TestCreateChat(t *testing.T) {
	mockAI := domain.Actor{Type: domain.AI}
	mockHuman := domain.Actor{Type: domain.Human}

	mockDB := &mockChatDB{}
	mockDB.On("FindActorByType", mock.Anything, domain.AI).Return(mockAI, nil)
	mockDB.On("FindActorByType", mock.Anything, domain.Human).Return(mockHuman, nil)
	mockDB.On("CreateChat", mock.Anything, []domain.Actor{mockAI, mockHuman}).Return(domain.Chat{}, nil)

	app := &App{chatDB: mockDB}

	chat, err := app.CreateChat(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, chat)

	mockDB.AssertExpectations(t)
}
