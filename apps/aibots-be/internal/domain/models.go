package domain

import "context"

type Models interface {
	Add(ctx context.Context, model Model) error
	Update(ctx context.Context, model Model) error
	Delete(ctx context.Context, id ModelID) error
}
