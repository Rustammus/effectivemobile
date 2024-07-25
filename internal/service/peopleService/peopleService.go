package peopleService

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/internal/schemas/externalApi"
	"EffectiveMobile/internal/schemas/peopleSchemas"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PeopleService struct {
	Repo repos.PeopleRepository
}

func (r *PeopleService) Create(p peopleSchemas.RequestCreatePeople) (pgtype.UUID, error) {
	pass := strings.Split(p.PassportNumber, " ")
	if len(pass) < 2 {
		return pgtype.UUID{}, errors.New("error in passportNumber - not enough substrings. Example: '1234 567890'")
	}
	passportSerie, err := strconv.Atoi(pass[0])
	passportNumber, err := strconv.Atoi(pass[1])
	if err != nil {
		return pgtype.UUID{}, err
	}

	people, err := r.requestPeopleInfo(passportSerie, passportNumber)
	if err != nil {
		return pgtype.UUID{}, err
	}

	peopleCreate := dto.CreatePeopleDTO{
		PassportSerie:  passportSerie,
		PassportNumber: passportNumber,
		Surname:        people.Surname,
		Name:           people.Name,
		Patronymic:     people.Patronymic,
		Address:        people.Address,
	}

	uuid, err := r.Repo.Create(context.TODO(), peopleCreate)
	if err != nil {
		return pgtype.UUID{}, err
	}

	return uuid, nil
}

func (r *PeopleService) requestPeopleInfo(passportSerie, passportNumber int) (externalApi.ExResponsePeople, error) {
	conf := config.GetConfig()
	urlBase := conf.Server.ExternalURL
	people := externalApi.ExResponsePeople{}
	//TODO refactor this
	url := fmt.Sprintf("%s?passportSerie=%d&passportNumber=%d", urlBase, passportSerie, passportNumber)

	client := http.DefaultClient
	client.Timeout = time.Second * 5
	response, err := client.Get(url)
	if err != nil {
		return people, err
	}
	defer response.Body.Close()

	data := make([]byte, response.ContentLength)
	_, err = response.Body.Read(data)
	if err != nil {
		return people, err
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		return people, err
	}

	return people, nil
}
