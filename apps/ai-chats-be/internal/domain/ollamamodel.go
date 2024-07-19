package domain

type OllamaModel struct {
	Model string `json:"model"`
}

func NewOllamaModel(model string) OllamaModel {
	return OllamaModel{
		Model: model,
	}
}

func (m OllamaModel) String() string {
	return m.Model
}

func (m *OllamaModel) Scan(value any) error {
	*m = NewOllamaModel(value.(string))
	return nil
}

func (m OllamaModel) Value() (any, error) {
	return m.String(), nil
}
