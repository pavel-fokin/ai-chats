package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Run("empty model", func(t *testing.T) {
		model := NewModel("")
		assert.Equal(t, "", model.String())
	})

	t.Run("create a new model", func(t *testing.T) {
		model := NewModel("model:latest")

		assert.Equal(t, "model", model.Name)
		assert.Equal(t, "latest", model.Tag)
	})

	t.Run("model as string", func(t *testing.T) {
		model := NewModel("test:latest")

		assert.Equal(t, "test:latest", model.String())
	})
}
