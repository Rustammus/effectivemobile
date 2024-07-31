package v1

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/schemas"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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
// @Param People body  schemas.RequestCreatePeople true "People base"
// @Success 200 {object} IResponseBase[schemas.ResponseUUID]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /people [post]
func (h *Handler) peopleCreate(c *gin.Context) {
	peopleRequest := schemas.RequestCreatePeople{}

	err := c.ShouldBindJSON(&peopleRequest)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on binding JSON")
		return
	}

	uuid, err := h.Services.People.Create(peopleRequest)
	if err != nil {
		IWriteResponseErr(c, 500, err, "error on create people")
		return
	}

	IWriteResponse(c, 200, schemas.ResponseUUID{uuid}, "people created")
}

// peopleFindByUUID godoc
// @Tags         People API
// @Summary      Find People by uuid
// @Description  Find People by uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "People UUID" format(uuid)
// @Success 200 {object} IResponseBase[schemas.ResponsePeople]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /people/{uuid} [get]
func (h *Handler) peopleFindByUUID(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on scan param 'uuid")
		return
	}

	peopleDTO, err := h.Services.People.FindByUUID(uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist")
			return
		}
		IWriteResponseErr(c, 500, err, "error on list people")
		return
	}

	responsePeople := schemas.ResponsePeople{}
	responsePeople.ScanDTO(peopleDTO)

	IWriteResponse(c, 200, responsePeople, "people found")
}

// peopleListByFilter godoc
// @Tags         People API
// @Summary      List Peoples by filter
// @Description  List Peoples by filter
// @Accept       json
// @Produce      json
// @Param PeopleFilter query schemas.RequestFilterPeople true "People base"
// @Param Pagination query crud.Pagination true "Pagination base"
// @Success      200  {object}  IResponseBasePaginated[schemas.ResponsePeople]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /people [get]
func (h *Handler) peopleListByFilter(c *gin.Context) {
	var err error

	filter := schemas.RequestFilterPeople{}
	pagination := crud.Pagination{}
	err = c.ShouldBindQuery(&filter)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on bind filter queries")
		return
	}

	err = c.ShouldBindQuery(&pagination)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on bind pagination queries")
		return
	}

	uuidS, ok := c.GetQuery("uuid")
	if ok {
		err = filter.UUID.Scan(uuidS)
		if err != nil {
			IWriteResponseErr(c, 400, err, "error on scan uuid query")
			return
		}
	}

	if pagination.Limit > maxRowLimit || pagination.Limit < 1 {
		pagination.Limit = maxRowLimit
	}
	if pagination.Offset < 0 {
		pagination.Offset = 0
	}

	var peoples []dto.ReadPeople
	var nextPag crud.Pagination
	if !filter.Valid() {
		peoples, nextPag, err = h.Services.People.FindAllByOffset(pagination)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				IWriteResponseErr(c, 200, err, "no peoples found")
				return
			}
			IWriteResponseErr(c, 500, err, "error on list people")
			return
		}
	} else {
		peoples, nextPag, err = h.Services.People.FindByFilterOffset(filter, pagination)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				IWriteResponseErr(c, 200, err, "no peoples found by filter")
				return
			}
			IWriteResponseErr(c, 500, err, "error on list people")
			return
		}
	}

	responsePeoples := make([]schemas.ResponsePeople, len(peoples), len(peoples))
	for i := 0; i < len(peoples); i++ {
		responsePeoples[i].ScanDTO(peoples[i])
	}

	IWriteResponsePaginated(c, 200, responsePeoples, nextPag, "peoples found by filter")
}

// peopleUpdate godoc
// @Tags         People API
// @Summary      Update people
// @Description  Update people
// @Accept       json
// @Produce      json
// @Param UpdatePeople body schemas.RequestUpdatePeople true "People base"
// @Param uuid path string false "People UUID" format(uuid)
// @Success 	 200 {object} IResponseBase[schemas.ResponsePeople]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /people/{uuid} [put]
func (h *Handler) peopleUpdate(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on scan param 'uuid'")
		return
	}

	uPeople := schemas.RequestUpdatePeople{}
	err = c.ShouldBindJSON(&uPeople)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on bind people update body")
		return
	}
	if !uPeople.Valid() {
		IWriteResponseErr(c, 400, err, "update people body required")
		return
	}

	peopleDTO, err := h.Services.People.UpdateByUUID(uuid, uPeople)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist")
			return
		}
		IWriteResponseErr(c, 500, err, "error on update people")
		return
	}

	responsePeople := schemas.ResponsePeople{}
	responsePeople.ScanDTO(peopleDTO)

	IWriteResponse(c, 200, responsePeople, "people updated")
}

// peopleDelete godoc
// @Tags         People API
// @Summary      Delete people
// @Description  Delete people
// @Accept       json
// @Produce      json
// @Param uuid path string false "People UUID" format(uuid)
// @Success 200 {object} IResponseBase[schemas.ResponseUUID]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /people/{uuid} [delete]
func (h *Handler) peopleDelete(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on scan param 'uuid'")
		return
	}

	uuid, err = h.Services.People.DeleteByUUID(uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist")
			return
		}
		IWriteResponseErr(c, 500, err, "error on delete people")
		return
	}
	IWriteResponse(c, 200, schemas.ResponseUUID{uuid}, "people deleted")
}

// peopleTaskStart godoc
// @Tags         People API
// @Summary      Create and start task by People uuid
// @Description  Create and start task by People uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "People UUID" format(uuid)
// @Param name query string false "Task name"
// @Success 200 {object} IResponseBase[schemas.ResponseUUID]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /people/{uuid}/start-task [post]
func (h *Handler) peopleTaskStart(c *gin.Context) {
	nameQ, ok := c.GetQuery("name")
	var name *string = &nameQ
	if !ok {
		name = nil
	}

	uuid, err := scanUUIDParam(c)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on scan param 'uuid'")
		return
	}

	_, err = h.Services.People.FindByUUID(uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist")
			return
		}
	}

	taskUUID, err := h.Services.Task.StartNew(dto.CreateTask{
		PeopleUUID: uuid,
		Name:       name,
	})
	if err != nil {
		IWriteResponseErr(c, 500, err, "error on create task")
		return
	}
	IWriteResponse(c, 200, schemas.ResponseUUID{taskUUID}, "people tasks started")
}

// peopleTaskList godoc
// @Tags         People API
// @Summary      List all tasks by People uuid
// @Description  List all tasks by People uuid
// @Accept       json
// @Produce      json
// @Param uuid path string true "People UUID" format(uuid)
// @Success 200 {object} IResponseBaseMulti[schemas.ResponseTask]
// @Failure      400  {object}	IResponseBaseErr
// @Failure      500  {object}	IResponseBaseErr
// @Router       /people/{uuid}/tasks [get]
func (h *Handler) peopleTaskList(c *gin.Context) {
	uuid, err := scanUUIDParam(c)
	if err != nil {
		IWriteResponseErr(c, 400, err, "error on scan param 'uuid'")
		return
	}

	tasksDTO, err := h.Services.Task.ListByPeople(uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			IWriteResponseErr(c, 200, err, "uuid not exist or no tasks found")
			return
		}
		IWriteResponseErr(c, 500, err, "error on list people tasksDTO")
		return
	}

	responseTasks := make([]schemas.ResponseTask, len(tasksDTO), len(tasksDTO))
	for i := 0; i < len(tasksDTO); i++ {
		responseTasks[i].ScanDTO(tasksDTO[i])
	}

	IWriteResponse(c, 200, responseTasks, "people tasksDTO read")
}

func scanUUIDParam(c *gin.Context) (pgtype.UUID, error) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	return uuid, err
}
