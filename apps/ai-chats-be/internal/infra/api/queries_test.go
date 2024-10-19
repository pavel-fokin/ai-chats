package api

import (
	"ai-chats/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOllamaModelsQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		want    domain.OllamaModelsFilter
		wantErr bool
	}{
		{
			name:  "empty query",
			query: "",
			want:  domain.OllamaModelsFilter{},
		},
		{
			name:  "status=",
			query: "status=",
			want:  domain.OllamaModelsFilter{},
		},
		{
			name:  "pulling",
			query: "status=pulling",
			want: domain.OllamaModelsFilter{
				Status: domain.OllamaModelStatusPulling,
			},
		},
		{
			name:  "available",
			query: "status=available",
			want: domain.OllamaModelsFilter{
				Status: domain.OllamaModelStatusAvailable,
			},
		},
		{
			name:    "invalid status",
			query:   "status=any",
			wantErr: true,
		},
		{
			name:    "invalid status",
			query:   "status=invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOllamaModelsQuery(tt.query)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
