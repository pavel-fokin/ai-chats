package sqlite

import (
	"context"
	"database/sql"
	"ai-chats/internal/domain"
)

type Models struct {
	db DB
}

func NewModels(db *sql.DB) *Models {
	return &Models{DB{db: db}}
}

func (m *Models) AddModelCard(ctx context.Context, modelDescription domain.ModelCard) error {
	_, err := m.db.DBTX(ctx).ExecContext(
		ctx,
		`INSERT INTO model_card
		(model, description)
		VALUES (?, ?)`,
		modelDescription.Model,
		modelDescription.Description,
	)
	return err
}

func (m *Models) AllModelCards(ctx context.Context) ([]domain.ModelCard, error) {
	rows, err := m.db.DBTX(ctx).QueryContext(
		ctx,
		`SELECT model, description
		FROM model_card`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	descriptions := []domain.ModelCard{}
	for rows.Next() {
		var model string
		var description string
		if err := rows.Scan(&model, &description); err != nil {
			return nil, err
		}

		descriptions = append(descriptions, domain.ModelCard{
			Model:       model,
			Description: description,
		})
	}

	return descriptions, nil
}
