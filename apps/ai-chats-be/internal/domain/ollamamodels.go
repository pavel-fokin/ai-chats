package domain

import "context"

type OllamaModels interface {
	Add(ctx context.Context, model OllamaModel) error
	AllAdded(ctx context.Context) ([]OllamaModel, error)
	Delete(ctx context.Context, model OllamaModel) error
	Exists(ctx context.Context, model string) (bool, error)
	Find(ctx context.Context, model string) (OllamaModel, error)
	Save(ctx context.Context, model OllamaModel) error
}
