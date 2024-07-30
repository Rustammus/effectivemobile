package repos

import (
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/pkg/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	People PeopleRepository
	Task   TaskRepository
}

func NewRepositories(logger logging.Logger, pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		People: crud.NewPeopleCRUD(logger, pool),
		Task:   crud.NewTaskCRUD(logger, pool),
	}
}
