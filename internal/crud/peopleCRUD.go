package crud

import (
	"EffectiveMobile/pkg/client/postgres"
	"context"
)

type PeopleCRUD struct {
	client postgres.Client
}

func (p PeopleCRUD) Create(ctx context.Context) {

}
