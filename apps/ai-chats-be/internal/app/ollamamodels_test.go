package app

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockModels struct {
	mock.Mock
}

func (m *MockModels) List(ctx context.Context) ([]domain.Model, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.Model), args.Error(1)
}

func (m *MockModels) Pull(ctx context.Context, modelID domain.Model) error {
	args := m.Called(ctx, modelID)
	return args.Error(0)
}

func (m *MockModels) Delete(ctx context.Context, modelID domain.Model) error {
	args := m.Called(ctx, modelID)
	return args.Error(0)
}

func TestOllamaAllModels(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		models := []domain.Model{
			domain.NewModel("model1", "latest"),
			domain.NewModel("model2", "latest"),
		}

		mockModels := new(MockModels)
		mockModels.On("List", ctx).Return(models, nil)

		app := New(nil, nil, mockModels, nil, nil)

		result, err := app.ListModels(ctx)
		assert.NoError(t, err)
		assert.Equal(t, models, result)
		mockModels.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockModels := new(MockModels)
		mockModels.On("List", ctx).Return(nil, assert.AnError)

		app := New(nil, nil, mockModels, nil, nil)

		_, err := app.ListModels(ctx)
		assert.Error(t, err)
		mockModels.AssertExpectations(t)
	})
}
