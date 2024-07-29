package domain

import (
	"strings"
)

type OllamaPullingFinalStatus string

const (
	OllamaPullingFinalStatusSuccess OllamaPullingFinalStatus = "success"
	OllamaPullingFinalStatusFailed  OllamaPullingFinalStatus = "failed"
)

// OllamaModel represents an Ollama model.
type OllamaModel struct {
	Model       string `json:"model"`
	Description string `json:"description"`
	IsPulling   bool   `json:"isPulling"`
}

func NewOllamaModel(model, description string) OllamaModel {
	return OllamaModel{
		Model:       model,
		Description: description,
	}
}

func (om OllamaModel) String() string {
	return om.Model
}

func (om OllamaModel) Name() string {
	parts := strings.Split(om.Model, ":")
	if len(parts) == 2 {
		return parts[0]
	}
	return om.Model
}

// func (om *OllamaModel) Scan(value any) error {
// 	*om = *NewOllamaModel(value.(string))
// 	return nil
// }

// func (om OllamaModel) Value() (any, error) {
// 	return om.String(), nil
// }
