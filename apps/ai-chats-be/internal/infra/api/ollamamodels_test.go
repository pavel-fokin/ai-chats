package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ai-chats/internal/domain"
)

type MockOllamaApp struct {
	mock.Mock
}

func (m *MockOllamaApp) ListModels(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaApp) PullModel(ctx context.Context, modelName string) error {
	args := m.Called(ctx, modelName)
	return args.Error(0)
}

func (m *MockOllamaApp) DeleteModel(ctx context.Context, modelName string) error {
	args := m.Called(ctx, modelName)
	return args.Error(0)
}

func TestGetOllamaModels(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama-models", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		models := []domain.OllamaModel{
			*domain.NewOllamaModel("model1"),
			*domain.NewOllamaModel("model2"),
		}

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("ListModels", mock.MatchedBy(matchChiContext)).Return(models, nil)

		router := chi.NewRouter()
		router.Get("/api/ollama-models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama-models", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("ListModels", mock.MatchedBy(matchChiContext)).Return(nil, assert.AnError)

		router := chi.NewRouter()
		router.Get("/api/ollama-models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestPostOllamaModels(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		body := `{"model":"model1"}`

		req, _ := http.NewRequest("POST", "/api/ollama-models", strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("PullModel", mock.MatchedBy(matchChiContext), "model1").Return(nil)

		router := chi.NewRouter()
		router.Post("/api/ollama-models", PostOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		body := `{"model":"model1"}`

		req, _ := http.NewRequest("POST", "/api/ollama-models", strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("PullModel", mock.MatchedBy(matchChiContext), "model1").Return(assert.AnError)

		router := chi.NewRouter()
		router.Post("/api/ollama-models", PostOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestDeleteOllamaModels(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/ollama-models/model1", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("DeleteModel", mock.MatchedBy(matchChiContext), "model1").Return(nil)

		router := chi.NewRouter()
		router.Delete("/api/ollama-models/{model}", DeleteOllamaModel(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/ollama-models/model1", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("DeleteModel", mock.MatchedBy(matchChiContext), "model1").Return(assert.AnError)

		router := chi.NewRouter()
		router.Delete("/api/ollama-models/{model}", DeleteOllamaModel(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}
