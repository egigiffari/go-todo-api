package exception

import (
	"net/http"

	"github.com/egigiffari/go-todo-api.git/constants"
)

func Unauthorized(resource string) Contract {
	return &Exception{
		code: http.StatusUnauthorized,
		messages: map[string][]string{
			resource: {constants.ErrUnauthorized},
		},
	}
}
