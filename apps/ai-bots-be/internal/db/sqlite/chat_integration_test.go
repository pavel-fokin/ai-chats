package sqlite

import (
	"context"
	"fmt"
	"testing"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateChat(t *testing.T) {
	db, close := New(":memory:")
	defer close()

	user, err := db.CreateUser(context.Background(), "test", "test")
	assert.NoError(t, err)

	// Create some test actors.
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

	// Call the CreateChat method.
	chat, err := db.CreateChat(context.Background(), user.ID, actors)
	assert.NoError(t, err)
	assert.NotNil(t, chat)
}

func TestAllChats(t *testing.T) {
	t.Run("no chats", func(t *testing.T) {
		db, close := New(":memory:")
		defer close()

		user, err := db.CreateUser(context.Background(), "test", "test")
		assert.NoError(t, err)

		// Call the AllChats method.
		chats, err := db.AllChats(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Empty(t, chats)
	})

	t.Run("multiple chats", func(t *testing.T) {
		db, close := New(":memory:")
		defer close()

		user, err := db.CreateUser(context.Background(), "test", "test")
		assert.NoError(t, err)

		// Create some test actors.
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

		// Create some chats.
		for i := 0; i < 3; i++ {
			_, err := db.CreateChat(context.Background(), user.ID, actors)
			assert.NoError(t, err)
		}

		// Call the AllChats method.
		chats, err := db.AllChats(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Len(t, chats, 3)
	})

	t.Run("multiple users", func(t *testing.T) {
		db, close := New(":memory:")
		defer close()

		// Create some test actors.
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

		// Create some chats.
		for i := 0; i < 3; i++ {
			user, err := db.CreateUser(
				context.Background(),
				fmt.Sprintf("username_%d", i),
				"password",
			)
			assert.NoError(t, err)

			_, err = db.CreateChat(context.Background(), user.ID, actors)
			assert.NoError(t, err)
		}

		// Call the AllChats method.
		chats, err := db.AllChats(context.Background(), uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, chats)
	})
}

func TestCreateActor(t *testing.T) {
	db, close := New(":memory:")
	defer close()

	// Call the CreateActor method.
	actorType := domain.ActorType("User")
	actor, err := db.CreateActor(context.Background(), actorType)
	assert.NoError(t, err)
	assert.NotNil(t, actor)
}

func TestAddMessages(t *testing.T) {
	db, close := New(":memory:")
	defer close()

	// Create a new user.
	user, err := db.CreateUser(context.Background(), "test", "test")
	assert.NoError(t, err)

	// Create some test actors.
	actors := []domain.Actor{
		{
			ID:   uuid.New(),
			Type: domain.Human,
		},
		{
			ID:   uuid.New(),
			Type: domain.AI,
		},
	}

	// Create a new chat.
	chat, err := db.CreateChat(context.Background(), user.ID, actors)
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
	// Create a new instance of Sqlite.
	db, close := New(":memory:")
	defer close()

	// Create a new user.
	user, err := db.CreateUser(context.Background(), "test", "test")
	assert.NoError(t, err)

	// Create some test actors.
	ai, err := db.CreateActor(context.Background(), domain.AI)
	assert.NoError(t, err)

	human, err := db.CreateActor(context.Background(), domain.Human)
	assert.NoError(t, err)

	// Create a new chat.
	chat, err := db.CreateChat(context.Background(), user.ID, []domain.Actor{ai, human})
	assert.NoError(t, err)

	// Create some test messages.
	messages := []domain.Message{
		{
			ID:    uuid.New(),
			Actor: human,
			Text:  "Hello, bot!",
		},
		{
			ID:    uuid.New(),
			Actor: ai,
			Text:  "Hello, user!",
		},
	}

	// Add the messages to the chat.
	for _, message := range messages {
		err := db.AddMessage(context.Background(), chat, message.Actor, message.Text)
		assert.NoError(t, err)
	}

	// Call the AllMessages method.
	allMessages, err := db.AllMessages(context.Background(), chat.ID)
	assert.NoError(t, err)
	assert.Equal(t, len(messages), len(allMessages))
}
