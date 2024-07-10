package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/crypto"
)

func TestSqliteAddUser(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)
	t.Run("add user", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("username", "password"))
		assert.NoError(t, err)
	})

	t.Run("add user with empty username", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("", "password"))
		assert.Error(t, err)
	})

	t.Run("add user with empty password", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("username", ""))
		assert.Error(t, err)
	})

	t.Run("add user with the same username", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("another_username", "password"))
		assert.NoError(t, err)

		err = users.Add(context.Background(), domain.NewUser("another_username", "password"))
		assert.ErrorIs(t, err, domain.ErrUserAlreadyExists)
	})

	t.Run("add user with the same username but different case", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("username", "password"))
		assert.Error(t, err)
	})
}

func TestSqliteFindByUsernameWithPassword(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	crypto.InitBcryptCost(1)

	users := NewUsers(db)

	t.Run("user found", func(t *testing.T) {
		user := domain.NewUser("username", "password")
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		foundUser, err := users.FindByUsernameWithPassword(context.Background(), "username")
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user.ID, foundUser.ID)
	})

	t.Run("user not found", func(t *testing.T) {
		_, err := users.FindByUsernameWithPassword(context.Background(), "unknown")
		assert.Error(t, err)
	})
}
