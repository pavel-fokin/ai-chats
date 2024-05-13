package sqlite

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddMessages(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)
	messages := NewMessages(db)

	user := domain.NewUser("username")
	err = users.AddUser(context.Background(), user)
	assert.NoError(t, err)

	// Create a new chat.
	chat := domain.NewChat(user)
	err = chats.Add(context.Background(), chat)
	assert.NoError(t, err)

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
		err := messages.Add(context.Background(), chat, message)
		assert.NoError(t, err)
	}
}
