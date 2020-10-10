package main

import (
	authDelivery "backend/internal/pkg/authentication/delivery"
	authRepository "backend/internal/pkg/authentication/repository"
	authUseCase "backend/internal/pkg/authentication/usecase"
	cinemaDelivery "backend/internal/pkg/cinemaService/delivery"
	cinemaRepository "backend/internal/pkg/cinemaService/repository"
	cinemaUsecase "backend/internal/pkg/cinemaService/usecase"
	movieDelivery "backend/internal/pkg/movieService/delivery"
	movieRepository "backend/internal/pkg/movieService/repository"
	movieUsecase "backend/internal/pkg/movieService/usecase"
	profileDelivery "backend/internal/pkg/profile/delivery"
	profileRepository "backend/internal/pkg/profile/repository"
	profileUseCase "backend/internal/pkg/profile/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const StaticPath = "../../static/2020_2_Jigglypuff/public/"
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
	movHandler := movieDelivery.NewMovieHandler(movUseCase, userRepository)
	profHandler := profileDelivery.NewProfileHandler(profUseCase)

	return &ServerStruct{
		authHandler: userHandler,
		cinemaHandler: cinHandler,
		movieHandler: movHandler,
		profileHandler: profHandler,
	}
}
func mainHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")

	file, err := os.Open(StaticPath + "index.html")
	if err != nil{
		_, responseErr := w.Write([]byte(`404 NOT FOUND`))
		log.Println("Cannot open index.html file")
		if responseErr != nil{
			log.Println(err)
		}
	}
	defer file.Close()

	streamBytes, inputErr := ioutil.ReadAll(file)
	if inputErr != nil{
		log.Println(inputErr)
	}

	_, outputErr := w.Write(streamBytes)
	if outputErr != nil{
		log.Println(outputErr)
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
	handler.HandleFunc("/ratemovie/", func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.movieHandler.RateMovie)
	})
	handler.HandleFunc("/getmovierating/", func(w http.ResponseWriter, r *http.Request){
		CORSDecorator(w,r, application.movieHandler.GetMovieRating)
	})

	handler.HandleFunc("/", mainHandler)


	mediaHandler := http.StripPrefix(
		"/media/",
		http.FileServer(http.Dir(MediaPath)),
	)
	handler.Handle("/media/", mediaHandler)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(StaticPath)))
	handler.Handle("/static/", staticHandler)
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