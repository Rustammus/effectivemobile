package taskService

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/pkg/logging"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskService struct {
	Repo   repos.TaskRepository
	Logger logging.Logger
}

func (r *TaskService) StartNew(dto dto.CreateTask) (pgtype.UUID, error) {

	uuid, err := r.Repo.Create(context.TODO(), dto)
	if err != nil {
		r.Logger.Infof("error on create task %s", err.Error())
		return pgtype.UUID{}, err
	}
	r.Logger.Debugf("created and start task %x", uuid.Bytes)
	return uuid, nil
}

func (r *TaskService) Stop(uuid pgtype.UUID) (dto.ReadTask, error) {
	task, err := r.Repo.UpdateTaskStop(context.TODO(), uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.Logger.Debugf("no rows updated by uuid: %x", uuid.Bytes)
			return dto.ReadTask{}, err
		}
		r.Logger.Infof("error on stop task %s", err.Error())
		return dto.ReadTask{}, err
	}
	r.Logger.Debugf("stoped task %+v", task)
	return task, nil
}

func (r *TaskService) ListByPeople(uuid pgtype.UUID) ([]dto.ReadTask, error) {
	tasks, err := r.Repo.ListByPeopleUUID(context.TODO(), uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.Logger.Debugf("no rows found by people: %x", uuid.Bytes)
			return nil, err
		}
		r.Logger.Infof("error on list tasks by people %s", err.Error())
		return nil, err
	}
	r.Logger.Debugf("list tasks by people %+v", tasks)
	return tasks, nil
}
