package peopleService

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/crud"
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
	//TODO change schema to dto?
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

func (r *PeopleService) FindByUUID(uuid pgtype.UUID) (dto.ReadPeopleDTO, error) {
	people, err := r.Repo.FindByUUID(context.TODO(), uuid)
	if err != nil {
		return dto.ReadPeopleDTO{}, err
	}
	return people, nil
}

func (r *PeopleService) FindAll() ([]dto.ReadPeopleDTO, error) {
	peoples, err := r.Repo.FindAll(context.TODO())
	if err != nil {
		return nil, err
	}
	return peoples, nil
}

func (r *PeopleService) FindAllByOffset(pag crud.Pagination) ([]dto.ReadPeopleDTO, crud.Pagination, error) {
	peoples, err := r.Repo.FindAllByOffset(context.TODO(), pag)
	if err != nil {
		return []dto.ReadPeopleDTO{}, crud.Pagination{}, err
	}

	pag.Offset += pag.Limit

	return peoples, pag, nil
}

func (r *PeopleService) FindByFilterOffset(f peopleSchemas.RequestFilterPeople, pag crud.Pagination) ([]dto.ReadPeopleDTO, crud.Pagination, error) {
	if f.UUID.Valid {
		people, err := r.Repo.FindByUUID(context.TODO(), f.UUID)
		ppls := make([]dto.ReadPeopleDTO, 0)
		ppls = append(ppls, people)
		pag.Offset = 0
		pag.Limit = 1
		return ppls, pag, err
	}

	filter := dto.FilterPeopleDTO{
		UUID:           f.UUID,
		PassportSerie:  f.PassportSerie,
		PassportNumber: f.PassportNumber,
		Surname:        f.Surname,
		Name:           f.Name,
		Patronymic:     f.Patronymic,
		Address:        f.Address,
	}

	peoples, err := r.Repo.FindByFilterOffset(context.TODO(), filter, pag)
	if err != nil {
		return nil, crud.Pagination{}, err
	}

	pag.Offset += pag.Limit

	return peoples, pag, nil
}

func (r *PeopleService) UpdateByUUID(uuid pgtype.UUID, p peopleSchemas.RequestUpdatePeople) (dto.ReadPeopleDTO, error) {
	updateDTO := dto.UpdatePeopleDTO{
		PassportSerie:  p.PassportSerie,
		PassportNumber: p.PassportNumber,
		Surname:        p.Surname,
		Name:           p.Name,
		Patronymic:     p.Patronymic,
		Address:        p.Address,
	}

	rPeople, err := r.Repo.Update(context.TODO(), uuid, updateDTO)
	if err != nil {
		return dto.ReadPeopleDTO{}, err
	}
	return rPeople, nil
}

func (r *PeopleService) DeleteByUUID(uuid pgtype.UUID) (pgtype.UUID, error) {
	rUUID, err := r.Repo.Delete(context.TODO(), uuid)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return rUUID, nil
}
