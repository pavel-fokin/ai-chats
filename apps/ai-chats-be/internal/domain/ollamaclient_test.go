package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainOllama_OllamaClientModel(t *testing.T) {
	t.Run("create new", func(t *testing.T) {
		model := NewOllamaClientModel("model:latest")
		assert.Equal(t, "model:latest", model.Model)
	})

	t.Run("name", func(t *testing.T) {
		model := NewOllamaClientModel("model:latest")
		assert.Equal(t, "model", model.Name())
	})

	t.Run("name with empty model", func(t *testing.T) {
		model := NewOllamaClientModel("")
		assert.Equal(t, "", model.Name())
	})
}
