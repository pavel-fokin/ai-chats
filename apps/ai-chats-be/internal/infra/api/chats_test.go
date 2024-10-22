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

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

func matchChiContext(ctx context.Context) bool {
	key := ctx.Value(chi.RouteCtxKey)
	return key != nil
}

func TestApiChats_CreateChat(t *testing.T) {
	userID := domain.NewUserID()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("missed UserID", func(t *testing.T) {
		defer func() { recover() }()

		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		PostChats(&MockChats{})(w, req)

		assert.Fail(t, "expected panic")
	})

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On("CreateChat", ctx, userID, "", "").Return(domain.Chat{}, nil)

		PostChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		mockChats.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"CreateChat", ctx, userID, "", "",
		).Return(domain.Chat{}, errors.New("failed to create chat"))

		PostChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		mockChats.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("with empty JSON", func(t *testing.T) {
		body := `{}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"CreateChat", ctx, userID, "", "",
		).Return(domain.Chat{}, nil)

		PostChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		mockChats.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("with invalid JSON", func(t *testing.T) {
		body := `{"12312"}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"CreateChat", ctx, userID, "", "",
		).Return(domain.Chat{}, nil)

		PostChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)

		mockChats.AssertNumberOfCalls(t, "CreateChat", 0)
	})

	t.Run("success with message", func(t *testing.T) {
		body := `{"defaultModel": "model","message": "message"}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"CreateChat", ctx, userID, "model", "message",
		).Return(domain.Chat{}, nil)

		PostChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		mockChats.AssertNumberOfCalls(t, "CreateChat", 1)
	})

	t.Run("failure with message", func(t *testing.T) {
		body := `{"defaultModel": "model","message": "message"}`
		req, err := http.NewRequest("POST", "", strings.NewReader(body))
		assert.NoError(t, err)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"CreateChat", ctx, userID, "model", "message",
		).Return(domain.Chat{}, errors.New("failed to create chat"))

		PostChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		mockChats.AssertNumberOfCalls(t, "CreateChat", 1)
	})
}

func TestApiChats_DeleteChat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()
		ctx := context.WithValue(context.Background(), UserIDCtxKey, domain.NewUserID())

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On("DeleteChat", mock.MatchedBy(matchChiContext), chatID).Return(nil)

		router := chi.NewRouter()
		router.Delete("/api/chats/{uuid}", DeleteChat(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)
		mockChats.AssertNumberOfCalls(t, "DeleteChat", 1)
	})

	t.Run("internal error", func(t *testing.T) {
		chatID := uuid.New()
		ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On("DeleteChat", mock.MatchedBy(matchChiContext), chatID).Return(errors.New("failed to delete chat"))

		router := chi.NewRouter()
		router.Delete("/api/chats/{uuid}", DeleteChat(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat not found", func(t *testing.T) {
		chatID := uuid.New()
		ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On("DeleteChat", mock.MatchedBy(matchChiContext), chatID).Return(domain.ErrChatNotFound)

		router := chi.NewRouter()
		router.Delete("/api/chats/{uuid}", DeleteChat(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 404, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})
}

func TestApiChats_GetChats(t *testing.T) {
	userID := domain.NewUserID()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On("FindChatsByUserID", ctx, userID).Return([]domain.Chat{}, nil)

		GetChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("failure", func(t *testing.T) {
		req, _ := http.NewRequest("", "", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On("FindChatsByUserID", ctx, userID).Return([]domain.Chat{}, assert.AnError)

		GetChats(mockChats)(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})
}

func TestApiChats_GetMessages(t *testing.T) {
	userID := domain.NewUserID()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("success", func(t *testing.T) {
		chatID := domain.NewChatID()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByIDWithMessages", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/messages", GetMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat not found", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByIDWithMessages", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, domain.ErrChatNotFound)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/messages", GetMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 404, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat access denied", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByIDWithMessages", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, domain.ErrChatAccessDenied)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/messages", GetMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 403, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/messages", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByIDWithMessages", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, errors.New("failed to get messages")).Once()

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/messages", GetMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})
}

func TestApiChats_GetEvents(t *testing.T) {
	userID := domain.NewUserID()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		sse := NewSSEConnections()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, nil)

		eventsCh := make(chan types.Message)
		defer close(eventsCh)
		eventsMock := &EventsMock{}
		eventsMock.On("Subscribe", mock.MatchedBy(matchChiContext), chatID.String()).
			Return(eventsCh, nil)
		eventsMock.On("Unsubscribe", mock.MatchedBy(matchChiContext), chatID.String(), mock.Anything).
			Return(nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/events", GetChatEvents(mockChats, sse, eventsMock))
		go router.ServeHTTP(w, req)

		eventsCh <- domain.MessageAdded{}
		time.Sleep(time.Millisecond)

		resp := w.Result()
		// Read server-sent events from repsonse body, unmarshal each to JSON and compare.
		assert.Equal(t, 200, resp.StatusCode)
		body := make([]byte, 1024)
		resp.Body.Read(body)
		assert.Contains(t, string(body), "event: messageAdded\ndata:")

		// Close all connections and wait for goroutines to finish.
		sse.CloseAll()
		time.Sleep(time.Millisecond)

		mockChats.AssertExpectations(t)
		eventsMock.AssertNumberOfCalls(t, "Subscribe", 1)
		eventsMock.AssertNumberOfCalls(t, "Unsubscribe", 1)
	})

	t.Run("failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		sse := NewSSEConnections()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, errors.New("failed to find chat"))

		eventsMock := &EventsMock{}
		eventsMock.On(
			"Subscribe", mock.MatchedBy(matchChiContext), chatID.String(),
		).Return(make(chan types.Message), errors.New("failed to subscribe"))
		eventsMock.On(
			"Unsubscribe", mock.MatchedBy(matchChiContext), chatID.String(), mock.Anything,
		).Return(errors.New("failed to unsubscribe"))

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/events", GetChatEvents(mockChats, sse, eventsMock))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		mockChats.AssertExpectations(t)
		eventsMock.AssertNumberOfCalls(t, "Subscribe", 0)
		eventsMock.AssertNumberOfCalls(t, "Unsubscribe", 0)
	})

	t.Run("chat not found", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s/events", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		sse := NewSSEConnections()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, domain.ErrChatNotFound)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}/events", GetChatEvents(mockChats, sse, &EventsMock{}))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 404, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})
}

func TestApiChats_GetChat(t *testing.T) {
	userID := domain.NewUserID()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, nil)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}", GetChat(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat access denied", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, domain.ErrChatAccessDenied)

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}", GetChat(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 403, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("failure", func(t *testing.T) {
		chatID := uuid.New()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/chats/%s", chatID), nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"FindChatByID", mock.MatchedBy(matchChiContext), userID, chatID,
		).Return(domain.Chat{}, errors.New("failed to get chat"))

		router := chi.NewRouter()
		router.Get("/api/chats/{uuid}", GetChat(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})
}

func TestApiChats_PostMessages(t *testing.T) {
	userID := domain.NewUserID()
	ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

	t.Run("success", func(t *testing.T) {
		chatID := uuid.New()

		body := `{"text": "text"}`
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/chats/%s/messages", chatID), strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"SendMessage", mock.MatchedBy(matchChiContext), userID, chatID, "text",
		).Return(nil)

		router := chi.NewRouter()
		router.Post("/api/chats/{uuid}/messages", PostMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat not found", func(t *testing.T) {
		chatID := uuid.New()

		body := `{"text": "text"}`
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/chats/%s/messages", chatID), strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"SendMessage", mock.MatchedBy(matchChiContext), userID, chatID, "text",
		).Return(domain.ErrChatNotFound)

		router := chi.NewRouter()
		router.Post("/api/chats/{uuid}/messages", PostMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 404, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat access denied", func(t *testing.T) {
		chatID := uuid.New()

		body := `{"text": "text"}`
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/chats/%s/messages", chatID), strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"SendMessage", mock.MatchedBy(matchChiContext), userID, chatID, "text",
		).Return(domain.ErrChatAccessDenied)

		router := chi.NewRouter()
		router.Post("/api/chats/{uuid}/messages", PostMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 403, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})

	t.Run("internal error", func(t *testing.T) {
		chatID := uuid.New()

		body := `{"text": "text"}`
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/chats/%s/messages", chatID), strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockChats := &MockChats{}
		mockChats.On(
			"SendMessage", mock.MatchedBy(matchChiContext), userID, chatID, "text",
		).Return(assert.AnError)

		router := chi.NewRouter()
		router.Post("/api/chats/{uuid}/messages", PostMessages(mockChats))
		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
		mockChats.AssertExpectations(t)
	})
}
