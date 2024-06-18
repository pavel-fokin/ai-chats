package domain

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ModelID uuid.UUID

func NewModelID() ModelID {
	return ModelID(uuid.New())
}

func (m ModelID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(m))
}

type Model struct {
	ID   ModelID `json:"id"`
	Name string  `json:"name"`
	Tag  string  `json:"tag"`
}

func NewModel(name, tag string) Model {
	return Model{
		ID:   NewModelID(),
		Name: name,
		Tag:  tag,
	}
}
