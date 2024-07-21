//go:build !race

package app

import (
	"context"
	"errors"
	"ai-chats/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateChat(t *testing.T) {
	ctx := context.Background()

	user := domain.NewUser("username", "password")
	mockUsers := &MockUsers{}
	mockUsers.On("FindByID", mock.Anything, user.ID).Return(user, nil)

	mockPubSub := &MockPubSub{}
	mockPubSub.On(
		"Publish",
		ctx,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("[]uint8"),
	).Return(nil)

	mockMessages := &MockMessages{}
	mockMessages.On(
		"Add",
		ctx,
		mock.AnythingOfType("uuid.UUID"),
		mock.AnythingOfType("domain.Message"),
	).Return(nil)

	mockTx := &MockTx{}

	t.Run("with empty message", func(t *testing.T) {
		mockChats := &MockChats{}
		mockChats.On("Add", ctx, mock.AnythingOfType("domain.Chat")).Return(nil)

		app := &App{chats: mockChats, users: mockUsers, tx: mockTx}

		chat, err := app.CreateChat(ctx, user.ID, "model", "")
		assert.NoError(t, err)
		assert.NotNil(t, chat)

		mockChats.AssertExpectations(t)
	})

	t.Run("with message", func(t *testing.T) {
		mockChats := &MockChats{}
		mockChats.On("Add", ctx, mock.AnythingOfType("domain.Chat")).Return(nil)
		mockChats.On(
			"AddMessage",
			ctx,
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("domain.Message"),
		).Return(nil)

		app := &App{chats: mockChats, users: mockUsers, pubsub: mockPubSub, tx: mockTx}

		chat, err := app.CreateChat(ctx, user.ID, "message", "model")
		assert.NoError(t, err)
		assert.NotNil(t, chat)

		mockChats.AssertExpectations(t)
	})
}

func TestChat_Delete(t *testing.T) {
	t.Run("chat exists", func(t *testing.T) {
		ctx := context.Background()
		assert := assert.New(t)

		chat := domain.NewChat(
			domain.NewUser("username", "password"),
			domain.NewModel("model"),
		)

		mockChats := &MockChats{}
		mockChats.On("Delete", ctx, chat.ID).Return(nil)

		app := &App{chats: mockChats}

		err := app.DeleteChat(ctx, chat.ID)
		assert.NoError(err)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		ctx := context.Background()
		assert := assert.New(t)

		chatID := uuid.New()
		expectedErr := errors.New("chat not found")

		mockChats := &MockChats{}
		mockChats.On("Delete", ctx, chatID).Return(expectedErr)

		app := &App{chats: mockChats}

		err := app.DeleteChat(ctx, chatID)
		assert.Error(err)
		assert.Equal(expectedErr, err)
	})

}

func TestChat_FindById(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	t.Run("chat exists", func(t *testing.T) {
		chat := domain.NewChat(
			domain.NewUser("username", "password"),
			domain.NewModel("model"),
		)

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)

		app := &App{chats: mockChats}

		foundChat, err := app.chats.FindByID(ctx, chat.ID)
		assert.NoError(err)

		assert.Equal(chat.ID, foundChat.ID)
		assert.Equal(chat.Title, foundChat.Title)
		assert.Equal(chat.CreatedAt, foundChat.CreatedAt)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		chatID := uuid.New()
		expectedErr := errors.New("chat not found")

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chatID).Return(domain.Chat{}, expectedErr)

		app := &App{chats: mockChats}

		_, err := app.chats.FindByID(ctx, chatID)
		assert.Error(err)
		assert.Equal(expectedErr, err)
	})
}
