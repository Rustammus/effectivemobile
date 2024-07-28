package route

import (
	v1 "EffectiveMobile/internal/route/api/v1"
	"EffectiveMobile/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) Init(r *gin.Engine) {
	handlerV1 := v1.NewHandler(h.Services)
	apiV1 := r.Group("/api/v1")
	handlerV1.Init(apiV1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
