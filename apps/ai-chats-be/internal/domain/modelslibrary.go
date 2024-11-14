package domain

import (
	"context"
	"errors"
)

var (
	ErrModelNotFound = errors.New("model not found")
)

type ModelsLibrary interface {
	FindAll(context.Context) ([]*ModelCard, error)
	FindDescription(context.Context, string) (string, error)
	FindByName(context.Context, string) (*ModelCard, error)
}
