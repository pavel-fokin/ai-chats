package sqlite

import (
	"context"
	"fmt"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/crypto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSqliteAddMessages(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)
	chats := NewChats(db)
	messages := NewMessages(db)

	user := domain.NewUser("username", "password")
	err := users.Add(ctx, user)
	assert.NoError(err)

	t.Run("success", func(t *testing.T) {
		chat := domain.NewChat(user, "model:latest")
		err := chats.Add(ctx, chat)
		assert.NoError(err)

		msgs := []domain.Message{
			domain.NewMessage("User", "Hello, bot!"),
			domain.NewMessage("AI", "Hello, user!"),
		}

		for _, message := range msgs {
			err := messages.Add(ctx, chat.ID, message)
			assert.NoError(err)
		}

		allMessages, err := messages.AllMessages(ctx, chat.ID)
		assert.NoError(err)
		assert.Equal(len(msgs), len(allMessages))
		for i, message := range allMessages {
			assert.Equal(msgs[i], message)
		}
	})

	t.Run("chat does not exist", func(t *testing.T) {
		err := messages.Add(ctx, uuid.New(), domain.NewMessage("User", "Hello, bot!"))
		assert.Error(err)
	})

	t.Run("empty text", func(t *testing.T) {
		chat := domain.NewChat(user, "model:latest")
		err := chats.Add(ctx, chat)
		assert.NoError(err)

		err = messages.Add(ctx, chat.ID, domain.NewMessage("User", ""))
		fmt.Println(err)
		assert.Error(err)
	})

	t.Run("empty sender", func(t *testing.T) {
		chat := domain.NewChat(user, "model:latest")
		err := chats.Add(ctx, chat)
		assert.NoError(err)

		err = messages.Add(ctx, chat.ID, domain.NewMessage("", "Hello, bot!"))
		fmt.Println(err)
		assert.Error(err)
	})
}

func TestSqliteAllMessages(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)
	chats := NewChats(db)
	messages := NewMessages(db)

	user := domain.NewUser("username", "password")
	err := users.Add(ctx, user)
	assert.NoError(err)

	chat := domain.NewChat(user, "model:latest")
	err = chats.Add(ctx, chat)
	assert.NoError(err)

	t.Run("success", func(t *testing.T) {
		msgs := []domain.Message{
			domain.NewMessage("User", "Hello, bot!"),
			domain.NewMessage("AI", "Hello, user!"),
		}

		for _, message := range msgs {
			err := messages.Add(ctx, chat.ID, message)
			assert.NoError(err)
		}

		allMessages, err := messages.AllMessages(ctx, chat.ID)
		assert.NoError(err)
		assert.Equal(len(msgs), len(allMessages))
		for i, message := range allMessages {
			assert.Equal(msgs[i], message)
		}
	})

	t.Run("chat does not exist", func(t *testing.T) {
		_, err := messages.AllMessages(ctx, uuid.New())
		assert.NoError(err)
	})
}
