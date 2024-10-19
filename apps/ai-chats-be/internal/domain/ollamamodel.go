package domain

import (
	"strings"
)

type OllamaPullingFinalStatus string

const (
	OllamaPullingFinalStatusSuccess OllamaPullingFinalStatus = "success"
	OllamaPullingFinalStatusFailed  OllamaPullingFinalStatus = "failed"
)

// OllamaModelStatus is the status of an Ollama model.
type OllamaModelStatus string

const (
	OllamaModelStatusPulling   OllamaModelStatus = "pulling"
	OllamaModelStatusAvailable OllamaModelStatus = "available"
)

// OllamaModel represents an Ollama model.
type OllamaModel struct {
	Model       string            `json:"model"`
	Description string            `json:"description"`
	Status      OllamaModelStatus `json:"status"`
}

func NewOllamaModel(model string) OllamaModel {
	return OllamaModel{
		Model: model,
	}
}

func (om *OllamaModel) SetStatus(status OllamaModelStatus) {
	om.Status = status
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
