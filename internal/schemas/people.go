package schemas

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type RequestCreatePeople struct {
	PassportNumber string `json:"passportNumber"`
}

type RequestUpdatePeople struct {
	PassportSerie  *int    `json:"passportSerie" example:"1234"`
	PassportNumber *int    `json:"passportNumber" example:"567890"`
	Surname        *string `json:"surname" example:"Иванов"`
	Name           *string `json:"name" example:"Иван"`
	Patronymic     *string `json:"patronymic" example:"Иванович"`
	Address        *string `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}

type RequestFilterPeople struct {
	UUID           pgtype.UUID `json:"uuid" form:"-" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	PassportSerie  int         `json:"passportSerie" form:"passportSerie" example:"1234"`
	PassportNumber int         `json:"passportNumber" form:"passportNumber" example:"567890"`
	Surname        string      `json:"surname" form:"surname" example:"Иванов"`
	Name           string      `json:"name" form:"name" example:"Иван"`
	Patronymic     string      `json:"patronymic" form:"patronymic" example:"Иванович"`
	Address        string      `json:"address" form:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}

type ResponsePeople struct {
	UUID           pgtype.UUID        `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	PassportSerie  int                `json:"passportSerie" example:"1234"`
	PassportNumber int                `json:"passportNumber" example:"567890"`
	Surname        string             `json:"surname" example:"Иванов"`
	Name           string             `json:"name" example:"Иван"`
	Patronymic     string             `json:"patronymic" example:"Иванович"`
	Address        string             `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at" example:"2020-01-01T00:00:00Z" swaggertype:"string"`
	CreatedAt      pgtype.Timestamptz `json:"created_at" example:"2020-01-01T00:00:00Z" swaggertype:"string"`
}
