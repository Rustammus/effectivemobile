package schemas

import (
	"EffectiveMobile/internal/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResponseTask struct {
	UUID       pgtype.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	PeopleUUID pgtype.UUID `json:"people_uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name       *string     `json:"name,omitempty" example:"do nothing"`
	IsStopped  bool        `json:"isStopped" example:"true"`
	Hours      int         `json:"hours" example:"25"`
	Minutes    int         `json:"minutes" example:"59"`
	//StartTime  pgtype.Timestamptz `json:"start_time" example:"2020-01-01T00:00:00Z" swaggertype:"string" format:"timestamp"`
	//EndTime    pgtype.Timestamptz `json:"end_time" example:"2020-01-01T00:00:00Z" swaggertype:"string" format:"timestamp"`
}

func (t *ResponseTask) ScanDTO(dto dto.ReadTask) {
	minutes, hours := 0, 0
	if dto.EndTime.Valid {
		duration := dto.EndTime.Time.Sub(dto.StartTime.Time)
		minutes = int(duration.Minutes()) % 60
		hours = int(duration.Hours())
		t.IsStopped = true
	} else {
		t.IsStopped = false
	}
	t.UUID = dto.UUID
	t.PeopleUUID = dto.PeopleUUID
	t.Name = dto.Name
	t.Hours = hours
	t.Minutes = minutes
}
