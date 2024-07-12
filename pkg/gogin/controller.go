package gogin

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/egigiffari/go-todo-api.git/constants"
	"github.com/egigiffari/go-todo-api.git/pkg/gogin/exception"
	"github.com/egigiffari/go-todo-api.git/pkg/gogin/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (ctrl *Controller) Ok(ctx *gin.Context, data any) {
	ctx.JSON(
		http.StatusOK,
		response.Ok{
			Data:      data,
			TraceId:   ctx.GetHeader("x-trace-id"),
			Timestamp: ctx.GetHeader("x-timestamp"),
		},
	)
}

func (ctrl *Controller) Paginate(ctx *gin.Context, data any, total int) {
	page, size := ctrl.getPageAndSize(ctx)
	lastPage := ctrl.getLastPage(size, total)
	ctx.JSON(
		http.StatusOK,
		response.Ok{
			Data: response.Paginate{
				Data:     data,
				Page:     page,
				LastPage: lastPage,
				Size:     size,
				Total:    total,
			},
			TraceId:   ctx.GetHeader("x-trace-id"),
			Timestamp: ctx.GetHeader("x-timestamp"),
		},
	)
}

func (ctrl *Controller) NoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
}

func (ctrl *Controller) Created(ctx *gin.Context, data any) {
	ctx.JSON(
		http.StatusCreated,
		response.Ok{
			Data:      data,
			TraceId:   ctx.GetHeader("x-trace-id"),
			Timestamp: ctx.GetHeader("x-timestamp"),
		},
	)
}

func (ctrl *Controller) GetToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")

	if len(token) <= 0 {
		panic(exception.Forbidden("token", []string{constants.ErrInvalidToken}))
	}

	return token
}

func (ctrl *Controller) BindForm(ctx *gin.Context, req any) {
	if err := ctx.ShouldBind(req); err != nil {
		panic(exception.BadRequest(map[string][]string{
			"request": {constants.ErrInvalidRequest},
		}))
	}
}

func (ctrl *Controller) BindJSON(ctx *gin.Context, req any) {
	if err := ctx.ShouldBindJSON(req); err != nil {
		panic(exception.BadRequest(map[string][]string{
			"request": {constants.ErrInvalidRequest},
		}))
	}
}

func (ctrl *Controller) getPageAndSize(ctx *gin.Context) (int, int) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "20"))
	return page, size
}

func (ctrl *Controller) getLastPage(size int, total int) (page int) {
	perPage := float64(total) / float64(size)
	lastPage := int(math.Ceil(perPage))

	if lastPage <= 0 {
		lastPage = 1
	}

	return int(lastPage)
}
