package domain

import "time"

// Status represents the status of a model.
type Status string

const (
	// StatusAdded represents a model that has been created.
	StatusAdded Status = "added"
	// StatusPulling represents a model that is being pulled.
	StatusPulling Status = "pulling"
	// StatusAvaialable represents an available model.
	StatusAvailable Status = "available"
	// StatusDeleted represents a model that has been deleted.
	StatusDeleted Status = "deleted"
)

// OllamaModel represents an Ollama model.
type OllamaModel struct {
	Model       string    `json:"model"`
	Description string    `json:"description"`
	AddedAt     time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
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

func (om *OllamaModel) Update(model string) {
	om.UpdatedAt = time.Now().UTC()
}

func (om *OllamaModel) Available() {
	om.Status = StatusAvailable
}

func (om *OllamaModel) Pulling() {
	om.Status = StatusPulling
}

func (om *OllamaModel) Deleted() {
	om.DeletedAt = time.Now().UTC()
	om.Status = StatusDeleted
}

func (om OllamaModel) String() string {
	return om.Model
}

func (om *OllamaModel) Scan(value any) error {
	*om = *NewOllamaModel(value.(string))
	return nil
}

func (om OllamaModel) Value() (any, error) {
	return om.String(), nil
}
