package dto

import "github.com/jackc/pgx/v5/pgtype"

type CreatePeopleDTO struct {
	PassportSerie  int    `json:"passportSerie" example:"1234"`
	PassportNumber int    `json:"passportNumber" example:"567890"`
	Surname        string `json:"surname" example:"Иванов"`
	Name           string `json:"name" example:"Иван"`
	Patronymic     string `json:"patronymic" example:"Иванович"`
	Address        string `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	//IsDeleted      bool               `json:"is_deleted" example:"false"`
	//UpdatedAt      pgtype.Timestamptz `json:"updated_at" example:"2020-01-01T00:00:00Z"`
	//CreatedAt      pgtype.Timestamptz `json:"created_at" example:"2020-01-01T00:00:00Z"`
}

type ReadPeopleDTO struct {
	UUID           pgtype.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	PassportSerie  int         `json:"passportSerie" example:"1234"`
	PassportNumber int         `json:"passportNumber" example:"567890"`
	Surname        string      `json:"surname" example:"Иванов"`
	Name           string      `json:"name" example:"Иван"`
	Patronymic     string      `json:"patronymic" example:"Иванович"`
	Address        string      `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	//IsDeleted      bool               `json:"is_deleted" example:"false"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt" example:"2020-01-01T00:00:00Z"`
	CreatedAt pgtype.Timestamptz `json:"createdAt" example:"2020-01-01T00:00:00Z"`
}

type UpdatePeopleDTO struct {
	//UUID           pgtype.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	PassportSerie  int    `json:"passportSerie" example:"1234"`
	PassportNumber int    `json:"passportNumber" example:"567890"`
	Surname        string `json:"surname" example:"Иванов"`
	Name           string `json:"name" example:"Иван"`
	Patronymic     string `json:"patronymic" example:"Иванович"`
	Address        string `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	// IsDeleted      bool               `json:"is_deleted" example:"false"`
	//UpdatedAt pgtype.Timestamptz `json:"updated_at" example:"2020-01-01T00:00:00Z"`
	//CreatedAt pgtype.Timestamptz `json:"created_at" example:"2020-01-01T00:00:00Z"`
}