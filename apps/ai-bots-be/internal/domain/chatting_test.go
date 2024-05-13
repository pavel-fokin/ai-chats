package domain

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChats struct {
	mock.Mock
}

func (m *MockChats) Add(ctx context.Context, chat Chat) error {
	args := m.Called(ctx, chat)
	return args.Error(0)
}

func (m *MockChats) FindChat(ctx context.Context, id uuid.UUID) (Chat, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Chat), args.Error(1)
}

func (m *MockChats) AllChats(ctx context.Context, userID uuid.UUID) ([]Chat, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]Chat), args.Error(1)
}

type MockMessages struct {
	mock.Mock
}

func (m *MockMessages) Add(ctx context.Context, chat Chat, message Message) error {
	args := m.Called(ctx, chat, message)
	return args.Error(0)
}

func (m *MockMessages) AllMessages(ctx context.Context, chatID uuid.UUID) ([]Message, error) {
	args := m.Called(ctx, chatID)
	return args.Get(0).([]Message), args.Error(1)
}

type MockLLM struct {
	mock.Mock
}

func (m *MockLLM) GenerateResponse(ctx context.Context, messages []Message) (Message, error) {
	args := m.Called(ctx, messages)
	return args.Get(0).(Message), args.Error(1)
}

func TestSendMessage(t *testing.T) {
	ctx := context.Background()

	// Create a mock Chatting instance
	chats := &MockChats{}
	chats.On("FindChat", ctx, mock.Anything).Return(Chat{}, nil)

	messages := &MockMessages{}
	messages.On("Add", ctx, mock.Anything, mock.Anything).Return(nil)
	messages.On("AllMessages", ctx, mock.Anything).Return([]Message{}, nil)

	chatting := NewChatting(chats, messages)

	// Create a mock chat ID and message
	chatID := uuid.New()
	message := NewMessage("User", "Hello!")

	// Mock the newLLM function
	mockLLM := &MockLLM{}
	mockLLM.On("GenerateResponse", ctx, mock.Anything).Return(Message{
		Sender: "Bot",
		Text:   "Hi!",
	}, nil)

	// Call the SendMessage method
	message, err := chatting.SendMessage(ctx, mockLLM, chatID, message)
	assert.NoError(t, err)
	assert.Equal(t, "Hi!", message.Text)

	// Verify that the Add method was called twice
	messages.AssertNumberOfCalls(t, "Add", 2)
}
