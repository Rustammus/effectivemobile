package taskService

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/pkg/logging"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskService struct {
	Repo   repos.TaskRepository
	Logger logging.Logger
}

func (r TaskService) StartNew(dto dto.CreateTask) (pgtype.UUID, error) {
	uuid, err := r.Repo.Create(context.TODO(), dto)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func (r TaskService) Stop(uuid pgtype.UUID) (dto.ReadTask, error) {
	task, err := r.Repo.UpdateTaskStop(context.TODO(), uuid)
	if err != nil {
		return dto.ReadTask{}, err
	}
	return task, nil
}

func (r TaskService) ListByPeople(uuid pgtype.UUID) ([]dto.ReadTask, error) {
	tasks, err := r.Repo.ListByPeopleUUID(context.TODO(), uuid)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
