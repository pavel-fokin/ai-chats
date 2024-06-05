package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	t.Run("NewModel", func(t *testing.T) {
		name := "test"
		model := NewModel(name)

		assert.Equal(t, name, model.Name)
		assert.NotEmpty(t, model.ID)
	})

	t.Run("AddTag", func(t *testing.T) {
		model := NewModel("test")
		tag := NewModelTag("tag")

		model.AddTag(tag)

		assert.Len(t, model.Tags, 1)
		assert.Equal(t, tag, model.Tags[0])
	})

	t.Run("RemoveTag", func(t *testing.T) {
		model := NewModel("test")
		tag := NewModelTag("tag")

		model.AddTag(tag)

		err := model.RemoveTag(tag)

		assert.Nil(t, err)
		assert.Len(t, model.Tags, 0)
	})

	t.Run("RemoveTagNotFound", func(t *testing.T) {
		model := NewModel("test")
		tag := NewModelTag("tag")

		err := model.RemoveTag(tag)

		assert.Equal(t, ErrTagNotFound, err)
	})
}
