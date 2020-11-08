package main

import (
	_ "backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func main() {
	http.HandleFunc("/docs/", httpSwagger.WrapHandler)
	http.ListenAndServe(":8081",nil)
}
