package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"ai-chats/internal/domain"
)

func TestSqliteModelsLibrary_FindAll(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	LoadFixtures(db)

	t.Run("valid", func(t *testing.T) {
		modelsLibrary := NewModelsLibrary(db)
		modelCards, err := modelsLibrary.FindAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, modelCards, 2)
	})
}

func TestSqliteModelsLibrary_FindByName(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)
	LoadFixtures(db)

	modelsLibrary := NewModelsLibrary(db)
	t.Run("valid", func(t *testing.T) {
		modelCard, err := modelsLibrary.FindByName(ctx, "llama3")
		assert.NotNil(t, modelCard)
		assert.NoError(t, err)

		assert.Equal(t, "llama3", modelCard.Model)
		assert.Equal(t, "Meta Llama 3: The most capable openly available LLM to date.", modelCard.Description)
		assert.Equal(t, []string{"70b", "8b"}, modelCard.Tags)
	})

	t.Run("not found", func(t *testing.T) {
		modelCard, err := modelsLibrary.FindByName(ctx, "not found")
		assert.Nil(t, modelCard)
		assert.ErrorIs(t, err, domain.ErrModelNotFound)
	})
}
