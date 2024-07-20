package domain

import "context"

type Models interface {
	AllModelCards(context.Context) ([]ModelCard, error)
}
