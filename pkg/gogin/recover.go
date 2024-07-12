package gogin

import (
	"net/http"

	"github.com/egigiffari/go-todo-api.git/pkg/gogin/exception"
	"github.com/egigiffari/go-todo-api.git/pkg/gogin/response"
	"github.com/gin-gonic/gin"
)

func Recovery(ctx *gin.Context, err any) {
	switch exc := err.(type) {
	case exception.Contract:
		ctx.AbortWithStatusJSON(
			exc.Code(),
			response.Error{
				Errors:    exc.Messages(),
				TraceId:   ctx.GetHeader("x-trace-id"),
				Timestamp: ctx.GetHeader("x-timestamp"),
			},
		)
	default:
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
