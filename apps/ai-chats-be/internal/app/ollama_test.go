package app

import (
	"ai-chats/internal/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOllamaModels struct {
	mock.Mock
}

func (m *MockOllamaModels) List(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaModels) Pull(ctx context.Context, modelID domain.OllamaModel) error {
	args := m.Called(ctx, modelID)
	return args.Error(0)
}

func (m *MockOllamaModels) Delete(ctx context.Context, modelID domain.OllamaModel) error {
	args := m.Called(ctx, modelID)
	return args.Error(0)
}

type MockModels struct {
	mock.Mock
}

func (m *MockModels) AllModelCards(ctx context.Context) ([]domain.ModelCard, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.ModelCard), args.Error(1)
}

func (m *MockModels) AddModelCard(ctx context.Context, model domain.ModelCard) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockModels) FindModelCard(ctx context.Context, model string) (domain.ModelCard, error) {
	args := m.Called(ctx, model)
	return args.Get(0).(domain.ModelCard), args.Error(1)
}

func TestOllamaAllModels(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		models := []domain.OllamaModel{
			domain.NewOllamaModel("model1"),
			// domain.NewOllamaModel("mqodel2"),
		}

		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("List", ctx).Return(models, nil)
		mockModels := &MockModels{}
		mockModels.On("FindModelCard", ctx, "model1").Return(domain.ModelCard{Description: "description1"}, nil)

		app := New(nil, nil, mockModels, mockOllamaModels, nil, nil)

		_, err := app.ListModels(ctx)
		assert.NoError(t, err)
		// assert.Equal(t, models, result)
		mockOllamaModels.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("List", ctx).Return(nil, assert.AnError)
		mockModels := &MockModels{}

		app := New(nil, nil, mockModels, mockOllamaModels, nil, nil)

		_, err := app.ListModels(ctx)
		assert.Error(t, err)
		mockModels.AssertExpectations(t)
	})
}
