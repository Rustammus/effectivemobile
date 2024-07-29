package schemas

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type BaseResp struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
	Data    any    `json:"data"`
}

type ResponseUUID struct {
	UUID pgtype.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
}
