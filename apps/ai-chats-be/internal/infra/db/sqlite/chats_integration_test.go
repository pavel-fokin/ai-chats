package sqlite

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/crypto"
)

func TestSqliteAddChat(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)
	chats := NewChats(db)

	user := domain.NewUser("test", "password")
	err := users.Add(ctx, user)
	assert.NoError(t, err)

	t.Run("valid", func(t *testing.T) {
		chat := domain.NewChat(user, "model:latest")
		err = chats.Add(ctx, chat)
		assert.NoError(t, err)
	})

	t.Run("add chat without user", func(t *testing.T) {
		chat := domain.NewChat(domain.User{}, "model:latest")
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})

	t.Run("add chat with empty title", func(t *testing.T) {
		chat := domain.NewChat(user, "model:latest")
		chat.Title = ""
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})

	t.Run("add chat with invalid user", func(t *testing.T) {
		chat := domain.NewChat(domain.User{ID: uuid.New()}, "model:latest")
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})

	t.Run("add chat with empty model", func(t *testing.T) {
		chat := domain.NewChat(user, "")
		err = chats.Add(ctx, chat)
		assert.Error(t, err)
	})
}

func TestSqliteDeleteChat(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username", "password")
		err := users.Add(ctx, user)
		assert.NoError(err)

		chat := domain.NewChat(user, "model:latest")
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

func TestSqliteAllChats(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("no chats", func(t *testing.T) {
		user := domain.NewUser("user-no-chats", "password")
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		// Call the AllChats method.
		allChats, err := chats.AllChats(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Empty(t, allChats)
	})

	t.Run("multiple chats", func(t *testing.T) {
		user := domain.NewUser("user-multiple-chats", "password")
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		// Create some chats.
		for i := 0; i < 3; i++ {
			chat := domain.NewChat(user, "model:latest")
			err := chats.Add(context.Background(), chat)
			assert.NoError(t, err)
		}

		// Call the AllChats method.
		allChats, err := chats.AllChats(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Len(t, allChats, 3)
	})

	t.Run("multiple users", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			user := domain.NewUser(fmt.Sprintf("user-number-%d", i), "password")
			err := users.Add(
				context.Background(),
				user,
			)
			assert.NoError(t, err)

			chat := domain.NewChat(user, "model:latest")
			err = chats.Add(context.Background(), chat)
			assert.NoError(t, err)
		}

		allChats, err := chats.AllChats(context.Background(), uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, allChats)
	})
}

func TestSqliteFindChat(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username", "password")
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		chat := domain.NewChat(user, "model:latest")
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
	crypto.InitBcryptCost(1)

	chats := NewChats(db)
	users := NewUsers(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username", "password")
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		chat := domain.NewChat(user, "model:latest")
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
