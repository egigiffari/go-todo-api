package exception

import "net/http"

func BadRequest(messages map[string][]string) Contract {
	return &Exception{
		code:     http.StatusBadRequest,
		messages: messages,
	}
}
