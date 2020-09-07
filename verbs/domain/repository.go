package domain

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Verb, error)
	FindById(id int) (Verb, error)
	GetSize() int
}
