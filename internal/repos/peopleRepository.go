package repos

import (
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/dto"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type PeopleRepository interface {
	Create(ctx context.Context, dto dto.CreatePeopleDTO) (pgtype.UUID, error)
	FindAll(ctx context.Context) ([]dto.ReadPeopleDTO, error)
	FindAllByOffset(ctx context.Context, pag crud.Pagination) ([]dto.ReadPeopleDTO, error)
	FindByFilterOffset(ctx context.Context, filter dto.FilterPeopleDTO, pag crud.Pagination) ([]dto.ReadPeopleDTO, error)
	FindByUUID(ctx context.Context, uuid pgtype.UUID) (dto.ReadPeopleDTO, error)
	Update(ctx context.Context, uuid pgtype.UUID, user dto.UpdatePeopleDTO) (dto.ReadPeopleDTO, error)
	Delete(ctx context.Context, uuid pgtype.UUID) (pgtype.UUID, error)
}
