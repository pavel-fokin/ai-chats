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

type MockOllamaModels struct {
	mock.Mock
}

func (m *MockOllamaModels) AllModelsWithPullingInProgress(ctx context.Context) ([]string, error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockOllamaModels) AddModelPullingStarted(ctx context.Context, model string) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockOllamaModels) AddModelPullingFinished(ctx context.Context, model string, status domain.OllamaPullingFinalStatus) error {
	args := m.Called(ctx, model, status)
	return args.Error(0)
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
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("AllModelsWithPullingInProgress", ctx).Return([]string{}, nil)
		mockOllamaModels.On("AddModelPullingStarted", ctx, "model1").Return(nil)
		mockOllamaModels.On("AddModelPullingFinished", ctx, "model1", mock.Anything).Return(nil)

		app := &App{
			models:       mockModels,
			ollamaClient: mockOllamaClient,
			ollamaModels: mockOllamaModels,
		}

		_, err := app.ListOllamaModels(ctx)
		assert.NoError(t, err)
		mockOllamaClient.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockOllamaClient := &MockOllamaClient{}
		mockOllamaClient.On("List", ctx).Return(nil, assert.AnError)
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("AllModelsWithPullingInProgress", ctx).Return([]string{}, nil)

		app := &App{
			ollamaClient: mockOllamaClient,
			ollamaModels: mockOllamaModels,
		}

		_, err := app.ListOllamaModels(ctx)
		assert.Error(t, err)
		mockOllamaClient.AssertExpectations(t)
	})
}
