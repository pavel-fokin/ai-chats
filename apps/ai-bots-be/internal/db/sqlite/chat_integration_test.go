package sqlite

import (
	"context"
	"fmt"
	"testing"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSqlite_CreateChat(t *testing.T) {
	// Create a new instance of Sqlite
	db, close := New(":memory:")
	defer close()

	// Create some test actors
	actors := []domain.Actor{
		{
			ID:   uuid.New(),
			Type: "User",
		},
		{
			ID:   uuid.New(),
			Type: "Bot",
		},
	}

	// Call the CreateChat method
	chat, err := db.CreateChat(context.Background(), actors)
	assert.NoError(t, err)
	assert.NotNil(t, chat)
}

func TestSqlite_CreateActor(t *testing.T) {
	// Create a new instance of Sqlite
	db, close := New(":memory:")
	defer close()

	// Call the CreateActor method
	actorType := "User"
	actor, err := db.CreateActor(context.Background(), actorType)
	assert.NoError(t, err)
	assert.NotNil(t, actor)
}

func TestAddMessages(t *testing.T) {
	// Create a new instance of Sqlite
	db, close := New(":memory:")
	defer close()

	// Create some test actors
	actors := []domain.Actor{
		{
			ID:   uuid.New(),
			Type: "User",
		},
		{
			ID:   uuid.New(),
			Type: "Bot",
		},
	}

	// Create a new chat
	chat, err := db.CreateChat(context.Background(), actors)
	assert.NoError(t, err)

	// Create some test messages
	messages := []domain.Message{
		{
			ID:    uuid.New(),
			Actor: actors[0],
			Text:  "Hello, bot!",
		},
		{
			ID:    uuid.New(),
			Actor: actors[1],
			Text:  "Hello, user!",
		},
	}

	// Add the messages to the chat
	for _, message := range messages {
		err := db.AddMessage(context.Background(), chat, message.Actor, message.Text)
		assert.NoError(t, err)
	}
}

func TestAllMessages(t *testing.T) {
	// Create a new instance of Sqlite
	db, close := New("/tmp/test.db")
	defer close()

	// Create some test actors
	actors := []domain.Actor{
		{
			ID:   uuid.New(),
			Type: "User",
		},
		{
			ID:   uuid.New(),
			Type: "Bot",
		},
	}

	// Create a new chat
	chat, err := db.CreateChat(context.Background(), actors)
	assert.NoError(t, err)

	// Create some test messages
	messages := []domain.Message{
		{
			ID:    uuid.New(),
			Actor: actors[0],
			Text:  "Hello, bot!",
		},
		{
			ID:    uuid.New(),
			Actor: actors[1],
			Text:  "Hello, user!",
		},
	}

	// Add the messages to the chat
	for _, message := range messages {
		err := db.AddMessage(context.Background(), chat, message.Actor, message.Text)
		assert.NoError(t, err)
	}

	// Call the AllMessages method
	fmt.Println(messages)
	allMessages, err := db.AllMessages(context.Background(), chat.ID)
	fmt.Println(allMessages)
	assert.NoError(t, err)
	assert.Equal(t, len(messages), len(allMessages))
}
