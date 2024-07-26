package v1

import (
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/dto"
	"EffectiveMobile/internal/schemas/peopleSchemas"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strconv"
)

func (h *Handler) initPeopleRouter(r *gin.RouterGroup) {
	p := r.Group("/people")
	{
		p.POST("", h.peopleCreate)
		p.GET("/:uuid", h.peopleFindByUUID)
		p.GET("/list", h.peopleListByFilter)
		p.PUT("/:uuid", h.peopleUpdate)
		p.DELETE("/:uuid", h.peopleDelete)

		p.POST("/:uuid/start-task", h.peopleTaskStart)
		p.POST("/:uuid/end-task/:task-uuid", h.peopleTaskEnd)
		p.GET("/:uuid/tasks", h.peopleTaskList)
	}
}

func (h *Handler) peopleCreate(c *gin.Context) {
	peopleRequest := peopleSchemas.RequestCreatePeople{}

	err := c.ShouldBindJSON(&peopleRequest)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uuid, err := h.Services.People.Create(peopleRequest)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string      `json:"message"`
		UUID    pgtype.UUID `json:"uuid"`
	}{
		Message: "people created",
		UUID:    uuid,
	})

}

func (h *Handler) peopleFindByUUID(c *gin.Context) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	people, err := h.Services.People.FindByUUID(uuid)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(
		http.StatusOK,
		struct {
			Message string            `json:"message"`
			Peoples dto.ReadPeopleDTO `json:"people"`
		}{
			Message: "people found",
			Peoples: people,
		},
	)
}

func (h *Handler) peopleList(c *gin.Context) {
	offsetN, limitN := 0, 10
	var err error
	offsetS, ok := c.GetQuery("offset")
	limitS, ok2 := c.GetQuery("limit")
	if ok && ok2 {
		offsetN, err = strconv.Atoi(offsetS)
		limitN, err = strconv.Atoi(limitS)
		if err != nil {
			//TODO error response
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if limitN > 40 {
			limitN = 40
		}
	}

	peoples, newPag, err := h.Services.People.FindAllByOffset(crud.Pagination{
		Offset:    offsetN,
		Limit:     limitN,
		Timestamp: pgtype.Timestamptz{},
	})
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(
		http.StatusOK,
		struct {
			Message        string              `json:"message"`
			Peoples        []dto.ReadPeopleDTO `json:"peoples"`
			NextPagination crud.Pagination     `json:"pagination"`
		}{
			Message:        "people read",
			Peoples:        peoples,
			NextPagination: newPag,
		},
	)
}

func (h *Handler) peopleListByFilter(c *gin.Context) {
	offsetN, limitN := 0, 10
	var err error
	offsetS, ok := c.GetQuery("offset")
	limitS, ok2 := c.GetQuery("limit")
	if ok && ok2 {
		offsetN, err = strconv.Atoi(offsetS)
		limitN, err = strconv.Atoi(limitS)
		if err != nil {
			//TODO error response
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if limitN > 40 {
			//TODO configurable max limit
			limitN = 40
		}
		if offsetN < 0 {
			offsetN = 0
		}
	}

	filter := peopleSchemas.RequestFilterPeople{}
	pagination := crud.Pagination{}
	err = c.BindQuery(&filter)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = c.BindQuery(&pagination)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	peoples, nextPag, err := h.Services.People.FindByFilterOffset(filter, pagination)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(
		http.StatusOK,
		struct {
			Message        string              `json:"message"`
			Peoples        []dto.ReadPeopleDTO `json:"peoples"`
			NextPagination crud.Pagination     `json:"next_pagination"`
		}{
			Message:        "people read with filter",
			Peoples:        peoples,
			NextPagination: nextPag,
		},
	)
}

func (h *Handler) peopleUpdate(c *gin.Context) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uPeople := peopleSchemas.RequestUpdatePeople{}

	err = c.ShouldBindJSON(&uPeople)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rPeople, err := h.Services.People.UpdateByUUID(uuid, uPeople)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(
		http.StatusOK,
		struct {
			Message string            `json:"message"`
			Peoples dto.ReadPeopleDTO `json:"peoples"`
		}{
			Message: "people updated",
			Peoples: rPeople,
		},
	)
}

func (h *Handler) peopleDelete(c *gin.Context) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}
	err := uuid.Scan(uuidS)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	uuid, err = h.Services.People.DeleteByUUID(uuid)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(
		http.StatusOK,
		struct {
			Message string      `json:"message"`
			UUID    pgtype.UUID `json:"uuid"`
		}{
			Message: "people deleted",
			UUID:    uuid,
		},
	)
}

func (h *Handler) peopleTaskStart(c *gin.Context) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}

	err := uuid.Scan(uuidS)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//TODO create task by people uuid
}

func (h *Handler) peopleTaskEnd(c *gin.Context) {
	peopleUUIDstr := c.Param("uuid")
	taskUUIDstr := c.Param("task-uuid")
	peopleUUID := pgtype.UUID{}
	taskUUID := pgtype.UUID{}
	err := peopleUUID.Scan(peopleUUIDstr)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = taskUUID.Scan(taskUUIDstr)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//TODO end task by people uuid
}

func (h *Handler) peopleTaskList(c *gin.Context) {
	uuidS := c.Param("uuid")
	uuid := pgtype.UUID{}

	err := uuid.Scan(uuidS)
	if err != nil {
		//TODO error response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//TODO list all task by people uuid order by time
}
