package models

import (
	"encoding/json"
	"net/http"
)

type IncorrectGetParameters struct{}

func (t IncorrectGetParameters) Error() string {
	return "Incorrect get parameters!"
}

type ServerResponse struct {
	StatusCode int
	Response   string
}

func BadBodyHTTPResponse(w *http.ResponseWriter, err error) {
	response, _ := json.Marshal(ServerResponse{
		StatusCode: http.StatusBadRequest,
		Response:   err.Error(),
	})

	(*w).WriteHeader(http.StatusBadRequest)
	_, _ = (*w).Write(response)
}

func BadMethodHTTPResponse(w *http.ResponseWriter) {
	response, _ := json.Marshal(ServerResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Response:   "MethodNotAllowed!",
	})

	(*w).WriteHeader(http.StatusMethodNotAllowed)
	_, _ = (*w).Write(response)
}

func UnauthorizedHTTPResponse(w *http.ResponseWriter) {
	response, _ := json.Marshal(ServerResponse{
		StatusCode: http.StatusUnauthorized,
		Response:   "You not authorized!",
	})
	(*w).WriteHeader(http.StatusUnauthorized)
	_, _ = (*w).Write(response)
}
