package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"
)

type ChatMock struct {
	mock.Mock
}

func (m *ChatMock) AllChats(ctx context.Context) ([]domain.Chat, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func (m *ChatMock) CreateChat(ctx context.Context) (domain.Chat, error) {
	args := m.Called(ctx)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *ChatMock) SendMessage(ctx context.Context, chatID uuid.UUID, message string) (app.Message, error) {
	args := m.Called(ctx, chatID, message)
	return args.Get(0).(app.Message), args.Error(1)
}

func (m *ChatMock) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func matchChiContext(ctx context.Context) bool {
	key := ctx.Value(chi.RouteCtxKey)
	return key != nil
}

func TestCreateChat(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("CreateChat", context.Background()).Return(domain.Chat{}, nil)

		PostChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		chat.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("CreateChat", context.Background()).Return(domain.Chat{}, errors.New("failed to create chat"))

		PostChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		chat.AssertNumberOfCalls(t, "CreateChat", 1)
	})
}

func TestGetChats(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("AllChats", context.Background()).Return([]domain.Chat{}, nil)

		GetChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("AllChats", context.Background()).Return([]domain.Chat{}, errors.New("failed to get chats"))

		GetChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestGetMessages(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("AllMessages", mock.MatchedBy(matchChiContext), chatID).Return([]domain.Message{}, nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/messages", GetMessages(chat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On(
			"AllMessages", mock.MatchedBy(matchChiContext), chatID,
		).Return([]domain.Message{}, errors.New("failed to get messages")).Once()

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/messages", GetMessages(chat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}
