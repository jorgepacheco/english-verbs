package domain

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Verb, error)
}
