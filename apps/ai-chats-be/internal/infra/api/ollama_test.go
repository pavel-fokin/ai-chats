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

func (m *MockOllamaApp) AllOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaApp) FindOllamaModelsAvailable(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaApp) FindOllamaModelsPullingInProgress(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)
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

		models := []domain.OllamaModel{
			domain.NewOllamaModel("model1", "description1"),
			domain.NewOllamaModel("model2", "description2"),
		}

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.
			On("FindOllamaModelsAvailable", mock.MatchedBy(matchChiContext)).
			Return(models, nil)
		mockOllamaApp.
			On("FindOllamaModelsPullingInProgress", mock.MatchedBy(matchChiContext)).
			Return(models, nil)

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
			On("FindOllamaModelsAvailable", mock.MatchedBy(matchChiContext)).
			Return([]domain.OllamaModel{}, assert.AnError)
		mockOllamaApp.
			On("FindOllamaModelsPullingInProgress", mock.MatchedBy(matchChiContext)).
			Return([]domain.OllamaModel{}, assert.AnError)

		router := chi.NewRouter()
		router.Get("/api/ollama/models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)
	})

	t.Run("invalid query", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models?onlyPulling=invalid", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.
			On("FindOllamaModelsAvailable", mock.MatchedBy(matchChiContext)).
			Return([]domain.OllamaModel{}, assert.AnError)
		mockOllamaApp.
			On("FindOllamaModelsPullingInProgress", mock.MatchedBy(matchChiContext)).
			Return([]domain.OllamaModel{}, assert.AnError)

		router := chi.NewRouter()
		router.Get("/api/ollama/models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("onlyPulling", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama/models?onlyPulling", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockOllamaApp := &MockOllamaApp{}
		mockOllamaApp.
			On("FindOllamaModelsPullingInProgress", mock.MatchedBy(matchChiContext)).
			Return([]domain.OllamaModel{
				domain.NewOllamaModel("model1", "description1"),
			}, nil)

		router := chi.NewRouter()
		router.Get("/api/ollama/models", GetOllamaModels(mockOllamaApp))

		router.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
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
