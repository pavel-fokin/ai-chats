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

type MockOllamaModels struct {
	mock.Mock
}

func (m *MockOllamaModels) Add(ctx context.Context, model domain.OllamaModel) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockOllamaModels) AllAdded(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaModels) AllAvailable(ctx context.Context) ([]domain.OllamaModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaModels) Delete(ctx context.Context, model domain.OllamaModel) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockOllamaModels) Exists(ctx context.Context, model string) (bool, error) {
	args := m.Called(ctx, model)
	return args.Bool(0), args.Error(1)
}

func (m *MockOllamaModels) Find(ctx context.Context, model string) (domain.OllamaModel, error) {
	args := m.Called(ctx, model)
	return args.Get(0).(domain.OllamaModel), args.Error(1)
}

func (m *MockOllamaModels) Save(ctx context.Context, model domain.OllamaModel) error {
	args := m.Called(ctx, model)
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
		mockOllamaModels := &MockOllamaModels{}
		mockOllamaModels.On("Add", ctx, mock.Anything).Return(nil)
		mockOllamaModels.On("AllAdded", ctx).Return([]domain.OllamaModel{}, nil)
		mockOllamaModels.On("Find", ctx, "model1:latest").Return(domain.OllamaModel{Model: "model1:latest"}, nil)
		mockOllamaModels.On("Save", ctx, mock.Anything).Return(nil)

		// app := New(nil, nil, mockModels, mockOllamaClient, nil, nil, nil)
		app := &App{
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

		app := &App{
			ollamaClient: mockOllamaClient,
		}

		_, err := app.ListOllamaModels(ctx)
		assert.Error(t, err)
		mockOllamaClient.AssertExpectations(t)
	})
}
