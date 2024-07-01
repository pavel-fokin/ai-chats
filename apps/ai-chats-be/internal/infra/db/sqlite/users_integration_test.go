package sqlite

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := NewDB(":memory:")
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)

	err := users.Add(context.Background(), domain.NewUser("username"))
	assert.NoError(t, err)
}

func TestFindByUsernameWithPassword(t *testing.T) {
	db := NewDB(":memory:")
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
