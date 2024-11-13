package app

import (
	"ai-chats/internal/domain"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestApp_CreateChat(t *testing.T) {
	ctx := context.Background()

	user := domain.NewUser("username")
	mockUsers := &MockUsers{}
	mockUsers.On("FindByID", mock.Anything, user.ID).Return(user, nil)

	mockPubSub := &MockPubSub{}
	mockPubSub.On(
		"Publish",
		ctx,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("domain.MessageAdded"),
	).Return(nil)

	mockTx := &MockTx{}

	t.Run("with empty message", func(t *testing.T) {
		mockChats := &MockChats{}

		app := &App{chats: mockChats, tx: mockTx}

		chat, err := app.CreateChat(ctx, user.ID, "model", "")
		assert.Error(t, err)
		assert.Equal(t, "message text is empty", err.Error())
		assert.Equal(t, domain.Chat{}, chat)

		mockChats.AssertExpectations(t)
	})

	t.Run("with message", func(t *testing.T) {
		mockChats := &MockChats{}
		mockChats.On("Add", ctx, mock.AnythingOfType("domain.Chat")).Return(nil)

		app := &App{chats: mockChats, pubsub: mockPubSub, tx: mockTx}

		chat, err := app.CreateChat(ctx, user.ID, "message", "model")
		assert.NoError(t, err)
		assert.NotNil(t, chat)

		mockChats.AssertExpectations(t)
	})
}

func TestApp_DeleteChat(t *testing.T) {
	ctx := context.Background()

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username")
		chat := domain.NewChat(
			user,
			domain.NewModelID("model"),
		)

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)
		mockChats.On("Delete", ctx, chat.ID).Return(nil)

		app := &App{chats: mockChats}

		err := app.DeleteChat(ctx, user.ID, chat.ID)
		assert.NoError(t, err)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		chatID := domain.NewChatID()
		userID := domain.NewUserID()

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chatID).Return(domain.Chat{}, domain.ErrChatNotFound)

		app := &App{chats: mockChats}

		err := app.DeleteChat(ctx, userID, chatID)
		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrChatNotFound)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat access denied", func(t *testing.T) {
		chat := domain.NewChat(domain.NewUser("otheruser"), domain.NewModelID("model"))

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)

		app := &App{chats: mockChats}

		err := app.DeleteChat(ctx, domain.NewUserID(), chat.ID)
		assert.ErrorIs(t, err, domain.ErrChatAccessDenied)
		mockChats.AssertExpectations(t)
	})
}

func TestApp_FindChatByID(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username")

		chat := domain.NewChat(
			user,
			domain.NewModelID("model"),
		)

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)

		app := &App{chats: mockChats}

		foundChat, err := app.FindChatByID(ctx, user.ID, chat.ID)
		assert.NoError(err)

		assert.Equal(chat.ID, foundChat.ID)
		assert.Equal(chat.Title, foundChat.Title)
		assert.Equal(chat.CreatedAt, foundChat.CreatedAt)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		chatID := uuid.New()
		userID := domain.NewUserID()

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chatID).Return(domain.Chat{}, domain.ErrChatNotFound)

		app := &App{chats: mockChats}

		_, err := app.FindChatByID(ctx, userID, chatID)
		assert.Error(err)
		assert.ErrorIs(err, domain.ErrChatNotFound)
	})

	t.Run("chat access denied", func(t *testing.T) {
		userID := domain.NewUserID()
		chat := domain.NewChat(domain.NewUser("username"), domain.NewModelID("model"))

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)
		app := &App{chats: mockChats}

		_, err := app.FindChatByID(ctx, userID, chat.ID)
		assert.ErrorIs(err, domain.ErrChatAccessDenied)
	})
}

func TestApp_SendMessage(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	user := domain.NewUser("username")
	mockUsers := &MockUsers{}
	mockUsers.On("FindByID", mock.Anything, user.ID).Return(user, nil)

	mockPubSub := &MockPubSub{}
	mockPubSub.On(
		"Publish",
		ctx,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("domain.MessageAdded"),
	).Return(nil)

	mockTx := &MockTx{}

	t.Run("chat exists", func(t *testing.T) {
		chat := domain.NewChat(
			user,
			domain.NewModelID("model"),
		)

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)
		mockChats.On("Update", ctx, mock.AnythingOfType("domain.Chat")).Return(nil)

		app := &App{chats: mockChats, pubsub: mockPubSub, tx: mockTx}

		err := app.SendMessage(ctx, user.ID, chat.ID, "Hello, how are you?")
		assert.NoError(err)
		mockChats.AssertExpectations(t)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		chatID := domain.NewChatID()

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chatID).Return(domain.Chat{}, domain.ErrChatNotFound)

		app := &App{chats: mockChats, pubsub: mockPubSub, tx: mockTx}

		err := app.SendMessage(ctx, user.ID, chatID, "Hello, how are you?")
		assert.Error(err)
		assert.ErrorContains(err, domain.ErrChatNotFound.Error())
	})

	t.Run("chat access denied", func(t *testing.T) {
		chat := domain.NewChat(domain.NewUser("otheruser"), domain.NewModelID("model"))

		mockChats := &MockChats{}
		mockChats.On("FindByID", ctx, chat.ID).Return(chat, nil)

		app := &App{chats: mockChats, pubsub: mockPubSub, tx: mockTx}

		err := app.SendMessage(ctx, user.ID, chat.ID, "Hello, how are you?")
		assert.ErrorIs(err, domain.ErrChatAccessDenied)
	})
}
