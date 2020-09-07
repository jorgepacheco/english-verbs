package app

import (
	"context"
	"english-verbs/verbs/domain"
)

// Interfaz

type VerbService struct {
	Repo domain.Repository
}

func (srv *VerbService) GetAll(ctx context.Context) ([]domain.Verb, error) {

	return srv.Repo.GetAll(ctx)
}
