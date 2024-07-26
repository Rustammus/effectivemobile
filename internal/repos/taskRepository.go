package repos

import (
	"EffectiveMobile/internal/dto"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskRepository interface {
	Create(ctx context.Context, dto dto.CreateTaskDTO) (pgtype.UUID, error)
	ListByPeopleUUID(ctx context.Context, uuid pgtype.UUID) ([]dto.ReadTaskDTO, error)
	UpdateTaskStop(ctx context.Context, uuid pgtype.UUID) (dto.ReadTaskDTO, error)
}
