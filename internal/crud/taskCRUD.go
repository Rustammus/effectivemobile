package crud

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/pkg/client/postgres"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskCRUD struct {
	client postgres.Client
}

func (c TaskCRUD) Create(ctx context.Context, dto dto.CreateTaskDTO) (pgtype.UUID, error) {
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

func (c TaskCRUD) FindByPeopleUUID(ctx context.Context, uuid pgtype.UUID) ([]dto.ReadTaskDTO, error) {
	q := `SELECT uuid, people_uuid, name, start_time, end_time 
		  FROM public.tasks 
		  WHERE people_uuid = $1 
		  ORDER BY (end_time - start_time) DESC`

	rows, err := c.client.Query(ctx, q, uuid)
	defer rows.Close()
	if err != nil {
		return []dto.ReadTaskDTO{}, err
	}
	tasks := make([]dto.ReadTaskDTO, 0)
	for rows.Next() {
		task := dto.ReadTaskDTO{}
		err = rows.Scan(&task.UUID, &task.PeopleUUID, &task.Name, &task.StartTime, &task.EndTime)
		if err != nil {
			return []dto.ReadTaskDTO{}, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
