package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"ai-chats/internal/domain"
)

func TestSqliteModels_AddModelCard(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	models := NewModels(db)

	t.Run("valid", func(t *testing.T) {
		err := models.AddModelCard(ctx, domain.ModelCard{
			Model:       "model",
			Description: "description",
		})
		assert.NoError(t, err)
	})

	t.Run("add description with empty model", func(t *testing.T) {
		err := models.AddModelCard(ctx, domain.ModelCard{
			Model:       "",
			Description: "description",
		})
		assert.Error(t, err)
	})

	t.Run("add description with empty description", func(t *testing.T) {
		err := models.AddModelCard(ctx, domain.ModelCard{
			Model:       "model",
			Description: "",
		})
		assert.Error(t, err)
	})
}

func TestSqliteModels_AllModelCards(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	models := NewModels(db)

	t.Run("empty", func(t *testing.T) {
		descriptions, err := models.AllModelCards(ctx)
		assert.NoError(t, err)
		assert.Empty(t, descriptions)
	})

	t.Run("valid", func(t *testing.T) {
		err := models.AddModelCard(ctx, domain.ModelCard{
			Model:       "model",
			Description: "description",
		})
		assert.NoError(t, err)

		descriptions, err := models.AllModelCards(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, descriptions)
		assert.Equal(t, 1, len(descriptions))
	})
}

func TestSqliteModels_FindModelCard(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	models := NewModels(db)

	t.Run("not found", func(t *testing.T) {
		description, err := models.FindModelCard(ctx, "model")
		assert.Error(t, err)
		assert.Equal(t, domain.ModelCard{}, description)
	})

	t.Run("found", func(t *testing.T) {
		err := models.AddModelCard(ctx, domain.ModelCard{
			Model:       "model",
			Description: "description",
		})
		assert.NoError(t, err)

		description, err := models.FindModelCard(ctx, "model")
		assert.NoError(t, err)
		assert.Equal(t, domain.ModelCard{
			Model:       "model",
			Description: "description",
		}, description)
	})
}
