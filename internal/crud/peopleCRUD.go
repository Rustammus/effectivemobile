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
	//TODO implement me
	q := `INSERT INTO peoples 
    (passportSerie, passportNumber, surname, name, patronymic, address)
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
	q := `SELECT uuid, passportSerie, passportNumber, surname, name, patronymic, address, updatedAt, createdAt FROM people`
}

func (p PeopleCRUD) FindByUUID(ctx context.Context, uuid pgtype.UUID) (*dto.ReadPeopleDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (p PeopleCRUD) Update(ctx context.Context, user dto.UpdatePeopleDTO, uuid pgtype.UUID) (dto.ReadPeopleDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (p PeopleCRUD) Delete(ctx context.Context, uuid pgtype.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p PeopleCRUD) SoftDelete(ctx context.Context, uuid pgtype.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewPeopleCRUD(client postgres.Client) *PeopleCRUD {
	return &PeopleCRUD{client: client}
}
