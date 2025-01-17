package v1

import (
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/schemas"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) initTaskRouter(r *gin.RouterGroup) {
	p := r.Group("/task")
	{
		p.POST("", h.taskStart)
		p.PUT("/:uuid", h.taskStop)
		p.GET("", h.taskListByPeople)
	}
}

// taskStart godoc
// @Tags         Task API
// @Summary      Create and start task by People uuid
// @Description  Create and start task by People uuid
// @Accept       json
// @Produce      json
// @Param name query string false "Task name"
// @Param people query string true "People" format(uuid)
// @Success 200 {object} IResponseBase[schemas.ResponseUUID]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /task [post]
func (h *Handler) taskStart(c *gin.Context) {
	nameQ, ok := c.GetQuery("name")
	var name *string = &nameQ
	if !ok {
		name = nil
	}

	uuidS, ok := c.GetQuery("people")
	if !ok {
		IWriteResponseErr(c, 400, nil, "people query not matched")
		return
	}
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error scan people uuid")
		return
	}

	_, err = h.Services.People.FindByUUID(uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist")
			return
		}
	}

	uuidCreated, err := h.Services.Task.StartNew(dto.CreateTask{
		PeopleUUID: uuid,
		Name:       name,
	})
	if err != nil {
		IWriteResponseErr(c, 500, err, "error on create task")
		return
	}
	IWriteResponse(c, 200, schemas.ResponseUUID{uuidCreated}, "create task successfully")
}

// taskStop godoc
// @Tags         Task API
// @Summary      Stop task by uuid
// @Description  Stop task by uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "Task UUID" format(uuid)
// @Success 200 {object} IResponseBase[schemas.ResponseTask]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /task/{uuid} [put]
func (h *Handler) taskStop(c *gin.Context) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error scan task uuid param")
		return
	}

	taskDTO, err := h.Services.Task.Stop(uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist or task already stopped")
			return
		}
		IWriteResponseErr(c, 500, err, "error stop task")
		return
	}

	responseTask := schemas.ResponseTask{}
	responseTask.ScanDTO(taskDTO)

	IWriteResponse(c, 200, responseTask, "stop task successfully")
}

// taskListByPeople godoc
// @Tags         Task API
// @Summary      List tasks by people
// @Description  List tasks by people
// @Accept       json
// @Produce      json
// @Param people query string true "People UUID" format(uuid)
// @Success 200 {object} IResponseBaseMulti[schemas.ResponseTask]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /task [get]
func (h *Handler) taskListByPeople(c *gin.Context) {

	uuidS, ok := c.GetQuery("people")
	if !ok {
		IWriteResponseErr(c, 400, nil, "people query not matched")
		return
	}
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error scan people uuid")
		return
	}

	tasksDTO, err := h.Services.Task.ListByPeople(uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist or no tasks")
			return
		}
		IWriteResponseErr(c, 500, err, "error list tasks by people")
		return
	}

	responseTasks := make([]schemas.ResponseTask, len(tasksDTO), len(tasksDTO))
	for i := 0; i < len(tasksDTO); i++ {
		responseTasks[i].ScanDTO(tasksDTO[i])
	}

	IWriteResponse(c, 200, responseTasks, "people tasks read")
}
