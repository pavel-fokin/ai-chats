package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainOllamaModel(t *testing.T) {
	t.Run("empty model", func(t *testing.T) {
		model := NewOllamaModel("")
		assert.Equal(t, "", model.Model)
	})

	t.Run("create a new model", func(t *testing.T) {
		model := NewOllamaModel("model:latest")
		assert.Equal(t, "model:latest", model.Model)
	})

	t.Run("model as string", func(t *testing.T) {
		model := NewOllamaModel("model:latest")
		assert.Equal(t, "model:latest", model.String())
	})
}
