package domain

type OllamaModel struct {
	Model       string `json:"model"`
	Description string `json:"description"`
}

func NewOllamaModel(model string) OllamaModel {
	return OllamaModel{
		Model: model,
	}
}

func (om OllamaModel) String() string {
	return om.Model
}

func (om *OllamaModel) Scan(value any) error {
	*om = NewOllamaModel(value.(string))
	return nil
}

func (om OllamaModel) Value() (any, error) {
	return om.String(), nil
}
