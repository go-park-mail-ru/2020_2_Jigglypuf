package main

import (
	"authentication/delivery"
	"authentication/repository"
	"authentication/usecase"
	"cookie"
	"log"
	"net/http"
	"sync"
	"time"
)

const StaticPath = "../static/"
const salt = "oisndoiqwe123"

type ServerStruct struct{
	authHandler *delivery.UserHandler
	httpServer *http.Server
}

func configureAPI() *ServerStruct{
	mutex := sync.RWMutex{}
	userRepository := repository.NewUserRepository(&mutex)

	userUseCase := usecase.NewUserUseCase(userRepository, salt)
	userHandler := delivery.NewUserHandler(userUseCase)

	return &ServerStruct{
		authHandler: userHandler,
	}
}

func configureRouter(application *ServerStruct) *http.ServeMux{
	handler := http.NewServeMux()
	handler.HandleFunc("/signIn/", func(w http.ResponseWriter, r *http.Request){
		if cookie.CheckCookie(r){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		application.authHandler.AuthHandler(w,r)
	})

	handler.HandleFunc("/signUp/", application.authHandler.RegisterHandler)

	return handler
}

func configureServer(port string, funcHandler *http.ServeMux) *http.Server{
	return &http.Server{
		Addr: ":" + port,
		Handler: funcHandler,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func main(){
	serverConfig := configureAPI()
	responseHandler := configureRouter(serverConfig)
	serverConfig.httpServer = configureServer("8080",responseHandler)

	err := serverConfig.httpServer.ListenAndServe()

	if err != nil{
		log.Fatalln(err)
	}
}