package domain

import (
	"github.com/google/uuid"
)

type ModelID uuid.UUID

func NewModelID() ModelID {
	return ModelID(uuid.New())
}

type ModelTag struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  string    `json:"created_at"`
	UploadedAt string    `json:"uploaded_at"`
	DeletedAt  string    `json:"deleted_at"`
}

func NewModelTag(name string) ModelTag {
	return ModelTag{
		ID:   uuid.New(),
		Name: name,
	}
}

type Model struct {
	ID        ModelID    `json:"id"`
	Name      string     `json:"name"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	DeletedAt string     `json:"deleted_at"`
	Tags      []ModelTag `json:"tags"`
}

func NewModel(name string) Model {
	return Model{
		ID:   NewModelID(),
		Name: name,
	}
}

func (m *Model) AddTag(tag ModelTag) error {
	m.Tags = append(m.Tags, tag)
	return nil
}

func (m *Model) RemoveTag(tag ModelTag) error {
	for i, t := range m.Tags {
		if t.ID == tag.ID {
			m.Tags = append(m.Tags[:i], m.Tags[i+1:]...)
			return nil
		}
	}

	return ErrTagNotFound
}
