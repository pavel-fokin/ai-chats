package api

import (
	"context"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"
)

type MockOllamaApp struct {
	mock.Mock
}

func (m *MockOllamaApp) ListModels(ctx context.Context) ([]domain.Model, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.Model), args.Error(1)
}

func TestGetOllamaModels(t *testing.T) {
	ctx := context.WithValue(context.Background(), apiutil.UserIDCtxKey, uuid.New())

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ollama-models", nil)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		models := []domain.Model{
			domain.NewModel("model1", "latest"),
			domain.NewModel("model2", "latest"),
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
