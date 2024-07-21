package app

import (
	"context"
	"ai-chats/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockModels struct {
	mock.Mock
}

func (m *MockModels) List(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockModels) Pull(ctx context.Context, modelID domain.OllamaModel) error {
	args := m.Called(ctx, modelID)
	return args.Error(0)
}

func (m *MockModels) Delete(ctx context.Context, modelID domain.OllamaModel) error {
	args := m.Called(ctx, modelID)
	return args.Error(0)
}

func TestOllamaAllModels(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		models := []domain.OllamaModel{
			domain.NewOllamaModel("model1"),
			domain.NewOllamaModel("model2"),
		}

		mockModels := new(MockModels)
		mockModels.On("List", ctx).Return(models, nil)

		app := New(nil, nil, nil, mockModels, nil, nil)

		result, err := app.ListModels(ctx)
		assert.NoError(t, err)
		assert.Equal(t, models, result)
		mockModels.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockModels := new(MockModels)
		mockModels.On("List", ctx).Return(nil, assert.AnError)

		app := New(nil, nil, nil, mockModels, nil, nil)

		_, err := app.ListModels(ctx)
		assert.Error(t, err)
		mockModels.AssertExpectations(t)
	})
}
