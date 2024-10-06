package api

import (
	"errors"
	"net/url"
)

// OllamaModelsQuery represents the query parameters for Ollama models.
type OllamaModelsQuery struct {
	OnlyPulling bool
}

// ParseOllamaModelsQuery parses the query string and returns the OllamaModelsQuery.
func ParseOllamaModelsQuery(query string) (OllamaModelsQuery, error) {
	values, err := url.ParseQuery(query)
	if err != nil {
		return OllamaModelsQuery{}, err
	}

	if !values.Has("onlyPulling") {
		return OllamaModelsQuery{}, nil
	}

	pulling := values.Get("onlyPulling")
	if pulling != "" {
		return OllamaModelsQuery{}, errors.New("invalid value for 'pulling' parameter")
	}

	return OllamaModelsQuery{OnlyPulling: true}, nil
}
