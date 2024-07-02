package domain

// type ModelID uuid.UUID

// func NewModelID() ModelID {
// 	return ModelID(uuid.New())
// }

// func (m ModelID) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(uuid.UUID(m))
// }

type Model struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func NewModel(name, tag string) Model {
	return Model{
		Name: name,
		Tag:  tag,
	}
}
