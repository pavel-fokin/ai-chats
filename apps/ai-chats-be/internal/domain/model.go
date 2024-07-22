package domain

type Vendor string

type Model struct {
	Model  string `json:"model"`
	Vendor Vendor `json:"vendor"`
}

func NewModel(model string) Model {
	return Model{
		Model:  model,
		Vendor: "ollama",
	}
}

func (m Model) AsOllamaModel() OllamaModel {
	return OllamaModel{
		Model: m.Model,
	}
}

func (m Model) String() string {
	return m.Model
}

func (m *Model) Scan(value any) error {
	*m = NewModel(value.(string))
	return nil
}

func (m Model) Value() (any, error) {
	return m.String(), nil
}
