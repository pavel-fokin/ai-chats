package domain

import (
	"errors"
	"strings"
)

var ErrOllamaModelInvalidModel = errors.New("model cannot be empty")

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
	Name        string            `json:"name"`
	Tag         string            `json:"tag"`
	Description string            `json:"description"`
	Status      OllamaModelStatus `json:"status"`
}

func NewOllamaModel(model string) (OllamaModel, error) {
	if model == "" {
		return OllamaModel{}, ErrOllamaModelInvalidModel
	}

	parts := strings.Split(model, ":")
	if len(parts) == 2 {
		return OllamaModel{
			Model: model,
			Name:  parts[0],
			Tag:   parts[1],
		}, nil
	}

	return OllamaModel{
		Model: model,
		Name:  parts[0],
		Tag:   "latest",
	}, nil
}

func (om *OllamaModel) SetStatus(status OllamaModelStatus) {
	om.Status = status
}

func (om OllamaModel) String() string {
	return om.Model
}

// func (om OllamaModel) Name() string {
// 	parts := strings.Split(om.Model, ":")
// 	if len(parts) == 2 {
// 		return parts[0]
// 	}
// 	return om.Model
// }

// func (om *OllamaModel) Scan(value any) error {
// 	*om = *NewOllamaModel(value.(string))
// 	return nil
// }

// func (om OllamaModel) Value() (any, error) {
// 	return om.String(), nil
// }
