package app

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ai-chats/internal/domain"
)

type MockOllamaClient struct {
	mock.Mock
}

func (m *MockOllamaClient) List(ctx context.Context) ([]domain.OllamaClientModel, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.OllamaClientModel), args.Error(1)
}

func (m *MockOllamaClient) Pull(ctx context.Context, model string) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockOllamaClient) Delete(ctx context.Context, model string) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

type MockModels struct {
	mock.Mock
}

func (m *MockModels) FindDescription(ctx context.Context, model string) (string, error) {
	args := m.Called(ctx, model)
	return args.String(0), args.Error(1)
}

func TestAppOllama_ListModels(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		models := []domain.OllamaClientModel{
			domain.NewOllamaClientModel("model1:latest"),
		}

		mockOllamaClient := &MockOllamaClient{}
		mockOllamaClient.On("List", ctx).Return(models, nil)
		mockModels := &MockModels{}
		mockModels.On("FindDescription", ctx, "model1").Return("description", nil)

		app := &App{
			ollamaClient: mockOllamaClient,
			models:       mockModels,
		}

		_, err := app.ListOllamaModels(ctx)
		assert.NoError(t, err)
		mockOllamaClient.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockOllamaClient := &MockOllamaClient{}
		mockOllamaClient.On("List", ctx).Return(nil, assert.AnError)

		app := &App{
			ollamaClient: mockOllamaClient,
		}

		_, err := app.ListOllamaModels(ctx)
		assert.Error(t, err)
		mockOllamaClient.AssertExpectations(t)
	})
}
