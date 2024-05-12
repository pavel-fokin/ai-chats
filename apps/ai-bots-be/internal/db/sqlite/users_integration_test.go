package sqlite

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)

	user := domain.NewUser("username")

	err = users.AddUser(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindUser(t *testing.T) {
	db, err := NewDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()
	CreateTables(db)

	users := NewUsers(db)

	user := domain.NewUser("username")
	err = users.AddUser(context.Background(), user)
	assert.NoError(t, err)

	foundUser, err := users.FindUser(context.Background(), "username")
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
}
