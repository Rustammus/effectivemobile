package v1

import (
	"EffectiveMobile/internal/crud"
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

type IResponseBase[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type IResponseBaseErr struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}

type IResponseBaseMulti[T any] struct {
	Message    string          `json:"message"`
	Pagination crud.Pagination `json:"next_pagination"`
	Data       []T             `json:"data"`
}

func IWriteResponse[DataType any](c *gin.Context, code int, data DataType, message string) {
	resp := IResponseBase[DataType]{
		Message: message,
		Data:    data,
	}
	c.JSON(code, resp)
}
func IWriteResponseMulti[DataType any](c *gin.Context, code int, data []DataType, nextPagination crud.Pagination, message string) {
	resp := IResponseBaseMulti[DataType]{
		Message:    message,
		Pagination: nextPagination,
		Data:       data,
	}
	c.JSON(code, resp)
}

func IWriteResponseErr(c *gin.Context, code int, err error, message string) {
	resp := IResponseBaseErr{
		Message: message,
		Error:   err,
	}
	c.JSON(code, resp)
}
