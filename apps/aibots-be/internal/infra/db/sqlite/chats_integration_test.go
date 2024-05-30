package sqlite

import (
	"context"
	"fmt"
	"testing"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddChat(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	user := domain.NewUser("test")
	err = users.Add(context.Background(), user)
	assert.NoError(t, err)

	// Call the AddChat method.
	chat := domain.NewChat(user)
	err = chats.Add(context.Background(), chat)
	assert.NoError(t, err)
}

func TestAllChats(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("no chats", func(t *testing.T) {
		user := domain.NewUser("user-no-chats")
		err = users.Add(context.Background(), user)
		assert.NoError(t, err)

		// Call the AllChats method.
		allChats, err := chats.AllChats(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Empty(t, allChats)
	})

	t.Run("multiple chats", func(t *testing.T) {
		user := domain.NewUser("user-multiple-chats")
		err = users.Add(context.Background(), user)
		assert.NoError(t, err)

		// Create some chats.
		for i := 0; i < 3; i++ {
			chat := domain.NewChat(user)
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
			user := domain.NewUser(fmt.Sprintf("user-number-%d", i))
			err := users.Add(
				context.Background(),
				user,
			)
			assert.NoError(t, err)

			chat := domain.NewChat(user)
			err = chats.Add(context.Background(), chat)
			assert.NoError(t, err)
		}

		allChats, err := chats.AllChats(context.Background(), uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, allChats)
	})
}

func TestFindChat(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username")
		err = users.Add(context.Background(), user)
		assert.NoError(t, err)

		chat := domain.NewChat(user)
		err = chats.Add(context.Background(), chat)
		assert.NoError(t, err)

		foundChat, err := chats.FindChat(context.Background(), chat.ID)
		assert.NoError(t, err)
		assert.Equal(t, chat.ID, foundChat.ID)
		assert.Equal(t, chat.Title, foundChat.Title)
		assert.Equal(t, chat.CreatedAt, foundChat.CreatedAt)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		_, err := chats.FindChat(context.Background(), uuid.New())
		assert.Error(t, err)
	})
}

func TestChats_Exists(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	chats := NewChats(db)
	users := NewUsers(db)

	t.Run("chat exists", func(t *testing.T) {
		user := domain.NewUser("username")
		err = users.Add(context.Background(), user)
		assert.NoError(t, err)

		chat := domain.NewChat(user)
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
