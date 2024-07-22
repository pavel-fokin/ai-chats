package domain

import "context"

type Models interface {
	AllModelCards(context.Context) ([]ModelCard, error)
	FindModelCard(context.Context, string) (ModelCard, error)
}
