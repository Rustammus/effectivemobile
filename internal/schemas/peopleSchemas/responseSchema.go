package peopleSchemas

import (
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type RespPeople struct {
	UUID           pgtype.UUID        `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	PassportSerie  int                `json:"passportSerie" example:"1234"`
	PassportNumber int                `json:"passportNumber" example:"567890"`
	Surname        string             `json:"surname" example:"Иванов"`
	Name           string             `json:"name" example:"Иван"`
	Patronymic     string             `json:"patronymic" example:"Иванович"`
	Address        string             `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at" example:"2020-01-01T00:00:00Z"`
	CreatedAt      pgtype.Timestamptz `json:"created_at" example:"2020-01-01T00:00:00Z"`
}

type RespPeoplePaginated struct {
	Peoples        []dto.ReadPeopleDTO `json:"peoples"`
	NextPagination crud.Pagination     `json:"next_pagination"`
}
