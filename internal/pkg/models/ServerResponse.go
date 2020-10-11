package models

import (
	"encoding/json"
	"net/http"
)

type IncorrectGetParameters struct{}

func(t IncorrectGetParameters) Error()string{
	return "Incorrect get parameters!"
}

type ServerResponse struct{
	StatusCode int
	Response string
}

func BadBodyHTTPResponse(w *http.ResponseWriter, err error){
	response, err := json.Marshal(ServerResponse{
		StatusCode: 401,
		Response:  err.Error(),
	})

	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Write(response)
}

func BadMethodHttpResponse(w *http.ResponseWriter){
	response, _ := json.Marshal(ServerResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Response:  "MethodNotAllowed!",
	})

	(*w).WriteHeader(http.StatusMethodNotAllowed)
	(*w).Write(response)
}

func UnauthorizedHttpResponse(w *http.ResponseWriter){
	response, _ := json.Marshal(ServerResponse{
		StatusCode: http.StatusUnauthorized,
		Response: "You not authorized!",
	})
	(*w).WriteHeader(http.StatusUnauthorized)
	(*w).Write(response)
}

