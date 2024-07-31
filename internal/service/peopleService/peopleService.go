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
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"io"
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
	pass := strings.Split(p.PassportNumber, " ")
	if len(pass) != 2 {
		return pgtype.UUID{}, errors.New("error in passportNumber - not valid or not present. Example: '1234 567890'")
	}
	passportSerie, err := strconv.Atoi(pass[0])
	passportNumber, err := strconv.Atoi(pass[1])
	if err != nil {
		return pgtype.UUID{}, err
	}

	peopleInfo, err := r.requestPeopleInfo(passportSerie, passportNumber)
	if err != nil {
		r.Logger.Infof("Error on request people info: %s", err)
		return pgtype.UUID{}, err
	}

	if !peopleInfo.Valid() {
		return pgtype.UUID{}, errors.New("requested people info is not valid or empty")
	}

	peopleCreate := dto.CreatePeople{
		PassportSerie:  passportSerie,
		PassportNumber: passportNumber,
		Surname:        peopleInfo.Surname,
		Name:           peopleInfo.Name,
		Patronymic:     peopleInfo.Patronymic,
		Address:        peopleInfo.Address,
	}

	uuid, err := r.Repo.Create(context.TODO(), peopleCreate)
	if err != nil {
		r.Logger.Infof("error on create people. Values: %+v", peopleCreate)
		return pgtype.UUID{}, err
	}
	r.Logger.Debugf("created people: %x\n\t Values: %+v", uuid.Bytes, peopleCreate)

	return uuid, nil
}

func (r *PeopleService) requestPeopleInfo(passportSerie, passportNumber int) (externalApi.ExResponsePeople, error) {
	conf := config.GetConfig()
	urlBase := conf.Server.ExternalURL
	people := externalApi.ExResponsePeople{}

	url := fmt.Sprintf("%s/info?passportSerie=%d&passportNumber=%d", urlBase, passportSerie, passportNumber)

	client := http.DefaultClient
	client.Timeout = time.Second * 5
	response, err := client.Get(url)
	if response != nil {
		defer response.Body.Close()
	} else {
		r.Logger.Errorf("error on get request people info. response is nil ")
		return people, errors.New("error on get request people info. response is nil")
	}

	if err != nil {
		r.Logger.Errorf("error on get request people info; url: %s. %s", url, err.Error())
		return externalApi.ExResponsePeople{}, err
	}

	data := make([]byte, response.ContentLength)
	_, err = response.Body.Read(data)
	if err != nil && err != io.EOF {
		r.Logger.Errorf("error on read response body: %s", err.Error())
		return externalApi.ExResponsePeople{}, err
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		return externalApi.ExResponsePeople{}, err
	}

	r.Logger.Debugf("got people info: %+v", people)
	return people, nil
}

func (r *PeopleService) FindByUUID(uuid pgtype.UUID) (dto.ReadPeople, error) {
	people, err := r.Repo.FindByUUID(context.TODO(), uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.Logger.Debugf("no rows found by uuid: %x", uuid.Bytes)
			return dto.ReadPeople{}, err
		}
		r.Logger.Infof("error on find people: %s", err.Error())
		return dto.ReadPeople{}, err
	}
	r.Logger.Debugf("people find by uuid: %+v", people)
	return people, nil
}

func (r *PeopleService) FindAllByOffset(pag crud.Pagination) ([]dto.ReadPeople, crud.Pagination, error) {
	peoples, err := r.Repo.FindAllByOffset(context.TODO(), pag)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.Logger.Debugf("no rows found on find all people")
			return nil, crud.Pagination{}, err
		}
		r.Logger.Infof("error on find peoples: %s", err.Error())
		return nil, crud.Pagination{}, err
	}

	pag.Offset += pag.Limit

	return peoples, pag, nil
}

func (r *PeopleService) FindByFilterOffset(f schemas.RequestFilterPeople, pag crud.Pagination) ([]dto.ReadPeople, crud.Pagination, error) {
	if f.UUID.Valid {
		r.Logger.Debugf("uuid contains in people filter, call find by uuid")
		people, err := r.Repo.FindByUUID(context.TODO(), f.UUID)
		ppls := make([]dto.ReadPeople, 0)
		ppls = append(ppls, people)
		pag.Offset = 0
		pag.Limit = 1
		return ppls, pag, err
	}

	filter := dto.FilterPeople{
		PassportSerie:  f.PassportSerie,
		PassportNumber: f.PassportNumber,
		Surname:        f.Surname,
		Name:           f.Name,
		Patronymic:     f.Patronymic,
		Address:        f.Address,
	}

	peoples, err := r.Repo.FindByFilterOffset(context.TODO(), filter, pag)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.Logger.Debugf("no rows found by filter: %+v", filter)
			return nil, crud.Pagination{}, err
		}
		r.Logger.Infof("error on find peoples by filter: %s", err.Error())
		return nil, crud.Pagination{}, err
	}

	pag.Offset += pag.Limit
	r.Logger.Debugf("peoples found by filter: %+v", peoples)
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
		if errors.Is(err, pgx.ErrNoRows) {
			r.Logger.Debugf("no rows updated by uuid: %x", uuid.Bytes)
			return dto.ReadPeople{}, err
		}
		r.Logger.Infof("error on update people: %s", err.Error())
		return dto.ReadPeople{}, err
	}
	r.Logger.Debugf("people updated: %+v", rPeople)
	return rPeople, nil
}

func (r *PeopleService) DeleteByUUID(uuid pgtype.UUID) (pgtype.UUID, error) {
	rUUID, err := r.Repo.Delete(context.TODO(), uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.Logger.Debugf("no rows deleted by uuid: %x", uuid.Bytes)
			return pgtype.UUID{}, err
		}
		r.Logger.Infof("error on delete people: %s", err.Error())
		return pgtype.UUID{}, err
	}
	r.Logger.Debugf("people deleted: %x", rUUID.Bytes)
	return rUUID, nil
}
