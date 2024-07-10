package sqlite

import (
	"context"
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

	// Create a new chat.
	chat := domain.NewChat(user, "model:latest")
	err = chats.Add(ctx, chat)
	assert.NoError(err)

	// Create some test messages
	msgs := []domain.Message{
		{
			ID:     uuid.New(),
			Sender: "User",
			Text:   "Hello, bot!",
		},
		{
			ID:     uuid.New(),
			Sender: "AI",
			Text:   "Hello, user!",
		},
	}

	// Add the messages to the chat
	for _, message := range msgs {
		err := messages.Add(ctx, chat.ID, message)
		assert.NoError(err)
	}

	// Retrieve all messages for the chat
	allMessages, err := messages.AllMessages(ctx, chat.ID)
	assert.NoError(err)
	assert.Equal(len(msgs), len(allMessages))
	for i, message := range allMessages {
		assert.Equal(msgs[i].ID, message.ID)
		assert.Equal(msgs[i].Sender, message.Sender)
		assert.Equal(msgs[i].Text, message.Text)
	}
}
