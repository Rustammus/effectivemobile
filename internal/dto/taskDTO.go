package dto

import "github.com/jackc/pgx/v5/pgtype"

type CreateTask struct {
	//UUID       pgtype.UUID        `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	PeopleUUID pgtype.UUID `json:"people_uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name       string      `json:"name" example:"do nothing"`
	//StartTime  pgtype.Timestamptz `json:"start_time" example:"2020-01-01T00:00:00Z"`
	//EndTime    pgtype.Timestamptz `json:"end_time" example:"2020-01-01T00:00:00Z"`
}

type ReadTask struct {
	UUID       pgtype.UUID        `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	PeopleUUID pgtype.UUID        `json:"people_uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name       string             `json:"name" example:"do nothing"`
	StartTime  pgtype.Timestamptz `json:"start_time" example:"2020-01-01T00:00:00Z" swaggertype:"string" format:"timestamp"`
	EndTime    pgtype.Timestamptz `json:"end_time" example:"2020-01-01T00:00:00Z" swaggertype:"string" format:"timestamp"`
}
