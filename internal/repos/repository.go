package repos

import (
	"EffectiveMobile/internal/crud"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	People PeopleRepository
	Task   TaskRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		People: crud.NewPeopleCRUD(pool),
		Task:   crud.NewTaskCRUD(pool),
	}
}
