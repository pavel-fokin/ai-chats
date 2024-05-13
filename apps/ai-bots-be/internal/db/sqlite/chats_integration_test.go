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
	err = users.AddUser(context.Background(), user)
	assert.NoError(t, err)

	// Call the AddChat method.
	chat := domain.NewChat(user)
	err = chats.Add(context.Background(), chat)
	assert.NoError(t, err)
}

func TestAllChats(t *testing.T) {
	t.Run("no chats", func(t *testing.T) {
		db, err := NewDB(":memory:")
		assert.NoError(t, err)
		defer db.Close()
		CreateTables(db)

		users := NewUsers(db)
		chats := NewChats(db)

		user := domain.NewUser("test")
		err = users.AddUser(context.Background(), user)
		assert.NoError(t, err)

		// Call the AllChats method.
		allChats, err := chats.AllChats(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Empty(t, allChats)
	})

	t.Run("multiple chats", func(t *testing.T) {
		db, err := NewDB(":memory:")
		assert.NoError(t, err)
		defer db.Close()
		CreateTables(db)

		users := NewUsers(db)
		chats := NewChats(db)

		user := domain.NewUser("test")
		err = users.AddUser(context.Background(), user)
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
		db, err := NewDB(":memory:")
		assert.NoError(t, err)
		defer db.Close()
		CreateTables(db)

		users := NewUsers(db)
		chats := NewChats(db)

		// Create some chats.
		for i := 0; i < 3; i++ {
			user := domain.NewUser(fmt.Sprintf("test_%d", i))
			err := users.AddUser(
				context.Background(),
				user,
			)
			assert.NoError(t, err)

			chat := domain.NewChat(user)
			err = chats.Add(context.Background(), chat)
			assert.NoError(t, err)
		}

		// Call the AllChats method.
		allChats, err := chats.AllChats(context.Background(), uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, allChats)
	})
}