package domain

import (
	"errors"
	"strings"
)

var ErrOllamaModelEmpty = errors.New("model cannot be empty")

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

	Events []OllamaModelPullEvent
}

// NewOllamaModel creates a new OllamaModel.
func NewOllamaModel(model string) (OllamaModel, error) {
	if model == "" {
		return OllamaModel{}, ErrOllamaModelEmpty
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

func (om *OllamaModel) PullStarted() {
	om.Events = append(om.Events, NewOllamaModelPullStarted(om.Model))
}

func (om *OllamaModel) PullCompleted() {
	om.Events = append(om.Events, NewOllamaModelPullCompleted(om.Model))
}

func (om *OllamaModel) PullFailed() {
	om.Events = append(om.Events, NewOllamaModelPullFailed(om.Model))
}

func (om *OllamaModel) ClearEvents() {
	om.Events = []OllamaModelPullEvent{}
}

func (om OllamaModel) String() string {
	return om.Model
}
