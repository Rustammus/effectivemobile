package v1

import (
	"EffectiveMobile/internal/schemas"
	"github.com/gin-gonic/gin"
)

func writeResp404(c *gin.Context, err error, message string) {
	c.JSON(404, schemas.BaseResp{
		Message: message,
		Error:   err,
		Data:    nil})
}

func writeResp500(c *gin.Context, err error, message string) {
	c.JSON(500, schemas.BaseResp{
		Message: message,
		Error:   err,
		Data:    nil})
}

func writeResp200(c *gin.Context, data any, message string) {
	c.JSON(500, schemas.BaseResp{
		Message: message,
		Error:   nil,
		Data:    data})
}
