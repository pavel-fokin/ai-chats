package domain

import (
	"testing"

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
		chat := NewChat(NewUser("username"), NewModelID("model"))
		assert.Equal(t, 0, len(chat.Events))
		message := NewMessage(NewSender("user:1"), "Hello, world!")

		chat.AddMessage(message)

		assert.Equal(t, 1, len(chat.Messages))
		assert.Equal(t, message, chat.Messages[0])
	})

	t.Run("update title", func(t *testing.T) {
		chat := NewChat(NewUser("username"), NewModelID("model"))
		assert.Equal(t, 0, len(chat.Events))
		assert.Equal(t, "New chat", chat.Title)

		chat.UpdateTitle("Test title")
		assert.Equal(t, 1, len(chat.Events))
		assert.Equal(t, "Test title", chat.Title)
	})

	t.Run("can user access", func(t *testing.T) {
		user := NewUser("username")

		chat := NewChat(user, NewModelID("model"))
		assert.NoError(t, chat.CanUserAccess(user.ID))
		assert.ErrorIs(t, chat.CanUserAccess(NewUserID()), ErrChatAccessDenied)
	})
}
