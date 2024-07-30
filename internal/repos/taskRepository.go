package repos

import (
	"EffectiveMobile/internal/dto"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskRepository interface {
	Create(ctx context.Context, dto dto.CreateTask) (pgtype.UUID, error)
	ListByPeopleUUID(ctx context.Context, uuid pgtype.UUID) ([]dto.ReadTask, error)
	UpdateTaskStop(ctx context.Context, uuid pgtype.UUID) (dto.ReadTask, error)
}
