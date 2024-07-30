package service

import (
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/internal/service/peopleService"
	"EffectiveMobile/internal/service/taskService"
	"EffectiveMobile/pkg/logging"
)

type Deps struct {
	Repos  *repos.Repositories
	Logger logging.Logger
}

type Services struct {
	People peopleService.PeopleService
	Task   taskService.TaskService
}

func NewServices(d Deps) *Services {
	return &Services{
		People: peopleService.PeopleService{
			Repo:   d.Repos.People,
			Logger: d.Logger,
		},
		Task: taskService.TaskService{
			Repo:   d.Repos.Task,
			Logger: d.Logger,
		},
	}
}
