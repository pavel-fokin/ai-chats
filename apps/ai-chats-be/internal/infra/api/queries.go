package api

import (
	"ai-chats/internal/domain"
	"fmt"
	"net/url"
)

// ParseOllamaModelsQuery parses the query string and returns the OllamaModelsFilter.
func ParseOllamaModelsQuery(query string) (domain.OllamaModelsFilter, error) {
	values, err := url.ParseQuery(query)
	if err != nil {
		return domain.OllamaModelsFilter{}, fmt.Errorf("failed to parse query string: %w", err)
	}

	status := values.Get("status")

	filter, err := domain.NewOllamaModelsFilter(status)
	if err != nil {
		return domain.OllamaModelsFilter{}, fmt.Errorf("failed to create ollama models filter: %w", err)
	}

	return filter, nil
}
