package sqlite

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"ai-chats/internal/domain"
)

func TestSqliteChats_Add(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	user := domain.NewUserWithPassword("test", "password", 1)
	err := users.Add(ctx, user)
	assert.NoError(t, err)
	modelID := domain.NewModelID("model")

	t.Run("valid", func(t *testing.T) {
		chat := domain.NewChat(user, modelID)
		err = chats.Add(ctx, chat)
		assert.NoError(t, err)
	})

	t.Run("add chat without user", func(t *testing.T) {
		chat := domain.NewChat(domain.User{}, modelID)
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})

	t.Run("add chat with empty title", func(t *testing.T) {
		chat := domain.NewChat(user, modelID)
		chat.Title = ""
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})

	t.Run("add chat with invalid user", func(t *testing.T) {
		chat := domain.NewChat(domain.NewUser(""), modelID)
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})

	t.Run("add chat with empty model id", func(t *testing.T) {
		chat := domain.NewChat(user, domain.NewModelID(""))
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})
}

func TestSqliteChats_Update(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	user := domain.NewUserWithPassword("test", "password", 1)
	err := users.Add(ctx, user)
	assert.NoError(t, err)
	modelID := domain.NewModelID("model")

	t.Run("update chat title", func(t *testing.T) {
		chat := domain.NewChat(user, modelID)
		err := chats.Add(ctx, chat)
		assert.NoError(t, err)

		chat.UpdateTitle("New title")
		err = chats.Update(ctx, chat)
		assert.NoError(t, err)

		updatedChat, err := chats.FindByID(ctx, chat.ID)
		assert.NoError(t, err)
		assert.Equal(t, "New title", updatedChat.Title)
	})

	t.Run("add message", func(t *testing.T) {
		chat := domain.NewChat(user, modelID)
		chats.Add(ctx, chat)
		chat.AddMessage(domain.NewUserMessage(user, "Hello, model!"))
		err = chats.Update(ctx, chat)
		assert.NoError(t, err)
	})
}

func TestSqliteChats_Delete(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUserWithPassword("username", "password", 1)
		err := users.Add(ctx, user)
		assert.NoError(err)

		chat := domain.NewChat(user, domain.NewModelID("model"))
		err = chats.Add(ctx, chat)
		assert.NoError(err)

		// Call the DeleteChat method.
		err = chats.Delete(ctx, chat.ID)
		assert.NoError(err)

		// Check that the chat was deleted.
		var deletedAt string
		err = db.QueryRowContext(ctx, "SELECT deleted_at FROM chat WHERE id = ?", chat.ID).Scan(&deletedAt)
		assert.NoError(err)
		assert.NotEmpty(deletedAt)

		exists, err := chats.Exists(ctx, chat.ID)
		assert.NoError(err)
		assert.False(exists)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		err := chats.Delete(ctx, uuid.New())
		assert.Error(err)
	})
}

func TestSqliteChats_FindByUserID(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("no chats", func(t *testing.T) {
		user := domain.NewUserWithPassword("user-no-chats", "password", 1)
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		userChats, err := chats.FindByUserID(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Empty(t, userChats)
	})

	t.Run("multiple chats", func(t *testing.T) {
		user := domain.NewUserWithPassword("user-multiple-chats", "password", 1)
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		// Create some chats.
		for i := 0; i < 3; i++ {
			chat := domain.NewChat(user, domain.NewModelID("model"))
			err := chats.Add(context.Background(), chat)
			assert.NoError(t, err)
		}

		// Call the FindByUserID method.
		userChats, err := chats.FindByUserID(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Len(t, userChats, 3)
	})

	t.Run("multiple users", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			user := domain.NewUserWithPassword(fmt.Sprintf("user-number-%d", i), "password", 1)
			err := users.Add(
				context.Background(),
				user,
			)
			assert.NoError(t, err)

			chat := domain.NewChat(user, domain.NewModelID("model"))
			err = chats.Add(context.Background(), chat)
			assert.NoError(t, err)
		}

		userChats, err := chats.FindByUserID(context.Background(), domain.NewUserID())
		assert.NoError(t, err)
		assert.Empty(t, userChats)
	})
}

func TestSqliteChats_FindByID(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUserWithPassword("username", "password", 1)
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		chat := domain.NewChat(user, domain.NewModelID("model"))
		err = chats.Add(context.Background(), chat)
		assert.NoError(t, err)

		foundChat, err := chats.FindByID(context.Background(), chat.ID)
		assert.NoError(t, err)
		assert.Equal(t, chat.ID, foundChat.ID)
		assert.Equal(t, chat.Title, foundChat.Title)
		assert.Equal(t, chat.CreatedAt, foundChat.CreatedAt)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		_, err := chats.FindByID(context.Background(), uuid.New())
		assert.Error(t, err)
	})
}

func TestSqliteChats_Exists(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	chats := NewChats(db)
	users := NewUsers(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUserWithPassword("username", "password", 1)
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		chat := domain.NewChat(user, domain.NewModelID("model"))
		err = chats.Add(context.Background(), chat)
		assert.NoError(t, err)

		exists, err := chats.Exists(context.Background(), chat.ID)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		exists, err := chats.Exists(context.Background(), uuid.New())
		assert.NoError(t, err)
		assert.False(t, exists)
	})
}

func TestSqliteChats_FindByIDWithMessages(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	modelID := domain.NewModelID("model")
	user := domain.NewUserWithPassword("username", "password", 1)
	err := users.Add(ctx, user)
	assert.NoError(err)

	chat := domain.NewChat(user, modelID)
	err = chats.Add(ctx, chat)
	assert.NoError(err)

	msgs := []domain.Message{
		domain.NewUserMessage(user, "Hello, model!"),
		domain.NewModelMessage(modelID, "Hello, user!"),
	}

	for _, message := range msgs {
		chat.AddMessage(message)
	}

	err = chats.Update(ctx, chat)
	assert.NoError(err)

	foundChat, err := chats.FindByIDWithMessages(ctx, chat.ID)
	assert.NoError(err)
	assert.Equal(len(msgs), len(foundChat.Messages))
	for i, message := range foundChat.Messages {
		assert.Equal(msgs[i], message)
	}
}
