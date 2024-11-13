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

func (m *MockOllamaClient) NewModel(model domain.OllamaModel) (domain.Model, error) {
	args := m.Called(model)
	return args.Get(0).(domain.Model), args.Error(1)
}

func (m *MockOllamaClient) List(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaClient) Pull(ctx context.Context, model string, fn domain.PullProgressFunc) error {
	args := m.Called(ctx, model, fn)
	return args.Error(0)
}

func (m *MockOllamaClient) Delete(ctx context.Context, model string) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

type MockModelsLibrary struct {
	mock.Mock
}

func (m *MockModelsLibrary) FindDescription(ctx context.Context, model string) (string, error) {
	args := m.Called(ctx, model)
	return args.String(0), args.Error(1)
}

type MockOllamaModels struct {
	mock.Mock
}

func (m *MockOllamaModels) Save(ctx context.Context, model domain.OllamaModel) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockOllamaModels) FindOllamaModelsPullInProgress(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.OllamaModel), args.Error(1)
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
		mockModelsLibrary := &MockModelsLibrary{}
		mockModelsLibrary.On("FindDescription", ctx, "model1").Return("description", nil)
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("FindOllamaModelsPullInProgress", ctx).Return([]domain.OllamaModel{
			{
				Model:  "model1",
				Name:   "model1",
				Tag:    "latest",
				Status: domain.OllamaModelStatusPulling,
			},
		}, nil)

		ollama := &Ollama{
			modelsLibrary: mockModelsLibrary,
			ollamaModels:  mockOllamaModels,
		}

		filter, err := domain.NewOllamaModelsFilter("pulling")
		assert.NoError(t, err)

		models, err := ollama.FindOllamaModels(ctx, filter)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:       "model1",
				Name:        "model1",
				Tag:         "latest",
				Description: "description",
				Status:      domain.OllamaModelStatusPulling,
			},
		}, models)
		mockOllamaModels.AssertExpectations(t)
	})

	t.Run("available status", func(t *testing.T) {
		mockModelsLibrary := &MockModelsLibrary{}
		mockModelsLibrary.On("FindDescription", ctx, "model1").Return("description", nil)
		mockOllamaClient := &MockOllamaClient{}
		mockOllamaClient.On("List", ctx).Return([]domain.OllamaModel{
			{
				Model:  "model1",
				Name:   "model1",
				Tag:    "latest",
				Status: domain.OllamaModelStatusAvailable,
			},
		}, nil)

		ollama := &Ollama{
			modelsLibrary: mockModelsLibrary,
			ollamaClient:  mockOllamaClient,
		}

		filter, err := domain.NewOllamaModelsFilter("available")
		assert.NoError(t, err)

		models, err := ollama.FindOllamaModels(ctx, filter)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:       "model1",
				Name:        "model1",
				Tag:         "latest",
				Description: "description",
				Status:      domain.OllamaModelStatusAvailable,
			},
		}, models)
		mockOllamaClient.AssertExpectations(t)
	})

	t.Run("any status", func(t *testing.T) {
		mockModelsLibrary := &MockModelsLibrary{}
		mockModelsLibrary.On("FindDescription", ctx, "model1").Return("description", nil)
		mockModelsLibrary.On("FindDescription", ctx, "model2").Return("description", nil)
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("FindOllamaModelsPullInProgress", ctx).Return([]domain.OllamaModel{
			{
				Model:  "model1",
				Name:   "model1",
				Tag:    "latest",
				Status: domain.OllamaModelStatusPulling,
			},
		}, nil)

		mockOllamaClient := &MockOllamaClient{}
		mockOllamaClient.On("List", ctx).Return([]domain.OllamaModel{
			{
				Model:  "model2",
				Name:   "model2",
				Tag:    "latest",
				Status: domain.OllamaModelStatusAvailable,
			},
		}, nil)

		ollama := &Ollama{
			modelsLibrary: mockModelsLibrary,
			ollamaModels:  mockOllamaModels,
			ollamaClient:  mockOllamaClient,
		}

		filter := domain.OllamaModelsFilter{}

		models, err := ollama.FindOllamaModels(ctx, filter)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:       "model1",
				Name:        "model1",
				Tag:         "latest",
				Description: "description",
				Status:      domain.OllamaModelStatusPulling,
			},
			{
				Model:       "model2",
				Name:        "model2",
				Tag:         "latest",
				Description: "description",
				Status:      domain.OllamaModelStatusAvailable,
			},
		}, models)
		mockOllamaModels.AssertExpectations(t)
	})
}
