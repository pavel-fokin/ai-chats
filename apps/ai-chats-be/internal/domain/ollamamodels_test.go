package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOllamaModelsFilter(t *testing.T) {
	t.Run("valid status", func(t *testing.T) {
		tests := []struct {
			status string
			want   OllamaModelStatus
		}{
			{status: "pulling", want: OllamaModelStatusPulling},
			{status: "available", want: OllamaModelStatusAvailable},
			{status: "", want: OllamaModelStatus("")},
		}

		for _, tt := range tests {
			filter, err := NewOllamaModelsFilter(tt.status)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, filter.Status)
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		tests := []string{
			"invalid",
			"any",
		}

		for _, tt := range tests {
			filter, err := NewOllamaModelsFilter(tt)
			assert.ErrorIs(t, err, ErrOllamaModelInvalidStatus)
			assert.Equal(t, OllamaModelsFilter{}, filter)
		}
	})
}
