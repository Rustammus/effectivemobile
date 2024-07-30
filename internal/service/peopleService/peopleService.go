package peopleService

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/internal/schemas"
	"EffectiveMobile/internal/schemas/externalApi"
	"EffectiveMobile/pkg/logging"
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
	Repo   repos.PeopleRepository
	Logger logging.Logger
}

func (r *PeopleService) Create(p schemas.RequestCreatePeople) (pgtype.UUID, error) {
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
		r.Logger.Warnf("dddddddddddddddddddddddddddddddddd")
		return pgtype.UUID{}, err
	}

	peopleCreate := dto.CreatePeople{
		PassportSerie:  passportSerie,
		PassportNumber: passportNumber,
		Surname:        people.Surname,
		Name:           people.Name,
		Patronymic:     people.Patronymic,
		Address:        people.Address,
	}

	uuid, err := r.Repo.Create(context.TODO(), peopleCreate)
	if err != nil {
		r.Logger.Error("error on create people. Values: %+v", peopleCreate)
		return pgtype.UUID{}, err
	}
	r.Logger.Debugf("created people: %x\n\t Values: %+v", uuid.Bytes, peopleCreate)

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
		r.Logger.Errorf("error on requesting people info. passport: %d %d. %s", passportSerie, passportNumber, err.Error())
		return people, err
	}
	defer response.Body.Close()

	data := make([]byte, response.ContentLength-1)
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

func (r *PeopleService) FindByUUID(uuid pgtype.UUID) (dto.ReadPeople, error) {
	people, err := r.Repo.FindByUUID(context.TODO(), uuid)
	if err != nil {
		return dto.ReadPeople{}, err
	}
	return people, nil
}

func (r *PeopleService) FindAll() ([]dto.ReadPeople, error) {
	peoples, err := r.Repo.FindAll(context.TODO())
	if err != nil {
		return nil, err
	}
	return peoples, nil
}

func (r *PeopleService) FindAllByOffset(pag crud.Pagination) ([]dto.ReadPeople, crud.Pagination, error) {
	peoples, err := r.Repo.FindAllByOffset(context.TODO(), pag)
	if err != nil {
		return []dto.ReadPeople{}, crud.Pagination{}, err
	}

	pag.Offset += pag.Limit

	return peoples, pag, nil
}

func (r *PeopleService) FindByFilterOffset(f schemas.RequestFilterPeople, pag crud.Pagination) ([]dto.ReadPeople, crud.Pagination, error) {
	if f.UUID.Valid {
		people, err := r.Repo.FindByUUID(context.TODO(), f.UUID)
		ppls := make([]dto.ReadPeople, 0)
		ppls = append(ppls, people)
		pag.Offset = 0
		pag.Limit = 1
		return ppls, pag, err
	}

	filter := dto.FilterPeople{
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

func (r *PeopleService) UpdateByUUID(uuid pgtype.UUID, p schemas.RequestUpdatePeople) (dto.ReadPeople, error) {
	updateDTO := dto.UpdatePeople{
		PassportSerie:  p.PassportSerie,
		PassportNumber: p.PassportNumber,
		Surname:        p.Surname,
		Name:           p.Name,
		Patronymic:     p.Patronymic,
		Address:        p.Address,
	}

	rPeople, err := r.Repo.Update(context.TODO(), uuid, updateDTO)
	if err != nil {
		return dto.ReadPeople{}, err
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
