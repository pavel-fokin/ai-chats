package app

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"

	"ai-chats/internal/domain"
)

type MockModel struct {
	mock.Mock
}

func (m *MockModel) ID() domain.ModelID {
	args := m.Called()
	return args.Get(0).(domain.ModelID)
}

func (m *MockModel) Chat(
	ctx context.Context,
	messages []domain.Message,
	fn domain.ModelResponseFunc,
) (domain.Message, error) {
	args := m.Called(ctx, messages, fn)
	return args.Get(0).(domain.Message), args.Error(1)
}

func TestApp_GenerateTitle(t *testing.T) {
	user := domain.NewUser("user1")
	chat := domain.NewChat(user, domain.NewModelID("model1"))
	chat.AddMessage(domain.NewUserMessage(chat.User, "Hello, how are you?"))
	chat.AddMessage(domain.NewModelMessage(domain.NewModelID("model1"), "Hello, how can I help you?"))

	mockChats := &MockChats{}
	mockChats.On("FindByIDWithMessages", mock.Anything, mock.Anything).Return(chat, nil)
	mockChats.On("Update", mock.Anything, mock.Anything).Return(nil)

	mockPubSub := &MockPubSub{}
	mockPubSub.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockModel := &MockModel{}
	mockModel.On("ID", mock.Anything).Return(domain.NewModelID("model1"))
	mockModel.On("Chat", mock.Anything, mock.Anything, mock.Anything).
		Return(domain.NewModelMessage(domain.NewModelID("model1"), "Conversation"), nil)

	mockOllamaClient := &MockOllamaClient{}
	mockOllamaClient.On("NewModel", mock.Anything).Return(mockModel, nil)

	mockTx := &MockTx{}

	llm := &LLM{
		chats:        mockChats,
		pubsub:       mockPubSub,
		ollamaClient: mockOllamaClient,
		tx:           mockTx,
	}

	llm.GenerateTitle(context.Background(), chat.ID)

	mockChats.AssertExpectations(t)
	mockPubSub.AssertExpectations(t)
	mockOllamaClient.AssertExpectations(t)
}
