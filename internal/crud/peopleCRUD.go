package crud

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/pkg/client/postgres"
	"EffectiveMobile/pkg/logging"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"strings"
)

type PeopleCRUD struct {
	client postgres.Client
	logger logging.Logger
}

type Pagination struct {
	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit" form:"limit"`
}

func (c *PeopleCRUD) Create(ctx context.Context, dto dto.CreatePeople) (pgtype.UUID, error) {
	q := `INSERT INTO public.peoples 
    	  (passport_serie, passport_number, surname, name, patronymic, address)
    	  VALUES ($1, $2, $3, $4, $5, $6)
    	  RETURNING uuid`
	uuid := pgtype.UUID{}
	err := c.client.QueryRow(ctx, q,
		dto.PassportSerie, dto.PassportNumber, dto.Surname, dto.Name, dto.Patronymic, dto.Address).
		Scan(&uuid)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func (c *PeopleCRUD) FindAllByOffset(ctx context.Context, pag Pagination) ([]dto.ReadPeople, error) {
	q := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at 
		  FROM public.peoples 
		  OFFSET $1 LIMIT $2`

	peoples := make([]dto.ReadPeople, 0)
	rows, err := c.client.Query(ctx, q, pag.Offset, pag.Limit)
	defer rows.Close()
	if err != nil {
		return peoples, err
	}
	for rows.Next() {
		people := dto.ReadPeople{}
		err = rows.Scan(&people.UUID, &people.PassportSerie, &people.PassportNumber, &people.Surname,
			&people.Name, &people.Patronymic, &people.Address, &people.UpdatedAt, &people.CreatedAt)
		if err != nil {
			return nil, err
		}
		peoples = append(peoples, people)
	}
	if len(peoples) == 0 {
		return nil, pgx.ErrNoRows
	}

	return peoples, nil
}

func (c *PeopleCRUD) FindByUUID(ctx context.Context, uuid pgtype.UUID) (dto.ReadPeople, error) {
	q := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at 
		  FROM public.peoples 
		  WHERE uuid=$1`
	people := dto.ReadPeople{}
	err := c.client.QueryRow(ctx, q, uuid).Scan(&people.UUID, &people.PassportSerie, &people.PassportNumber,
		&people.Surname, &people.Name, &people.Patronymic, &people.Address, &people.UpdatedAt, &people.CreatedAt)
	if err != nil {
		return dto.ReadPeople{}, err
	}
	return people, nil
}

func (c *PeopleCRUD) FindByFilterOffset(ctx context.Context, filter dto.FilterPeople, pag Pagination) ([]dto.ReadPeople, error) {
	query := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at 
			  FROM public.peoples WHERE`

	values, query := c.buildWhereCondition(query, filter)
	counter := len(values)
	if counter == 0 {
		c.logger.Error("got empty people filter")
		return nil, errors.New("empty people filter")
	}
	query = fmt.Sprintf("%s OFFSET $%d LIMIT $%d", query, counter+1, counter+2)
	values = append(values, pag.Offset, pag.Limit)

	rows, err := c.client.Query(ctx, query, values...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	peoples := make([]dto.ReadPeople, 0)

	for rows.Next() {
		people := dto.ReadPeople{}
		err = rows.Scan(&people.UUID, &people.PassportSerie, &people.PassportNumber, &people.Surname,
			&people.Name, &people.Patronymic, &people.Address, &people.UpdatedAt, &people.CreatedAt)
		if err != nil {
			return nil, err
		}
		peoples = append(peoples, people)
	}
	if len(peoples) == 0 {
		return nil, pgx.ErrNoRows
	}

	return peoples, nil
}

func (c *PeopleCRUD) buildWhereCondition(query string, filter dto.FilterPeople) ([]any, string) {

	values := make([]any, 0)
	counter := 0
	if filter.PassportSerie > 0 {
		counter++
		values = append(values, filter.PassportSerie)
		query = fmt.Sprintf("%s passport_serie = $%d AND", query, counter)
	}
	if filter.PassportNumber > 0 {
		counter++
		values = append(values, filter.PassportNumber)
		query = fmt.Sprintf("%s passport_number = $%d AND", query, counter)
	}
	if len(filter.Surname) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%%", filter.Surname))
		query = fmt.Sprintf("%s surname LIKE $%d AND", query, counter)
	}
	if len(filter.Name) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
		query = fmt.Sprintf("%s name LIKE $%d AND", query, counter)
	}
	if len(filter.Patronymic) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%%", filter.Patronymic))
		query = fmt.Sprintf("%s patronymic LIKE $%d AND", query, counter)
	}
	if len(filter.Address) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%%", filter.Address))
		query = fmt.Sprintf("%s address LIKE $%d AND", query, counter)
	}

	query = query[:len(query)-4]

	return values, query
}

func (c *PeopleCRUD) Update(ctx context.Context, uuid pgtype.UUID, people dto.UpdatePeople) (dto.ReadPeople, error) {
	baseQuery := `UPDATE public.peoples 
		  SET (%s) = (%s) 
		  WHERE uuid=$1 
		  RETURNING uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at`

	q, values := c.buildUpdateValues(baseQuery, uuid, people)

	rPeople := dto.ReadPeople{}
	err := c.client.QueryRow(ctx, q, values...).
		Scan(&rPeople.UUID, &rPeople.PassportSerie, &rPeople.PassportNumber, &rPeople.Surname,
			&rPeople.Name, &rPeople.Patronymic, &rPeople.Address, &rPeople.UpdatedAt, &rPeople.CreatedAt)
	if err != nil {
		return dto.ReadPeople{}, err
	}

	return rPeople, nil
}

func (c *PeopleCRUD) buildUpdateValues(baseQuery string, uuid pgtype.UUID, p dto.UpdatePeople) (string, []any) {
	var setL []string
	var setR []string
	var values []any
	counter := 1
	values = append(values, uuid)
	if p.PassportSerie != nil {
		counter++
		setL = append(setL, "passport_serie")
		setR = append(setR, fmt.Sprintf("$%d", counter))
		values = append(values, p.PassportSerie)
	}
	if p.PassportNumber != nil {
		counter++
		setL = append(setL, "passport_number")
		setR = append(setR, fmt.Sprintf("$%d", counter))
		values = append(values, p.PassportNumber)
	}
	if p.Surname != nil {
		counter++
		setL = append(setL, "surname")
		setR = append(setR, fmt.Sprintf("$%d", counter))
		values = append(values, p.Surname)
	}
	if p.Name != nil {
		counter++
		setL = append(setL, "name")
		setR = append(setR, fmt.Sprintf("$%d", counter))
		values = append(values, p.Name)
	}
	if p.Patronymic != nil {
		counter++
		setL = append(setL, "patronymic")
		setR = append(setR, fmt.Sprintf("$%d", counter))
		values = append(values, p.Patronymic)
	}
	if p.Address != nil {
		counter++
		setL = append(setL, "address")
		setR = append(setR, fmt.Sprintf("$%d", counter))
		values = append(values, p.Address)
	}

	counter++
	setL = append(setL, "updated_at")
	setR = append(setR, "CURRENT_TIMESTAMP(0)")
	q := fmt.Sprintf(baseQuery, strings.Join(setL, ","), strings.Join(setR, ","))

	return q, values
}

func (c *PeopleCRUD) Delete(ctx context.Context, uuid pgtype.UUID) (pgtype.UUID, error) {
	q := `DELETE FROM public.peoples WHERE uuid=$1	RETURNING uuid`
	delUUID := pgtype.UUID{}
	err := c.client.QueryRow(ctx, q, uuid).Scan(&delUUID)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return delUUID, nil
}

func NewPeopleCRUD(logger logging.Logger, client postgres.Client) *PeopleCRUD {
	return &PeopleCRUD{client: client, logger: logger}
}
