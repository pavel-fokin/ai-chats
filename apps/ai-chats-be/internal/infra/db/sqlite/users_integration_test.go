package sqlite

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteAddUser(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)
	t.Run("Add user", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("username"))
		assert.NoError(t, err)
	})

	t.Run("Add user with the same username", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("another_username"))
		assert.NoError(t, err)

		err = users.Add(context.Background(), domain.NewUser("another_username"))
		// assert.NoError(t, err)
		assert.ErrorIs(t, err, domain.ErrUserAlreadyExists)
		// assert.Error(t, err)
	})

	t.Run("Add user with the same username but different case", func(t *testing.T) {
		err := users.Add(context.Background(), domain.NewUser("username"))
		assert.Error(t, err)
	})
}

func TestFindByUsernameWithPassword(t *testing.T) {
	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	users := NewUsers(db)

	t.Run("User found", func(t *testing.T) {
		user := domain.NewUser("username")
		err := users.Add(context.Background(), user)
		assert.NoError(t, err)

		foundUser, err := users.FindByUsernameWithPassword(context.Background(), "username")
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user.ID, foundUser.ID)
	})

	t.Run("User not found", func(t *testing.T) {
		_, err := users.FindByUsernameWithPassword(context.Background(), "unknown")
		assert.Error(t, err)
	})
}
