package peopleService

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/repos"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"time"
)

type PeopleService struct {
	Repo repos.PeopleRepository
}

func (r *PeopleService) Create(dto dto.CreatePeopleDTO) (pgtype.UUID, error) {
	uuid, err := r.Repo.Create(context.TODO(), dto)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func (r *PeopleService) requestPeopleInfo(passportSerie, passportNumber int) {
	conf := config.GetConfig()
	url := conf.Server.ExternalURL
	//TODO do do
	client := http.DefaultClient
	client.Timeout = time.Second * 5
	response, err := client.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

}
