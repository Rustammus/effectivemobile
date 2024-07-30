package v1

import (
	"EffectiveMobile/internal/crud"
	"github.com/gin-gonic/gin"
)

type IResponseBase[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type IResponseBaseErr struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}

type IResponseBasePaginated[T any] struct {
	Message    string          `json:"message"`
	Pagination crud.Pagination `json:"next_pagination"`
	Data       []T             `json:"data"`
}

type IResponseBaseMulti[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

func IWriteResponse[DataType any](c *gin.Context, code int, data DataType, message string) {
	resp := IResponseBase[DataType]{
		Message: message,
		Data:    data,
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

func IWriteResponsePaginated[DataType any](c *gin.Context, code int, data []DataType, nextPagination crud.Pagination, message string) {
	resp := IResponseBasePaginated[DataType]{
		Message:    message,
		Pagination: nextPagination,
		Data:       data,
	}
	c.JSON(code, resp)
}

func IWriteResponseMulti[DataType any](c *gin.Context, code int, data []DataType, message string) {
	resp := IResponseBaseMulti[DataType]{
		Message: message,
		Data:    data,
	}
	c.JSON(code, resp)
}
