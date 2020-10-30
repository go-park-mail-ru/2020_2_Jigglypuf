package main

import (
	_ "backend/docs"
	authService "backend/internal/app/authserver"
	cinemaService "backend/internal/app/cinemaserver"
	cookieService "backend/internal/app/cookieserver"
	movieService "backend/internal/app/movieserver"
	profileService "backend/internal/app/profileserver"
	scheduleService "backend/internal/app/scheduleserver"
	authConfig "backend/internal/pkg/authentication"
	cinemaConfig "backend/internal/pkg/cinemaservice"
	"backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/middleware/cookie/middleware"
	"backend/internal/pkg/middleware/cors"
	movieConfig "backend/internal/pkg/movieservice"
	profileConfig "backend/internal/pkg/profile"
	scheduleConfig"backend/internal/pkg/schedule"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/swaggo/http-swagger"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"time"
)

type ServerStruct struct {
	authService    *authService.AuthService
	cinemaService  *cinemaService.CinemaService
	movieService   *movieService.MovieService
	profileService *profileService.ProfileService
	cookieService  *cookieService.CookieService
	scheduleService *scheduleService.ScheduleService
	httpServer     *http.Server
}

func configureAPI(cookieDBConnection *tarantool.Connection, mainDBConnection *sql.DB) (*ServerStruct, error) {
	NewCookieService, cookieErr := cookieService.Start(cookieDBConnection)
	if cookieErr != nil {
		log.Println("No Tarantool Cookie DB connection")
		return nil, cookieErr
	}
	newAuthService, authErr := authService.Start(NewCookieService.CookieRepository, mainDBConnection)
	if authErr != nil{
		log.Println(authErr)
		return nil, authErr
	}
	newCinemaService, cinemaErr := cinemaService.Start(mainDBConnection)
	newMovieService, movieErr := movieService.Start(mainDBConnection, newAuthService.AuthenticationRepository)
	newProfileService, profileErr := profileService.Start(mainDBConnection)
	newScheduleService, scheduleErr := scheduleService.Start(mainDBConnection)

	if cinemaErr != nil || movieErr != nil || profileErr != nil || scheduleErr != nil {
		log.Println(cinemaErr)
		return nil, cinemaErr
	}
	return &ServerStruct{
		authService:    newAuthService,
		cinemaService:  newCinemaService,
		movieService:   newMovieService,
		profileService: newProfileService,
		cookieService:  NewCookieService,
		scheduleService: newScheduleService,
	}, nil
}

func configureRouter(application *ServerStruct) http.Handler {
	handler := http.NewServeMux()

	handler.Handle(movieConfig.URLPattern, application.movieService.MovieRouter)
	handler.Handle(cinemaConfig.URLPattern, application.cinemaService.CinemaRouter)
	handler.Handle(authConfig.URLPattern, application.authService.AuthRouter)
	handler.Handle(profileConfig.URLPattern, application.profileService.ProfileRouter)
	handler.Handle(scheduleConfig.URLPattern, application.scheduleService.Router)

	handler.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI, http.StatusMovedPermanently)
	})
	handler.HandleFunc("/docs/", httpSwagger.WrapHandler)
	middlewareHandler := middleware.CookieMiddleware(handler)
	middlewareHandler = cors.MiddlewareCORS(middlewareHandler)

	return middlewareHandler
}

func configureServer(port string, funcHandler http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      funcHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func startDBWork() (*sql.DB, *tarantool.Connection, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		authConfig.Host, authConfig.Port, authConfig.User, authConfig.Password, authConfig.DBName)

	PostgreSQLConnection, DBErr := sql.Open("postgres", psqlInfo)
	if DBErr != nil {
		return nil, nil, errors.New("no postgresql connection")
	}

	TarantoolConnection, DBConnectionErr := tarantool.Connect(cookie.Host+cookie.Port, tarantool.Opts{
		User: cookie.User,
		Pass: cookie.Password,
	})
	if DBConnectionErr != nil {
		return nil, nil, errors.New("no tarantool connection")
	}

	return PostgreSQLConnection, TarantoolConnection, nil
}

// Backend doc
// @title CinemaScope Backend API
// @version 0.5
// @description This is a backend API
// @host https://cinemascope.space
// @BasePath /
func main() {
	mainDBConnection, cookieDBConnection, DBErr := startDBWork()
	if DBErr != nil {
		log.Fatalln(DBErr)
		return
	}

	serverConfig, configErr := configureAPI(cookieDBConnection, mainDBConnection)
	if configErr != nil {
		log.Fatalln(configErr)
		return
	}

	responseHandler := configureRouter(serverConfig)
	serverConfig.httpServer = configureServer("8080", responseHandler)
	log.Println("Starting server at port 8080")
	err := serverConfig.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if mainDBConnection != nil {
			_ = mainDBConnection.Close()
		}
		if cookieDBConnection != nil {
			_ = cookieDBConnection.Close()
		}
	}()
}
