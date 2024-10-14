package domain

import "context"

type ModelsLibrary interface {
	FindDescription(context.Context, string) (string, error)
}
