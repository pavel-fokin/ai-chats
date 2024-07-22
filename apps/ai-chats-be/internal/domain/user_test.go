package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainUser(t *testing.T) {
	t.Run("new user", func(t *testing.T) {
		user := NewUser("username")

		assert.NotEmpty(t, user.ID)
		assert.Equal(t, "username", user.Username)
	})

	t.Run("new user with empty username", func(t *testing.T) {
		user := NewUser("")

		assert.NotEmpty(t, user.ID)
		assert.Empty(t, user.Username)
	})

	t.Run("new user with password", func(t *testing.T) {
		user := NewUserWithPassword("username", "password", 1)

		assert.NotEmpty(t, user.ID)
		assert.Equal(t, "username", user.Username)
		assert.NotEmpty(t, user.PasswordHash)
	})
}
