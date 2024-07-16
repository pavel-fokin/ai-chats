package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserSender(t *testing.T) {
	id := uuid.New()
	user := User{ID: id}
	sender := NewUserSender(user)
	expected := Sender{sender: "user:" + id.String()}
	assert.Equal(t, expected, sender)
}

func TestNewModelSender(t *testing.T) {
	model := NewModel("model")
	sender := NewModelSender(model)
	expected := Sender{sender: "model:latest"}
	assert.Equal(t, expected, sender)
}

func TestSender_String(t *testing.T) {
	sender := Sender{sender: "user:123"}
	expected := "user:123"
	assert.Equal(t, expected, sender.String())
}

func TestSender_IsUser(t *testing.T) {
	sender := Sender{sender: "user:123"}
	assert.True(t, sender.IsUser())

	sender = Sender{sender: "model:latest"}
	assert.False(t, sender.IsUser())
}

func TestSender_IsModel(t *testing.T) {
	sender := Sender{sender: "user:123"}
	assert.False(t, sender.IsModel())

	sender = Sender{sender: "model:latest"}
	assert.True(t, sender.IsModel())
}
