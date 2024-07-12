package exception

import (
	"net/http"

	"github.com/egigiffari/go-todo-api.git/constants"
)

func NotFound(resource string) Contract {
	return &Exception{
		code: http.StatusNotFound,
		messages: map[string][]string{
			resource: {constants.ErrNotFound},
		},
	}
}
