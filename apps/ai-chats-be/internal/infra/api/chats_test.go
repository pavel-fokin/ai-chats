//go:build !race

package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server"
)

func matchChiContext(ctx context.Context) bool {
	key := ctx.Value(chi.RouteCtxKey)
	return key != nil
}

func TestApiCreateChat(t *testing.T) {
	userID := uuid.New()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("Missed UserID", func(t *testing.T) {
		defer func() { recover() }()

		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		PostChats(&MockChat{})(w, req)

		assert.Fail(t, "expected panic")
	})

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("CreateChat", ctx, userID, "", "").Return(domain.Chat{}, nil)

		PostChats(mockChat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		mockChat.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("CreateChat", ctx, userID, "", "").Return(domain.Chat{}, errors.New("failed to create chat"))

		PostChats(mockChat)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		mockChat.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("With empty JSON", func(t *testing.T) {
		body := `{}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("CreateChat", ctx, userID, "", "").Return(domain.Chat{}, nil)

		PostChats(mockChat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		mockChat.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("With invalid JSON", func(t *testing.T) {
		body := `{"12312"}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("CreateChat", ctx, userID, "", "").Return(domain.Chat{}, nil)

		PostChats(nil)(w, req)

		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)

		mockChat.AssertNumberOfCalls(t, "CreateChat", 0)
	})

	t.Run("success with message", func(t *testing.T) {
		body := `{"defaultModel": "model","message": "message"}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("CreateChat", ctx, userID, "model", "message").Return(domain.Chat{}, nil)

		PostChats(mockChat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		mockChat.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("failure with message", func(t *testing.T) {
		body := `{"defaultModel": "model","message": "message"}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On(
			"CreateChat", ctx, userID, "model", "message",
		).Return(domain.Chat{}, errors.New("failed to create chat"))

		PostChats(mockChat)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		mockChat.AssertNumberOfCalls(t, "CreateChat", 1)
	})
}

func TestApiDeleteChat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()
		ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("DeleteChat", mock.MatchedBy(matchChiContext), chatID).Return(nil)

		router := chi.NewRouter()
		router.Delete("/api/chats/{uuid}", DeleteChat(mockChat))
		router.ServeHTTP(w, req)

		mockChat.AssertNumberOfCalls(t, "DeleteChat", 1)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)

	})

	t.Run("internal error", func(t *testing.T) {
		chatID := uuid.New()
		ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("DeleteChat", mock.MatchedBy(matchChiContext), chatID).Return(errors.New("failed to delete chat"))

		router := chi.NewRouter()
		router.Delete("/api/chats/{uuid}", DeleteChat(mockChat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})

	t.Run("chat not found", func(t *testing.T) {
		chatID := uuid.New()
		ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("DeleteChat", mock.MatchedBy(matchChiContext), chatID).Return(domain.ErrChatNotFound)

		router := chi.NewRouter()
		router.Delete("/api/chats/{uuid}", DeleteChat(mockChat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 404, resp.StatusCode)
	})
}

func TestApiGetChats(t *testing.T) {
	userID := uuid.New()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On("AllChats", ctx, userID).Return([]domain.Chat{}, nil)

		GetChats(mockChat)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		chat := &MockChat{}
		chat.On("AllChats", ctx, userID).Return([]domain.Chat{}, errors.New("failed to get chats"))

		GetChats(chat)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestApiGetMessages(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		w := httptest.NewRecorder()

		chat := &MockChat{}
		chat.On(
			"AllMessages", mock.MatchedBy(matchChiContext), chatID,
		).Return([]domain.Message{}, nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/messages", GetMessages(chat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		w := httptest.NewRecorder()

		chat := &MockChat{}
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
func TestApiGetEvents(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		w := httptest.NewRecorder()

		sse := server.NewSSEConnections()

		app := &MockChat{}
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
		assert.Contains(t, string(body), "data: event\n\n")

		events.AssertNumberOfCalls(t, "Subscribe", 1)
		// events.AssertNumberOfCalls(t, "Unsubscribe", 1)
	})

	t.Run("failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		w := httptest.NewRecorder()

		sse := server.NewSSEConnections()

		app := &MockChat{}
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

	t.Run("chat not found", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		w := httptest.NewRecorder()

		sse := server.NewSSEConnections()

		app := &MockChat{}
		app.On("ChatExists", mock.MatchedBy(matchChiContext), chatID).
			Return(false, nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/events", GetEvents(app, sse, &EventsMock{}))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 404, resp.StatusCode)
	})
}

func TestApiGetChat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s", chatID), nil)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), chatID,
		).Return(domain.Chat{}, nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}", GetChat(mockChat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s", chatID), nil)
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), chatID,
		).Return(domain.Chat{}, errors.New("failed to get chat"))

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}", GetChat(mockChat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestApiPostMessages(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()

		body := `{"text": "text"}`
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/chats/%s/messages", chatID), strings.NewReader(body))
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On(
			"SendMessage", mock.MatchedBy(matchChiContext), chatID, "text",
		).Return(domain.Message{}, nil)

		router := chi.NewRouter()
		router.Post("/api/chats/{uuid}/messages", PostMessages(mockChat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		chatID := uuid.New()

		body := `{"text": "text"}`
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/chats/%s/messages", chatID), strings.NewReader(body))
		w := httptest.NewRecorder()

		mockChat := &MockChat{}
		mockChat.On(
			"SendMessage", mock.MatchedBy(matchChiContext), chatID, "text",
		).Return(domain.Message{}, errors.New("failed to send message"))

		router := chi.NewRouter()
		router.Post("/api/chats/{uuid}/messages", PostMessages(mockChat))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}
