package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Run("create a new model", func(t *testing.T) {
		name := "test"
		tag := "latest"
		model := NewModel(name, tag)

		assert.Equal(t, name, model.Name)
		assert.Equal(t, tag, model.Tag)
	})

	t.Run("model as string", func(t *testing.T) {
		name := "test"
		tag := "latest"
		model := NewModel(name, tag)

		assert.Equal(t, "test:latest", model.String())
	})
}
