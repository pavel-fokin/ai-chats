package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Run("NewModel", func(t *testing.T) {
		name := "test"
		tag := "latest"
		model := NewModel(name, tag)

		assert.Equal(t, name, model.Name)
		assert.Equal(t, tag, model.Tag)
		// assert.NotEmpty(t, model.ID)
	})
}
