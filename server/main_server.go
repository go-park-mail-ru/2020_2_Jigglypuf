package main

import (
	authDelivery "authentication/delivery"
	authRepository "authentication/repository"
	authUseCase"authentication/usecase"
	cinemaDelivery "cinemaService/delivery"
	cinemaRepository "cinemaService/repository"
	cinemaUsecase "cinemaService/usecase"
	"cookie"
	"log"
	"net/http"
	"sync"
	"time"
)

const StaticPath = "../static/"
const salt = "oisndoiqwe123"

type ServerStruct struct{
	authHandler *authDelivery.UserHandler
	cinemaHandler *cinemaDelivery.CinemaHandler
	httpServer *http.Server
}

func configureAPI() *ServerStruct{
	mutex := sync.RWMutex{}
	userRepository := authRepository.NewUserRepository(&mutex)
	cinRepository := cinemaRepository.NewCinemaRepository(&mutex)
	cinUseCase := cinemaUsecase.NewCinemaUseCase(cinRepository)
	userUseCase := authUseCase.NewUserUseCase(userRepository, salt)
	cinHandler := cinemaDelivery.NewCinemaHandler(cinUseCase)
	userHandler := authDelivery.NewUserHandler(userUseCase)

	return &ServerStruct{
		authHandler: userHandler,
		cinemaHandler: cinHandler,
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
	handler.HandleFunc("/getCinemaList/", application.cinemaHandler.GetCinemaList)
	handler.HandleFunc("/getCinema/", application.cinemaHandler.GetCinema)

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