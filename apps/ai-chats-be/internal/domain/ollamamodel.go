package domain

import (
	"strings"
	"time"
)

// Status represents the status of a model.
type Status string

const (
	// StatusAdded represents a model that has been created.
	StatusAdded Status = "added"
	// StatusPulling represents a model that is being pulled.
	StatusPulling Status = "pulling"
	// StatusDeleted represents a model that has been deleted.
	StatusDeleted Status = "deleted"
)

// OllamaModel represents an Ollama model.
type OllamaModel struct {
	Model       string    `json:"model"`
	Description string    `json:"description"`
	AddedAt     time.Time `json:"addedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
	Status      Status    `json:"status"`
}

func NewOllamaModel(model string) *OllamaModel {
	return &OllamaModel{
		Model:     model,
		AddedAt:   time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Status:    StatusAdded,
	}
}

func (om *OllamaModel) Pull() {
	om.Status = StatusPulling
}

func (om *OllamaModel) Delete() {
	om.DeletedAt = time.Now().UTC()
	om.Status = StatusDeleted
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

func (om *OllamaModel) Scan(value any) error {
	*om = *NewOllamaModel(value.(string))
	return nil
}

func (om OllamaModel) Value() (any, error) {
	return om.String(), nil
}
