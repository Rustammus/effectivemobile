package repos

import (
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/dto"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type PeopleRepository interface {
	Create(ctx context.Context, dto dto.CreatePeople) (pgtype.UUID, error)
	FindAll(ctx context.Context) ([]dto.ReadPeople, error)
	FindAllByOffset(ctx context.Context, pag crud.Pagination) ([]dto.ReadPeople, error)
	FindByFilterOffset(ctx context.Context, filter dto.FilterPeople, pag crud.Pagination) ([]dto.ReadPeople, error)
	FindByUUID(ctx context.Context, uuid pgtype.UUID) (dto.ReadPeople, error)
	Update(ctx context.Context, uuid pgtype.UUID, user dto.UpdatePeople) (dto.ReadPeople, error)
	Delete(ctx context.Context, uuid pgtype.UUID) (pgtype.UUID, error)
}
