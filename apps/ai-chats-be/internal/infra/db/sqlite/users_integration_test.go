package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"ai-chats/internal/domain"
)

func TestSqliteUsers_AddUser(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	t.Run("add user", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUserWithPassword("username", "password", 1))
		assert.NoError(t, err)
	})

	t.Run("add user with empty username", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUserWithPassword("", "password", 1))
		assert.Error(t, err)
	})

	t.Run("add user with empty password", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUserWithPassword("username", "", 1))
		assert.Error(t, err)
	})

	t.Run("add user with the same username", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUserWithPassword("another_username", "password", 1))
		assert.NoError(t, err)

		err = users.Add(context.Background(), domain.NewUserWithPassword("another_username", "password", 1))
		assert.ErrorIs(t, err, domain.ErrUserAlreadyExists)
	})

	t.Run("add user with the same username but different case", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUserWithPassword("username", "password", 1))
		assert.Error(t, err)
	})
}

func TestSqliteUsers_FindByUsernameWithPassword(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)

	t.Run("user found", func(t *testing.T) {
		user := domain.NewUserWithPassword("username", "password", 1)
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		foundUser, err := users.FindByUsernameWithPassword(context.Background(), "username")
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user.ID, foundUser.ID)
	})

	t.Run("user not found", func(t *testing.T) {
		_, err := users.FindByUsernameWithPassword(context.Background(), "unknown")
		assert.ErrorIs(t, err, domain.ErrUserNotFound)
	})
}
