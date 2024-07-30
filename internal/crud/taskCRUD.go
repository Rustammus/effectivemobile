package crud

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/pkg/client/postgres"
	"EffectiveMobile/pkg/logging"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskCRUD struct {
	client postgres.Client
	logger logging.Logger
}

func (c *TaskCRUD) Create(ctx context.Context, dto dto.CreateTask) (pgtype.UUID, error) {
	q := `INSERT INTO public.tasks (people_uuid, name, start_time) 
		  VALUES ($1, $2, CURRENT_TIMESTAMP(0)) 
		  RETURNING uuid`
	var row pgx.Row
	uuid := pgtype.UUID{}
	if len(dto.Name) > 0 {
		row = c.client.QueryRow(ctx, q, dto.PeopleUUID, dto.Name)
	} else {
		row = c.client.QueryRow(ctx, q, dto.PeopleUUID, nil)
	}
	err := row.Scan(&uuid)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func (c *TaskCRUD) ListByPeopleUUID(ctx context.Context, uuid pgtype.UUID) ([]dto.ReadTask, error) {
	q := `SELECT uuid, people_uuid, name, start_time, end_time 
		  FROM public.tasks 
		  WHERE people_uuid = $1 
		  ORDER BY (end_time - start_time) DESC`

	rows, err := c.client.Query(ctx, q, uuid)
	defer rows.Close()
	if err != nil {
		return []dto.ReadTask{}, err
	}
	tasks := make([]dto.ReadTask, 0)
	for rows.Next() {
		task := dto.ReadTask{}
		err = rows.Scan(&task.UUID, &task.PeopleUUID, &task.Name, &task.StartTime, &task.EndTime)
		if err != nil {
			return []dto.ReadTask{}, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (c *TaskCRUD) UpdateTaskStop(ctx context.Context, uuid pgtype.UUID) (dto.ReadTask, error) {
	q := `UPDATE public.tasks SET end_time = CURRENT_TIMESTAMP(0) WHERE end_time IS NULL AND uuid = $1 RETURNING uuid, people_uuid, name, start_time, end_time`
	task := dto.ReadTask{}
	err := c.client.QueryRow(ctx, q, uuid).Scan(&task.UUID, &task.PeopleUUID, &task.Name, &task.StartTime, &task.EndTime)
	if err != nil {
		return dto.ReadTask{}, err
	}
	return task, nil
}

func NewTaskCRUD(logger logging.Logger, client postgres.Client) *TaskCRUD {
	return &TaskCRUD{client: client, logger: logger}
}
