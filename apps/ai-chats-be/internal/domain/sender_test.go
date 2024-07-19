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
	model := NewModel("llama3")
	sender := NewModelSender(model)
	expected := Sender{sender: "model:llama3"}
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
func TestSender_Format_User(t *testing.T) {
	sender := Sender{sender: "user:123"}
	expected := "User"
	assert.Equal(t, expected, sender.Format())
}

func TestSender_Format_Model(t *testing.T) {
	sender := Sender{sender: "model:latest"}
	expected := "Model (latest)"
	assert.Equal(t, expected, sender.Format())
}

func TestSender_Format_Invalid(t *testing.T) {
	sender := Sender{sender: "invalid"}
	expected := ""
	assert.Equal(t, expected, sender.Format())
}
func TestSender_MarshalJSON(t *testing.T) {
	sender := Sender{sender: "user:123"}
	expected := []byte(`"user:123"`)
	actual, err := sender.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSender_MarshalJSON_Model(t *testing.T) {
	sender := Sender{sender: "model:llama3:latest"}
	expected := []byte(`"model:llama3:latest"`)
	actual, err := sender.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSender_MarshalJSON_Invalid(t *testing.T) {
	sender := Sender{sender: "invalid"}
	expected := []byte(`"invalid"`)
	actual, err := sender.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSender_UnmarshalJSON(t *testing.T) {
	data := []byte(`"user:123"`)
	expected := Sender{sender: "user:123"}
	var actual Sender
	err := actual.UnmarshalJSON(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSender_UnmarshalJSON_Model(t *testing.T) {
	data := []byte(`"model:llama3:latest"`)
	expected := Sender{sender: "model:llama3:latest"}
	var actual Sender
	err := actual.UnmarshalJSON(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSender_UnmarshalJSON_Invalid(t *testing.T) {
	data := []byte(`"invalid"`)
	expected := Sender{sender: "invalid"}
	var actual Sender
	err := actual.UnmarshalJSON(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
