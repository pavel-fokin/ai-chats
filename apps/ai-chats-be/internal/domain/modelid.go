package domain

type Vendor string

type ModelID struct {
	Model  string `json:"model"`
	Vendor Vendor `json:"vendor"`
}

func NewModelID(model string) ModelID {
	return ModelID{
		Model:  model,
		Vendor: "ollama",
	}
}

func (m ModelID) AsOllamaModel() OllamaModel {
	return OllamaModel{
		Model: m.Model,
	}
}

func (m ModelID) String() string {
	return m.Model
}

func (m *ModelID) Scan(value any) error {
	*m = NewModelID(value.(string))
	return nil
}

func (m ModelID) Value() (any, error) {
	return m.String(), nil
}
