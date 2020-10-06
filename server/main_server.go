package main

import (
	authDelivery "authentication/delivery"
	authRepository "authentication/repository"
	authUseCase"authentication/usecase"
	cinemaDelivery "cinemaService/delivery"
	cinemaRepository "cinemaService/repository"
	cinemaUsecase "cinemaService/usecase"
	movieDelivery "movieService/delivery"
	movieRepository "movieService/repository"
	movieUsecase "movieService/usecase"
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
	movieHandler *movieDelivery.MovieHandler

	httpServer *http.Server
}

func configureAPI() *ServerStruct{
	mutex := sync.RWMutex{}
	userRepository := authRepository.NewUserRepository(&mutex)
	cinRepository := cinemaRepository.NewCinemaRepository(&mutex)
	movRepository := movieRepository.NewMovieRepository(&mutex)

	cinUseCase := cinemaUsecase.NewCinemaUseCase(cinRepository)
	userUseCase := authUseCase.NewUserUseCase(userRepository, salt)
	movUseCase := movieUsecase.NewMovieUseCase(movRepository)

	cinHandler := cinemaDelivery.NewCinemaHandler(cinUseCase)
	userHandler := authDelivery.NewUserHandler(userUseCase)
	movHandler := movieDelivery.NewMovieHandler(movUseCase)

	return &ServerStruct{
		authHandler: userHandler,
		cinemaHandler: cinHandler,
		movieHandler: movHandler,
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
	handler.HandleFunc("/getMovie/", application.movieHandler.GetMovie)
	handler.HandleFunc("/getMovieList/", application.movieHandler.GetMovieList)

	staticHandler := http.StripPrefix(
		"/media/",
		http.FileServer(http.Dir("../../media")),
	)
	handler.Handle("/media/", staticHandler)

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