//go:build !race

package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"
)

type ChatMock struct {
	mock.Mock
}

func (m *ChatMock) AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Chat), args.Error(1)
}

func (m *ChatMock) CreateChat(ctx context.Context, userID uuid.UUID) (domain.Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(domain.Chat), args.Error(1)
}

func (m *ChatMock) SendMessage(ctx context.Context, chatID uuid.UUID, message string) (domain.Message, error) {
	args := m.Called(ctx, chatID, message)
	return args.Get(0).(domain.Message), args.Error(1)
}

func (m *ChatMock) AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]domain.Message), args.Error(1)
}

func (m *ChatMock) ChatExists(ctx context.Context, chatID uuid.UUID) (bool, error) {
	args := m.Called(ctx, chatID)
	return args.Bool(0), args.Error(1)
}

type EventsMock struct {
	mock.Mock
}

func (m *EventsMock) Subscribe(ctx context.Context, topic string) (chan []byte, error) {
	args := m.Called(ctx, topic)
	return args.Get(0).(chan []byte), args.Error(1)
}

func (m *EventsMock) Unsubscribe(ctx context.Context, topic string, channel chan []byte) error {
	args := m.Called(ctx, topic, channel)
	return args.Error(0)
}

func matchChiContext(ctx context.Context) bool {
	key := ctx.Value(chi.RouteCtxKey)
	return key != nil
}

func TestCreateChat(t *testing.T) {
	userID := uuid.New()
	ctx := context.WithValue(context.Background(), UserID("UserID"), userID)

	t.Run("Missed UserID", func(t *testing.T) {
		defer func() { recover() }()

		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		PostChats(&ChatMock{})(w, req)

		assert.Fail(t, "expected panic")
	})

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("CreateChat", ctx, userID).Return(domain.Chat{}, nil)

		PostChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		chat.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("CreateChat", ctx, userID).Return(domain.Chat{}, errors.New("failed to create chat"))

		PostChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		chat.AssertNumberOfCalls(t, "CreateChat", 1)
	})
}

func TestGetChats(t *testing.T) {
	userID := uuid.New()
	ctx := context.WithValue(context.Background(), UserID("UserID"), userID)

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("AllChats", ctx, userID).Return([]domain.Chat{}, nil)

		GetChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		chat := &ChatMock{}
		chat.On("AllChats", ctx, userID).Return([]domain.Chat{}, errors.New("failed to get chats"))

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
func TestGetEvents(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		w := httptest.NewRecorder()

		sse := apiutil.NewSSEConnections()

		app := &ChatMock{}
		app.On("ChatExists", mock.MatchedBy(matchChiContext), chatID).
			Return(true, nil)

		eventsCh := make(chan []byte)
		defer close(eventsCh)
		events := &EventsMock{}
		events.On("Subscribe", mock.MatchedBy(matchChiContext), chatID.String()).
			Return(eventsCh, nil)
		events.On("Unsubscribe", mock.MatchedBy(matchChiContext), chatID.String(), mock.Anything).
			Return(nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/events", GetEvents(app, sse, events))
		go router.ServeHTTP(w, req)

		eventsCh <- []byte("event")
		time.Sleep(time.Millisecond)

		resp := w.Result()
		// Read server-sent events from repsonse body, unmarshal each to JSON and compare.
		assert.Equal(t, 200, resp.StatusCode)
		body := make([]byte, 1024)
		resp.Body.Read(body)
		assert.Contains(t, string(body), "data: \"ZXZlbnQ=\"\n\n")

		events.AssertNumberOfCalls(t, "Subscribe", 1)
		// events.AssertNumberOfCalls(t, "Unsubscribe", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		w := httptest.NewRecorder()

		sse := apiutil.NewSSEConnections()

		app := &ChatMock{}
		app.On("ChatExists", mock.MatchedBy(matchChiContext), chatID).
			Return(true, nil)

		events := &EventsMock{}
		events.On("Subscribe", mock.MatchedBy(matchChiContext), chatID.String()).
			Return(make(chan []byte), errors.New("failed to subscribe"))
		events.On("Unsubscribe", mock.MatchedBy(matchChiContext), chatID.String(), mock.Anything).
			Return(errors.New("failed to unsubscribe"))

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/events", GetEvents(app, sse, events))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
		events.AssertNumberOfCalls(t, "Subscribe", 1)
		events.AssertNumberOfCalls(t, "Unsubscribe", 0)
	})

	t.Run("Chat not found", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		w := httptest.NewRecorder()

		sse := apiutil.NewSSEConnections()

		app := &ChatMock{}
		app.On("ChatExists", mock.MatchedBy(matchChiContext), chatID).
			Return(false, nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/events", GetEvents(app, sse, &EventsMock{}))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 404, resp.StatusCode)
	})
}
