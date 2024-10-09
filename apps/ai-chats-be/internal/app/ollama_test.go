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

func (m *MockOllamaClient) List(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaClient) Pull(ctx context.Context, model string, fn domain.PullingStreamFunc) error {
	args := m.Called(ctx, model, fn)
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

func (m *MockOllamaModels) FindOllamaModelsPullingInProgress(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaModels) AddModelPullingStarted(ctx context.Context, model string) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockOllamaModels) AddModelPullingFinished(ctx context.Context, model string, status domain.OllamaPullingFinalStatus) error {
	args := m.Called(ctx, model, status)
	return args.Error(0)
}

func TestAppOllama_FindOllamaModels(t *testing.T) {
	ctx := context.Background()

	t.Run("pulling status", func(t *testing.T) {
		mockModels := &MockModels{}
		mockModels.On("FindDescription", ctx, "model1").Return("description", nil)
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("FindOllamaModelsPullingInProgress", ctx).Return([]domain.OllamaModel{
			{
				Model: "model1",
			},
		}, nil)

		app := &App{
			models:       mockModels,
			ollamaModels: mockOllamaModels,
		}

		filter, err := domain.NewOllamaModelsFilter("pulling")
		assert.NoError(t, err)

		models, err := app.FindOllamaModels(ctx, filter)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:       "model1",
				Description: "description",
				IsPulling:   true,
			},
		}, models)
		mockOllamaModels.AssertExpectations(t)
	})

	t.Run("available status", func(t *testing.T) {
		mockModels := &MockModels{}
		mockModels.On("FindDescription", ctx, "model1").Return("description", nil)
		mockOllamaClient := &MockOllamaClient{}
		mockOllamaClient.On("List", ctx).Return([]domain.OllamaModel{
			{
				Model: "model1",
			},
		}, nil)

		app := &App{
			models:       mockModels,
			ollamaClient: mockOllamaClient,
		}

		filter, err := domain.NewOllamaModelsFilter("available")
		assert.NoError(t, err)

		models, err := app.FindOllamaModels(ctx, filter)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:       "model1",
				Description: "description",
				IsPulling:   false,
			},
		}, models)
		mockOllamaClient.AssertExpectations(t)
	})

	t.Run("any status", func(t *testing.T) {
		mockModels := &MockModels{}
		mockModels.On("FindDescription", ctx, "model1").Return("description", nil)
		mockModels.On("FindDescription", ctx, "model2").Return("description", nil)
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("FindOllamaModelsPullingInProgress", ctx).Return([]domain.OllamaModel{
			{
				Model: "model1",
			},
		}, nil)

		mockOllamaClient := &MockOllamaClient{}
		mockOllamaClient.On("List", ctx).Return([]domain.OllamaModel{
			{
				Model: "model2",
			},
		}, nil)

		app := &App{
			models:       mockModels,
			ollamaModels: mockOllamaModels,
			ollamaClient: mockOllamaClient,
		}

		filter := domain.OllamaModelsFilter{
			Status: domain.OllamaModelStatusAny,
		}

		models, err := app.FindOllamaModels(ctx, filter)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:       "model1",
				Description: "description",
				IsPulling:   true,
			},
			{
				Model:       "model2",
				Description: "description",
				IsPulling:   false,
			},
		}, models)
		mockOllamaModels.AssertExpectations(t)
	})
}
