package peopleSchemas

import "github.com/jackc/pgx/v5/pgtype"

type RequestCreatePeople struct {
	PassportNumber string `json:"passportNumber"`
}

type RequestUpdatePeople struct {
	PassportSerie  int    `json:"passportSerie" example:"1234"`
	PassportNumber int    `json:"passportNumber" example:"567890"`
	Surname        string `json:"surname" example:"Иванов"`
	Name           string `json:"name" example:"Иван"`
	Patronymic     string `json:"patronymic" example:"Иванович"`
	Address        string `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}

type RequestFilterPeople struct {
	UUID           pgtype.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	PassportSerie  int         `json:"passportSerie" example:"1234"`
	PassportNumber int         `json:"passportNumber" example:"567890"`
	Surname        string      `json:"surname" example:"Иванов"`
	Name           string      `json:"name" example:"Иван"`
	Patronymic     string      `json:"patronymic" example:"Иванович"`
	Address        string      `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}
