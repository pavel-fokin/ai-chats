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

func (m *MockOllamaApp) GetOllamaModelsLibrary(ctx context.Context) ([]*domain.ModelCard, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.ModelCard), args.Error(1)
}

func (m *MockOllamaApp) FindOllamaModels(ctx context.Context, filter domain.OllamaModelsFilter) ([]domain.OllamaModel, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaApp) PullOllamaModelAsync(ctx context.Context, modelName string) error {
	args := m.Called(ctx, modelName)
	return args.Error(0)
}

func (m *MockOllamaApp) DeleteOllamaModel(ctx context.Context, modelName string) error {
	args := m.Called(ctx, modelName)
	return args.Error(0)
}

func TestGetOllamaModels(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		model1, _ := domain.NewOllamaModel("model1")
		model2, _ := domain.NewOllamaModel("model2")
		models := []domain.OllamaModel{
			model1,
			model2,
		}

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On(
			"FindOllamaModels",
			mock.MatchedBy(matchChiContext),
			domain.OllamaModelsFilter{},
		).Return(models, nil)

		router := chi.NewRouter()
		router.Get("/api/ollama/models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.
			On(
				"FindOllamaModels",
				mock.MatchedBy(matchChiContext),
				domain.OllamaModelsFilter{},
			).
			Return([]domain.OllamaModel{}, assert.AnError)

		router := chi.NewRouter()
		router.Get("/api/ollama/models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})

	t.Run("invalid query", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models?status=invalid", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.
			On(
				"FindOllamaModels",
				mock.MatchedBy(matchChiContext),
				domain.OllamaModelsFilter{},
			).
			Return([]domain.OllamaModel{}, assert.AnError)

		router := chi.NewRouter()
		router.Get("/api/ollama/models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("status=pulling", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models?status=pulling", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		model1, _ := domain.NewOllamaModel("model1")
		model1.SetStatus(domain.OllamaModelStatusPulling)

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.
			On(
				"FindOllamaModels",
				mock.MatchedBy(matchChiContext),
				domain.OllamaModelsFilter{Status: domain.OllamaModelStatusPulling},
			).
			Return([]domain.OllamaModel{
				model1,
			}, nil)

		router := chi.NewRouter()
		router.Get("/api/ollama/models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
		mockOllamaApp.AssertExpectations(t)
	})
}

func TestApiOllama_GetOllamaModelsLibrary(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models-library", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("GetOllamaModelsLibrary", mock.MatchedBy(matchChiContext)).Return([]*domain.ModelCard{}, nil)

		router := chi.NewRouter()
		router.Get("/api/ollama/models-library", GetOllamaModelsLibrary(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
		mockOllamaApp.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models-library", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("GetOllamaModelsLibrary", mock.MatchedBy(matchChiContext)).Return([]*domain.ModelCard{}, assert.AnError)

		router := chi.NewRouter()
		router.Get("/api/ollama/models-library", GetOllamaModelsLibrary(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
		mockOllamaApp.AssertExpectations(t)
	})
}

func TestPostOllamaModels(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		body := `{"model":"model1"}`

		req, _ := http.NewRequest("POST", "/api/ollama/models", strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("PullOllamaModelAsync", mock.MatchedBy(matchChiContext), "model1").Return(nil)

		router := chi.NewRouter()
		router.Post("/api/ollama/models", PostOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		body := `{"model":"model1"}`

		req, _ := http.NewRequest("POST", "/api/ollama/models", strings.NewReader(body))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("PullOllamaModelAsync", mock.MatchedBy(matchChiContext), "model1").Return(assert.AnError)

		router := chi.NewRouter()
		router.Post("/api/ollama/models", PostOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestDeleteOllamaModels(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/ollama/models/model1", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("DeleteOllamaModel", mock.MatchedBy(matchChiContext), "model1").Return(nil)

		router := chi.NewRouter()
		router.Delete("/api/ollama/models/{model}", DeleteOllamaModel(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 204, resp.StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/ollama/models/model1", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.On("DeleteOllamaModel", mock.MatchedBy(matchChiContext), "model1").Return(assert.AnError)

		router := chi.NewRouter()
		router.Delete("/api/ollama/models/{model}", DeleteOllamaModel(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})
}
