package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainOllamaModel(t *testing.T) {
	t.Run("empty model", func(t *testing.T) {
		model, err := NewOllamaModel("")
		assert.ErrorIs(t, err, ErrOllamaModelEmpty)
		assert.Equal(t, OllamaModel{}, model)
	})

	t.Run("create a new model", func(t *testing.T) {
		model, err := NewOllamaModel("model")
		assert.NoError(t, err)
		assert.Equal(t, "model", model.Model)
		assert.Equal(t, "model", model.Name)
		assert.Equal(t, "latest", model.Tag)
	})

	t.Run("model with tag", func(t *testing.T) {
		model, err := NewOllamaModel("model:tag")
		assert.NoError(t, err)
		assert.Equal(t, "model:tag", model.Model)
		assert.Equal(t, "model", model.Name)
		assert.Equal(t, "tag", model.Tag)
	})

	t.Run("model as string", func(t *testing.T) {
		model, err := NewOllamaModel("model:latest")
		assert.NoError(t, err)
		assert.Equal(t, "model:latest", model.String())
	})
}
