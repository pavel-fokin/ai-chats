package domain

type ModelCard struct {
	ModelName   string   `json:"modelName"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func NewModelCard(modelName, description string, tags []string) *ModelCard {
	return &ModelCard{ModelName: modelName, Description: description, Tags: tags}
}
