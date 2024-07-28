package v1

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/schemas/peopleSchemas"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

var maxRowLimit = 40

func (h *Handler) initPeopleRouter(r *gin.RouterGroup) {
	maxRowLimit = config.GetConfig().Server.MaxRowLimit
	p := r.Group("/people")
	{
		p.POST("", h.peopleCreate)
		p.GET("/:uuid", h.peopleFindByUUID)
		p.GET("", h.peopleListByFilter)
		p.PUT("/:uuid", h.peopleUpdate)
		p.DELETE("/:uuid", h.peopleDelete)

		p.POST("/:uuid/start-task", h.peopleTaskStart)
		p.GET("/:uuid/tasks", h.peopleTaskList)
	}
}

// peopleCreate godoc
// @Tags         People API
// @Summary      Create people
// @Description  Create people
// @Accept       json
// @Produce      json
// @Param People body  peopleSchemas.RequestCreatePeople true "People base"
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /people [post]
func (h *Handler) peopleCreate(c *gin.Context) {
	peopleRequest := peopleSchemas.RequestCreatePeople{}

	err := c.ShouldBindJSON(&peopleRequest)
	if err != nil {
		writeResp404(c, err, "error on binding JSON")
		return
	}

	uuid, err := h.Services.People.Create(peopleRequest)
	if err != nil {
		writeResp500(c, err, "error on create people")
		return
	}

	writeResp200(c, uuid, "people created")
}

// peopleFindByUUID godoc
// @Tags         People API
// @Summary      Find People by uuid
// @Description  Find People by uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "People UUID" format(uuid)
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /people/{uuid} [get]
func (h *Handler) peopleFindByUUID(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		writeResp404(c, err, "error on scan param 'uuid'")
		return
	}

	people, err := h.Services.People.FindByUUID(uuid)
	if err != nil {
		writeResp404(c, err, "error on list people")
		return
	}

	writeResp200(c, people, "peoples found")
}

//TODO swagger pizdes

// peopleListByFilter godoc
// @Tags         People API
// @Summary      List Peoples by filter
// @Description  List Peoples by filter
// @Accept       json
// @Produce      json
// @Success      200  {object}   schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /people [get]
func (h *Handler) peopleListByFilter(c *gin.Context) {
	//offsetN, limitN := 0, 10
	var err error
	//TODO remove
	//offsetS, ok := c.GetQuery("offset")
	//limitS, ok2 := c.GetQuery("limit")
	//if ok && ok2 {
	//	offsetN, err = strconv.Atoi(offsetS)
	//	limitN, err = strconv.Atoi(limitS)
	//	if err != nil {
	//		writeResp404(c, err, "error on scan query 'offset' or 'limit'")
	//		return
	//	}
	//	if limitN > maxRowLimit {
	//		limitN = maxRowLimit
	//	}
	//	if offsetN < 0 {
	//		offsetN = 0
	//	}
	//}

	filter := peopleSchemas.RequestFilterPeople{}
	pagination := crud.Pagination{}
	err = c.BindQuery(&filter)
	if err != nil {
		writeResp404(c, err, "error on bind filter queries")
		return
	}

	err = c.BindQuery(&pagination)
	if err != nil {
		writeResp404(c, err, "error on bind filter queries")
		return
	}

	peoples, nextPag, err := h.Services.People.FindByFilterOffset(filter, pagination)
	if err != nil {
		writeResp500(c, err, "error on list people")
		return
	}

	resp := peopleSchemas.RespPeoplePaginated{
		Peoples:        peoples,
		NextPagination: nextPag,
	}
	writeResp200(c, resp, "peoples found")
}

// peopleUpdate godoc
// @Tags         People API
// @Summary      Update people
// @Description  Update people
// @Accept       json
// @Produce      json
// Param UpdatePeople body peopleSchemas.RequestUpdatePeople true "People base"
// @Param uuid path string false "People UUID" format(uuid)
// @Success 	 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /people/{uuid} [put]
func (h *Handler) peopleUpdate(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		writeResp404(c, err, "error on scan param 'uuid'")
		return
	}

	uPeople := peopleSchemas.RequestUpdatePeople{}
	err = c.ShouldBindJSON(&uPeople)
	if err != nil {
		writeResp404(c, err, "error on bind people update")
		return
	}

	rPeople, err := h.Services.People.UpdateByUUID(uuid, uPeople)
	if err != nil {
		writeResp500(c, err, "error on update people")
		return
	}

	writeResp200(c, rPeople, "peoples updated")
}

// peopleDelete godoc
// @Tags         People API
// @Summary      Delete people
// @Description  Delete people
// @Accept       json
// @Produce      json
// @Param uuid path string false "People UUID" format(uuid)
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /people/{uuid} [delete]
func (h *Handler) peopleDelete(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		writeResp404(c, err, "error on scan param 'uuid'")
		return
	}

	uuid, err = h.Services.People.DeleteByUUID(uuid)
	if err != nil {
		writeResp500(c, err, "error on update people")
		return
	}

	writeResp200(c, uuid, "peoples deleted")
}

// peopleTaskStart godoc
// @Tags         People API
// @Summary      Create and start task by People uuid
// @Description  Create and start task by People uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "People UUID" format(uuid)
// @Param name query string true "Task name"
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /people/{uuid}/start-task [post]
func (h *Handler) peopleTaskStart(c *gin.Context) {
	//TODO nil name
	name, _ := c.GetQuery("name")
	uuid, err := scanUUIDParam(c)
	if err != nil {
		writeResp404(c, err, "error on scan param 'uuid'")
		return
	}

	taskUUID, err := h.Services.Task.StartNew(dto.CreateTaskDTO{
		PeopleUUID: uuid,
		Name:       name,
	})
	if err != nil {
		writeResp500(c, err, "error on delete people")
		return
	}

	writeResp200(c, taskUUID, "people tasks read")
}

// peopleTaskList godoc
// @Tags         People API
// @Summary      List all tasks by People uuid
// @Description  List all tasks by People uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "People UUID" format(uuid)
// @Success 200 {object} schemas.BaseResp
// @Failure      404  {object}	schemas.BaseResp
// @Failure      500  {object}	schemas.BaseResp
// @Router       /people/{uuid}/tasks [get]
func (h *Handler) peopleTaskList(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		writeResp404(c, err, "error on scan uuid")
		return
	}

	tasks, err := h.Services.Task.ListByPeople(uuid)
	if err != nil {
		writeResp500(c, err, "error on delete people")
		return
	}

	writeResp200(c, tasks, "people tasks read")
}

func scanUUIDParam(c *gin.Context) (pgtype.UUID, error) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	return uuid, err
}
