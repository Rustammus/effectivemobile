package v1

import (
	"EffectiveMobile/internal/dto"
	"github.com/gin-gonic/gin"
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
// @Param people query string true "Task name" format(uuid)
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /task [post]
func (h *Handler) taskStart(c *gin.Context) {

	name, _ := c.GetQuery("name")
	uuidS, ok := c.GetQuery("people")
	if !ok {
		writeResp404(c, nil, "people query not matched")
		return
	}
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		writeResp404(c, err, "error scan people uuid")
		return
	}

	uuidCreated, err := h.Services.Task.StartNew(dto.CreateTaskDTO{
		PeopleUUID: uuid,
		Name:       name,
	})
	if err != nil {
		writeResp500(c, err, "error on create task")
		return
	}

	writeResp200(c, uuidCreated, "create task successfully")
}

// taskStop godoc
// @Tags         Task API
// @Summary      Stop task by uuid
// @Description  Stop task by uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "Task UUID" format(uuid)
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /task [put]
func (h *Handler) taskStop(c *gin.Context) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		writeResp404(c, err, "error scan task uuid param")
		return
	}

	task, err := h.Services.Task.Stop(uuid)
	if err != nil {
		writeResp500(c, err, "error scan task uuid param")
		return
	}

	writeResp200(c, task, "stop task successfully")
}

// taskListByPeople godoc
// @Tags         Task API
// @Summary      Stop task by uuid
// @Description  Stop task by uuid
// @Accept       json
// @Produce      json
// @Param people query string true "People UUID" format(uuid)
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /task [get]
func (h *Handler) taskListByPeople(c *gin.Context) {

	uuidS, ok := c.GetQuery("people")
	if !ok {
		writeResp404(c, nil, "people query not matched")
		return
	}
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		writeResp404(c, err, "error scan people uuid")
		return
	}

	tasks, err := h.Services.Task.ListByPeople(uuid)
	if err != nil {
		writeResp500(c, err, "error list tasks by people")
		return
	}

	writeResp200(c, tasks, "people tasks read")
}
