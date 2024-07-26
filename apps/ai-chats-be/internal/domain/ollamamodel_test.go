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
		assert.Equal(t, StatusAdded, model.Status)
		assert.NotEmpty(t, model.AddedAt)
		assert.NotEmpty(t, model.UpdatedAt)
		assert.Empty(t, model.DeletedAt)
	})

	t.Run("model as string", func(t *testing.T) {
		model := NewOllamaModel("model:latest")

		assert.Equal(t, "model:latest", model.String())
	})

	t.Run("pull model", func(t *testing.T) {
		model := NewOllamaModel("model:latest")
		model.Pull()

		assert.Equal(t, StatusPulling, model.Status)
	})

	// t.Run("delete model", func(t *testing.T) {
	// 	model := NewOllamaModel("model:latest")
	// 	model.Delete()

	// 	assert.Equal(t, StatusDeleted, model.Status)
	// 	assert.NotEmpty(t, model.DeletedAt)
	// })
}
