package v1

import (
	"EffectiveMobile/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) Init(r *gin.RouterGroup) {
	h.initPeopleRouter(r)
}
