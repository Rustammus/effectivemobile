package crud

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/pkg/client/postgres"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

type PeopleCRUD struct {
	client postgres.Client
}

func (c PeopleCRUD) Create(ctx context.Context, dto dto.CreatePeopleDTO) (pgtype.UUID, error) {
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

func (c PeopleCRUD) FindAll(ctx context.Context) ([]dto.ReadPeopleDTO, error) {
	q := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at 
		  FROM public.peoples`
	rows, err := c.client.Query(ctx, q)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	peoples := make([]dto.ReadPeopleDTO, 0)
	for rows.Next() {
		var people dto.ReadPeopleDTO
		err = rows.Scan(&people.UUID, &people.PassportSerie, &people.PassportNumber, &people.Surname,
			&people.Name, &people.Patronymic, &people.Address, &people.UpdatedAt, &people.CreatedAt)
		if err != nil {
			return nil, err
		}
		peoples = append(peoples, people)
	}
	return peoples, nil
}

func (c PeopleCRUD) FindByUUID(ctx context.Context, uuid pgtype.UUID) (dto.ReadPeopleDTO, error) {
	q := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at 
		  FROM public.peoples 
		  WHERE uuid=$1`
	people := dto.ReadPeopleDTO{}
	err := c.client.QueryRow(ctx, q, uuid).Scan(&people.UUID, &people.PassportSerie, &people.PassportNumber,
		&people.Surname, &people.Name, &people.Patronymic, &people.Address, &people.UpdatedAt, &people.CreatedAt)
	if err != nil {
		return dto.ReadPeopleDTO{}, err
	}
	return dto.ReadPeopleDTO{}, nil
}

func (c PeopleCRUD) FindByFilter(ctx context.Context, filter dto.FilterPeopleDTO) ([]dto.ReadPeopleDTO, error) {
	query := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at FROM public.peoples WHERE`

	if filter.UUID.Valid {
		people, err := c.FindByUUID(ctx, filter.UUID)
		ppls := make([]dto.ReadPeopleDTO, 0)
		ppls = append(ppls, people)
		return ppls, err
	}

	//TODO refactor me!

	values := make([]any, 0)
	counter := 0
	if filter.PassportSerie >= 0 {
		counter++
		values = append(values, filter.PassportSerie)
		query = fmt.Sprintf("%s passport_serie = $%d AND", query, counter)
	}
	if filter.PassportNumber >= 0 {
		counter++
		values = append(values, filter.PassportNumber)
		query = fmt.Sprintf("%s passport_number = $%d AND", query, counter)
	}
	if len(filter.Surname) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%", filter.Surname))
		query = fmt.Sprintf("%s surname LIKE $%d AND", query, counter)
	}
	if len(filter.Name) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%", filter.Name))
		query = fmt.Sprintf("%s name LIKE $%d AND", query, counter)
	}
	if len(filter.Patronymic) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%", filter.Patronymic))
		query = fmt.Sprintf("%s patronymic LIKE $%d AND", query, counter)
	}
	if len(filter.Address) > 0 {
		counter++
		values = append(values, fmt.Sprintf("%%%s%", filter.Address))
		query = fmt.Sprintf("%s address LIKE $%d AND", query, counter)
	}

	if counter <= 0 {
		return nil, nil
	}
	query = query[:len(query)-4]

	rows, err := c.client.Query(ctx, query, values...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	peoples := make([]dto.ReadPeopleDTO, 0)
	for rows.Next() {
		var people dto.ReadPeopleDTO
		err = rows.Scan(&people.UUID, &people.PassportSerie, &people.PassportNumber, &people.Surname,
			&people.Name, &people.Patronymic, &people.Address, &people.UpdatedAt, &people.CreatedAt)
		if err != nil {
			return nil, err
		}
		peoples = append(peoples, people)
	}
	return peoples, nil
}

func (c PeopleCRUD) Update(ctx context.Context, people dto.UpdatePeopleDTO, uuid pgtype.UUID) (dto.ReadPeopleDTO, error) {
	q := `UPDATE public.peoples 
		  SET (passport_serie, passport_number, surname, name, patronymic, address, updated_at) = ($2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP(0)) 
		  WHERE uuid=$1 
		  RETURNING uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at`

	rPeople := dto.ReadPeopleDTO{}
	err := c.client.QueryRow(ctx, q, uuid, people.PassportSerie, people.PassportNumber,
		people.Surname, people.Name, people.Patronymic, people.Address).
		Scan(&rPeople.UUID, &rPeople.PassportSerie, &rPeople.PassportNumber, &rPeople.Surname,
			&rPeople.Name, &rPeople.Patronymic, &rPeople.Address, &rPeople.UpdatedAt, &rPeople.CreatedAt)
	if err != nil {
		return dto.ReadPeopleDTO{}, err
	}

	return rPeople, nil
}

func (c PeopleCRUD) Delete(ctx context.Context, uuid pgtype.UUID) (pgtype.UUID, error) {
	q := `DELETE FROM public.peoples WHERE uuid=$1	RETURNING uuid`
	delUUID := pgtype.UUID{}
	err := c.client.QueryRow(ctx, q, uuid).Scan(&delUUID)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return delUUID, nil
}

func NewPeopleCRUD(client postgres.Client) *PeopleCRUD {
	return &PeopleCRUD{client: client}
}
