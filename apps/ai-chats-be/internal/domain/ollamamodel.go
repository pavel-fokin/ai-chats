package domain

import (
	"fmt"
	"strings"
)

type OllamaModel struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func NewModel(model string) OllamaModel {
	if model == "" {
		return OllamaModel{}
	}

	name := model
	tag := "latest"
	if strings.Contains(model, ":") {
		parts := strings.Split(model, ":")
		name = parts[0]
		tag = parts[1]
	}

	return OllamaModel{
		Name: name,
		Tag:  tag,
	}
}

func (m OllamaModel) String() string {
	if m.Tag == "" {
		return m.Name
	}
	return fmt.Sprintf("%s:%s", m.Name, m.Tag)
}

func (m *OllamaModel) Scan(value interface{}) error {
	*m = NewModel(value.(string))
	return nil
}

func (m OllamaModel) Value() (interface{}, error) {
	return m.String(), nil
}
