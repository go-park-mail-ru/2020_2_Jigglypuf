package main

import (
	authDelivery "authentication/delivery"
	authRepository "authentication/repository"
	authUseCase "authentication/usecase"
	cinemaDelivery "cinemaService/delivery"
	cinemaRepository "cinemaService/repository"
	cinemaUsecase "cinemaService/usecase"
	"log"
	movieDelivery "movieService/delivery"
	movieRepository "movieService/repository"
	movieUsecase "movieService/usecase"
	"net/http"
	profileDelivery "profile/delivery"
	profileRepository "profile/repository"
	profileUseCase "profile/usecase"
	"sync"
	"time"
)

const StaticPath = "../static/"
const MediaPath = "../../media/"
const salt = "oisndoiqwe123"


func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	log.Println(req.Header.Get("Origin"))
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func CORSDecorator(w http.ResponseWriter, r *http.Request, function func(w http.ResponseWriter, r *http.Request)){
	setupCORS(&w, r)
	if (*r).Method == http.MethodOptions {
		return
	}

	function(w,r)
}

type ServerStruct struct{
	authHandler *authDelivery.UserHandler
	cinemaHandler *cinemaDelivery.CinemaHandler
	movieHandler *movieDelivery.MovieHandler
	profileHandler *profileDelivery.ProfileHandler
	httpServer *http.Server
}

func configureAPI() *ServerStruct{
	mutex := sync.RWMutex{}
	userRepository := authRepository.NewUserRepository(&mutex)
	cinRepository := cinemaRepository.NewCinemaRepository(&mutex)
	movRepository := movieRepository.NewMovieRepository(&mutex)
	profRepository := profileRepository.NewProfileRepository(&mutex,userRepository)

	cinUseCase := cinemaUsecase.NewCinemaUseCase(cinRepository)
	userUseCase := authUseCase.NewUserUseCase(userRepository, salt)
	movUseCase := movieUsecase.NewMovieUseCase(movRepository)
	profUseCase := profileUseCase.NewProfileUseCase(profRepository)

	cinHandler := cinemaDelivery.NewCinemaHandler(cinUseCase)
	userHandler := authDelivery.NewUserHandler(userUseCase)
	movHandler := movieDelivery.NewMovieHandler(movUseCase)
	profHandler := profileDelivery.NewProfileHandler(profUseCase)

	return &ServerStruct{
		authHandler: userHandler,
		cinemaHandler: cinHandler,
		movieHandler: movHandler,
		profileHandler: profHandler,
	}
}

func configureRouter(application *ServerStruct) *http.ServeMux{
	handler := http.NewServeMux()
	handler.HandleFunc("/signin/", func(w http.ResponseWriter, r *http.Request){
		setupCORS(&w, r)
		if (*r).Method == http.MethodOptions {
			return
		}

		if /*cookie.CheckCookie(r) ||*/ r.Method != http.MethodPost{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		application.authHandler.AuthHandler(w,r)
	})

	handler.HandleFunc("/signup/", func(w http.ResponseWriter, r *http.Request){
		setupCORS(&w, r)
		if (*r).Method == http.MethodOptions {
			return
		}

		if /*cookie.CheckCookie(r) ||*/ r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		application.authHandler.RegisterHandler(w, r)
	})

	handler.HandleFunc("/getprofile/", func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.profileHandler.GetProfile)
	})
	handler.HandleFunc("/updateprofile/",func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.profileHandler.UpdateProfile)
	})

	handler.HandleFunc("/getcinemalist/", func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.cinemaHandler.GetCinemaList)
	})
	handler.HandleFunc("/getcinema/", func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.cinemaHandler.GetCinema)
	})
	handler.HandleFunc("/getmovie/", func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.movieHandler.GetMovie)
	})
	handler.HandleFunc("/getmovielist/", func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.movieHandler.GetMovieList)
	})


	staticHandler := http.StripPrefix(
		"/media/",
		http.FileServer(http.Dir(MediaPath)),
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