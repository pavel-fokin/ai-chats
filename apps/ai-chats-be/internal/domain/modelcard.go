package domain

type ModelCard struct {
	Model       string   `json:"model"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func NewModelCard(model, description string, tags []string) *ModelCard {
	return &ModelCard{Model: model, Description: description, Tags: tags}
}
