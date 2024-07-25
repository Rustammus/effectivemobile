package crud

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/pkg/client/postgres"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type PeopleCRUD struct {
	client postgres.Client
}

func (p PeopleCRUD) Create(ctx context.Context, dto dto.CreatePeopleDTO) (pgtype.UUID, error) {
	q := `INSERT INTO peoples 
    (passport_serie, passport_number, surname, name, patronymic, address)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING uuid`
	uuid := pgtype.UUID{}
	err := p.client.QueryRow(ctx, q,
		dto.PassportSerie, dto.PassportNumber, dto.Surname, dto.Name, dto.Patronymic, dto.Address).
		Scan(&uuid)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func (p PeopleCRUD) FindAll(ctx context.Context) ([]dto.ReadPeopleDTO, error) {
	//TODO implement me
	q := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at FROM peoples`
	rows, err := p.client.Query(ctx, q)
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

func (p PeopleCRUD) FindByUUID(ctx context.Context, uuid pgtype.UUID) (dto.ReadPeopleDTO, error) {
	q := `SELECT uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at FROM peoples WHERE uuid=$1`
	people := dto.ReadPeopleDTO{}
	err := p.client.QueryRow(ctx, q, uuid).Scan(&people.UUID, &people.PassportSerie, &people.PassportNumber,
		&people.Surname, &people.Name, &people.Patronymic, &people.Address, &people.UpdatedAt, &people.CreatedAt)
	if err != nil {
		return dto.ReadPeopleDTO{}, err
	}
	return dto.ReadPeopleDTO{}, nil
}

func (p PeopleCRUD) FindByFilter() {
	//TODO
}

func (p PeopleCRUD) Update(ctx context.Context, people dto.UpdatePeopleDTO, uuid pgtype.UUID) (dto.ReadPeopleDTO, error) {
	q := `UPDATE peoples 
			SET (passport_serie, passport_number, surname, name, patronymic, address) = ($2, $3, $4, $5, $6, $7) 
			WHERE uuid=$1 
			RETURNING uuid, passport_serie, passport_number, surname, name, patronymic, address, updated_at, created_at`

	rPeople := dto.ReadPeopleDTO{}
	err := p.client.QueryRow(ctx, q, uuid, people.PassportSerie, people.PassportNumber,
		people.Surname, people.Name, people.Patronymic, people.Address).
		Scan(&rPeople.UUID, &rPeople.PassportSerie, &rPeople.PassportNumber, &rPeople.Surname,
			&rPeople.Name, &rPeople.Patronymic, &rPeople.Address, &rPeople.UpdatedAt, &rPeople.CreatedAt)
	if err != nil {
		return dto.ReadPeopleDTO{}, err
	}
	return rPeople, nil
}

func (p PeopleCRUD) Delete(ctx context.Context, uuid pgtype.UUID) (pgtype.UUID, error) {
	q := `DELETE FROM peoples WHERE uuid=$1	RETURNING uuid`
	delUUID := pgtype.UUID{}
	err := p.client.QueryRow(ctx, q, uuid).Scan(&delUUID)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return delUUID, nil
}

func NewPeopleCRUD(client postgres.Client) *PeopleCRUD {
	return &PeopleCRUD{client: client}
}
