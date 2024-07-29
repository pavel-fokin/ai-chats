package domain

import "context"

type Models interface {
	FindDescription(context.Context, string) (string, error)
}
