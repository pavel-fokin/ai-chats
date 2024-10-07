package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOllamaModelsQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		want    OllamaModelsQuery
		wantErr bool
	}{
		{
			name:    "empty query",
			query:   "",
			want:    OllamaModelsQuery{},
			wantErr: false,
		},
		{
			name:    "onlyPulling parameter present",
			query:   "onlyPulling",
			want:    OllamaModelsQuery{OnlyPulling: true},
			wantErr: false,
		},
		{
			name:    "onlyPulling parameter present",
			query:   "onlyPulling=",
			want:    OllamaModelsQuery{OnlyPulling: true},
			wantErr: false,
		},
		{
			name:    "onlyPulling with value",
			query:   "onlyPulling=true",
			want:    OllamaModelsQuery{},
			wantErr: true,
		},
		{
			name:    "onlyPulling with false value",
			query:   "onlyPulling=false",
			want:    OllamaModelsQuery{},
			wantErr: true,
		},
		{
			name:    "other parameter",
			query:   "other=param",
			want:    OllamaModelsQuery{},
			wantErr: false,
		},
		{
			name:    "invalid query string",
			query:   "invalid=query&string",
			want:    OllamaModelsQuery{},
			wantErr: false,
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
