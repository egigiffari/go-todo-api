package exception

import (
	"net/http"
)

func Forbidden(resource string, message []string) Contract {
	return &Exception{
		code: http.StatusForbidden,
		messages: map[string][]string{
			resource: message,
		},
	}
}
