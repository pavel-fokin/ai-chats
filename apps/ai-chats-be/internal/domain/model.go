package domain

import (
	"fmt"
	"strings"
)

type Model struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func NewModel(model string) Model {
	if model == "" {
		return Model{}
	}

	name := model
	tag := "latest"
	if strings.Contains(model, ":") {
		parts := strings.Split(model, ":")
		name = parts[0]
		tag = parts[1]
	}

	return Model{
		Name: name,
		Tag:  tag,
	}
}

func (m Model) String() string {
	if m.Tag == "" {
		return m.Name
	}
	return fmt.Sprintf("%s:%s", m.Name, m.Tag)
}

func (m *Model) Scan(value interface{}) error {
	*m = NewModel(value.(string))
	return nil
}

func (m Model) Value() (interface{}, error) {
	return m.String(), nil
}
