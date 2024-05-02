package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, close := New(":memory:")
	defer close()

	user, err := db.CreateUser(context.Background(), "username", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindUser(t *testing.T) {
	db, close := New(":memory:")
	defer close()

	user, err := db.CreateUser(context.Background(), "username", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	foundUser, err := db.FindUser(context.Background(), "username")
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
}
