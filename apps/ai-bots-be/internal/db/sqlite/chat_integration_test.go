package sqlite

import (
	"context"
	"fmt"
	"testing"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateChat(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	user := domain.NewUser("test")
	err = users.AddUser(context.Background(), user)
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
	chat, err := chats.CreateChat(context.Background(), user.ID, actors)
	assert.NoError(t, err)
	assert.NotNil(t, chat)
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
			_, err := chats.CreateChat(context.Background(), user.ID, actors)
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
			user := domain.NewUser(fmt.Sprintf("test_%d", i))
			err := users.AddUser(
				context.Background(),
				user,
			)
			assert.NoError(t, err)

			_, err = chats.CreateChat(context.Background(), user.ID, actors)
			assert.NoError(t, err)
		}

		// Call the AllChats method.
		allChats, err := chats.AllChats(context.Background(), uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, allChats)
	})
}

func TestCreateActor(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	chats := NewChats(db)

	// Call the CreateActor method.
	actorType := domain.ActorType("User")
	actor, err := chats.CreateActor(context.Background(), actorType)
	assert.NoError(t, err)
	assert.NotNil(t, actor)
}

func TestAddMessages(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	user := domain.NewUser("test")
	err = users.AddUser(context.Background(), user)
	assert.NoError(t, err)

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
	chat, err := chats.CreateChat(context.Background(), user.ID, actors)
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
		err := chats.AddMessage(context.Background(), chat, message.Actor, message.Text)
		assert.NoError(t, err)
	}
}

func TestAllMessages(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	chats := NewChats(db)

	user := domain.NewUser("test")
	err = users.AddUser(context.Background(), user)
	assert.NoError(t, err)

	ai, err := chats.CreateActor(context.Background(), domain.AI)
	assert.NoError(t, err)

	human, err := chats.CreateActor(context.Background(), domain.Human)
	assert.NoError(t, err)

	chat, err := chats.CreateChat(context.Background(), user.ID, []domain.Actor{ai, human})
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
		err := chats.AddMessage(context.Background(), chat, message.Actor, message.Text)
		assert.NoError(t, err)
	}

	// Call the AllMessages method.
	allMessages, err := chats.AllMessages(context.Background(), chat.ID)
	assert.NoError(t, err)
	assert.Equal(t, len(messages), len(allMessages))
}
