package repos

import (
	"EffectiveMobile/internal/crud"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	User PeopleRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: crud.NewPeopleCRUD(pool),
	}
}
