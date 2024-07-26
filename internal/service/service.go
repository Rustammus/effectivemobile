package service

import (
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/internal/service/peopleService"
	"EffectiveMobile/internal/service/taskService"
)

type Deps struct {
	Repos *repos.Repositories
}

type Services struct {
	People peopleService.PeopleService
	Task   taskService.TaskService
}

func NewServices(d Deps) *Services {
	return &Services{
		People: peopleService.PeopleService{
			Repo: d.Repos.People,
		},
		Task: taskService.TaskService{
			Repo: d.Repos.Task,
		},
	}
}
