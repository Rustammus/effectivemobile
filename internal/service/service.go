package service

import (
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/internal/service/peopleService"
)

type Deps struct {
	Repos *repos.Repositories
}

type Services struct {
	People peopleService.PeopleService
}

func NewServices(d Deps) *Services {
	return &Services{
		People: peopleService.PeopleService{
			Repo: d.Repos.People,
		},
	}
}
