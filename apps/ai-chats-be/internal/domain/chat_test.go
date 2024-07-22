package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDomainChat(t *testing.T) {

	t.Run("new chat", func(t *testing.T) {
		user := NewUser("username")
		modelID := NewModelID("model")

		chat := NewChat(user, modelID)

		assert.NotEmpty(t, chat.ID)
		assert.Equal(t, "New chat", chat.Title)
		assert.Equal(t, user, chat.User)
		assert.Equal(t, modelID, chat.DefaultModel)
		assert.NotEmpty(t, chat.CreatedAt)
	})

	t.Run("add message", func(t *testing.T) {
		chat := Chat{
			ID:           NewChatID(),
			Title:        "New chat",
			User:         NewUser("username"),
			DefaultModel: NewModelID("model"),
			CreatedAt:    time.Now().UTC(),
		}

		message := NewMessage(NewSender("user:1"), "Hello, world!")

		chat.AddMessage(message)

		assert.Equal(t, 1, len(chat.Messages))
		assert.Equal(t, message, chat.Messages[0])
	})
}
