package models

import (
	"encoding/json"
	"net/http"
)

type ServerResponse struct{
	StatusCode int
	Response []byte
}

func BadBodyHTTPResponse(w *http.ResponseWriter, err error){
	response, err := json.Marshal(ServerResponse{
		StatusCode: 401,
		Response:  []byte(err.Error()),
	})

	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Write(response)
}

func BadMethodHttpResponse(w *http.ResponseWriter){
	response, _ := json.Marshal(ServerResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Response:  []byte("MethodNotAllowed"),
	})

	(*w).WriteHeader(http.StatusMethodNotAllowed)
	(*w).Write(response)
}

func UnauthorizedHttpResponse(w *http.ResponseWriter){
	response, _ := json.Marshal(ServerResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Response:  []byte("MethodNotAllowed"),
	})
	(*w).WriteHeader(http.StatusUnauthorized)
	(*w).Write(response)
}

