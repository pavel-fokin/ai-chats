package sqlite

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, close := New(":memory:")
	defer close()

	user := domain.NewUser("username")

	err := db.AddUser(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindUser(t *testing.T) {
	db, close := New(":memory:")
	defer close()

	user := domain.NewUser("username")
	err := db.AddUser(context.Background(), user)
	assert.NoError(t, err)

	foundUser, err := db.FindUser(context.Background(), "username")
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
}
