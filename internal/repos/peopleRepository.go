package repos

import (
	"EffectiveMobile/internal/dto"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type PeopleRepository interface {
	Create(ctx context.Context, dto dto.CreatePeopleDTO) (pgtype.UUID, error)
	FindAll(ctx context.Context) ([]dto.ReadPeopleDTO, error)
	FindByUUID(ctx context.Context, uuid pgtype.UUID) (*dto.ReadPeopleDTO, error)
	Update(ctx context.Context, user dto.UpdatePeopleDTO, uuid pgtype.UUID) (dto.ReadPeopleDTO, error)
	Delete(ctx context.Context, uuid pgtype.UUID) error
	SoftDelete(ctx context.Context, uuid pgtype.UUID) error
}
