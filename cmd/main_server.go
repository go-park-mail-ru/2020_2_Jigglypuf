package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-park-mail-ru/2020_2_Jigglypuf/docs"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/authserver"
	cinemaService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cinemaserver"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cookieserver"
	hallService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/hallserver"
	movieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/movieserver"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/profileserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/recserver"
	scheduleService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/scheduleserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/ticketservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	cinemaConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice"
	hallConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/csrf"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/logger"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	movieConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	profileConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	config "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	scheduleConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/middleware"
	ticketConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	_ "github.com/lib/pq"
	"github.com/swaggo/http-swagger"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"sync"
	"time"
)

type ServerStruct struct {
	authService     *authService.AuthService
	cinemaService   *cinemaService.CinemaService
	movieService    *movieService.MovieService
	profileService  *profileService.ProfileService
	cookieService   *cookieService.CookieService
	scheduleService *scheduleService.ScheduleService
	ticketService   *ticketservice.TicketService
	hallService     *hallService.HallService
	httpServer      *http.Server
	csrfMiddleware  *csrf.HashCSRFToken
	recommendationService *recserver.RecommendationService
}

func configureAPI(cookieDBConnection *tarantool.Connection, mainDBConnection *sql.DB) (*ServerStruct, error) {
	mutex := sync.RWMutex{}
	NewCookieService, cookieErr := cookieService.Start(cookieDBConnection)
	if cookieErr != nil {
		log.Println("No Tarantool Cookie DB connection")
		return nil, cookieErr
	}
	newProfileService, profileErr := profileService.Start(mainDBConnection)
	if profileErr != nil {
		return nil, profileErr
	}
	newAuthService, authErr := authService.Start(newProfileService.ProfileRepository, mainDBConnection)
	if authErr != nil {
		log.Println(authErr)
		return nil, authErr
	}
	newHallService, hallErr := hallService.Start(mainDBConnection)
	if hallErr != nil {
		log.Println(models.ErrFooInitFail)
		return nil, models.ErrFooInitFail
	}
	newCinemaService, cinemaErr := cinemaService.Start(mainDBConnection)
	newMovieService, movieErr := movieService.Start(mainDBConnection, newAuthService.AuthenticationRepository)
	newScheduleService, scheduleErr := scheduleService.Start(mainDBConnection)
	if scheduleErr != nil {
		log.Println(scheduleErr)
		return nil, models.ErrFooInitFail
	}
	newTicketService, ticketErr := ticketservice.Start(mainDBConnection, newAuthService.AuthenticationRepository, newHallService.Repository, newScheduleService.Repository)
	newHashCSRFMiddleware, csrfErr := csrf.NewHashCSRFToken(models.RandStringRunes(7), time.Hour*24)
	if cinemaErr != nil || movieErr != nil || ticketErr != nil || csrfErr != nil {
		log.Println(models.ErrFooInitFail)
		return nil, models.ErrFooInitFail
	}
	recommendationService, recErr := recserver.Start(mainDBConnection, &mutex, time.Minute*10)
	if recErr != nil{
		return nil,models.ErrFooInitFail
	}
	return &ServerStruct{
		authService:     newAuthService,
		cinemaService:   newCinemaService,
		movieService:    newMovieService,
		profileService:  newProfileService,
		cookieService:   NewCookieService,
		scheduleService: newScheduleService,
		ticketService:   newTicketService,
		hallService:     newHallService,
		csrfMiddleware:  newHashCSRFMiddleware,
		recommendationService: recommendationService,
	}, nil
}

func configureRouter(application *ServerStruct) http.Handler {
	handler := http.NewServeMux()

	handler.Handle(movieConfig.URLPattern, application.movieService.MovieRouter)
	handler.Handle(cinemaConfig.URLPattern, application.cinemaService.CinemaRouter)
	handler.Handle(configs.URLPattern, application.authService.AuthRouter)
	handler.Handle(profileConfig.URLPattern, application.profileService.ProfileRouter)
	handler.Handle(scheduleConfig.URLPattern, application.scheduleService.Router)
	handler.Handle(hallConfig.URLPattern, application.hallService.Router)
	handler.Handle(ticketConfig.URLPattern, application.ticketService.Router)
	handler.Handle(config.URLPattern, application.recommendationService.RecommendationRouter)
	handler.HandleFunc("/api/csrf/", application.csrfMiddleware.GenerateCSRFToken)

	handler.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI, http.StatusMovedPermanently)
	})
	handler.HandleFunc("/api/docs/", httpSwagger.WrapHandler)
	middlewareHandler := application.csrfMiddleware.CSRFMiddleware(handler)
	middlewareHandler = middleware.CookieMiddleware(middlewareHandler, application.cookieService.CookieDelivery)
	middlewareHandler = cors.MiddlewareCORS(middlewareHandler)
	middlewareHandler = logger.AccessLogMiddleware(middlewareHandler)
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
		configs.Host, configs.Port, configs.User, configs.Password, configs.DBName)

	PostgreSQLConnection, DBErr := sql.Open("postgres", psqlInfo)
	if DBErr != nil {
		return nil, nil, errors.New("no postgresql connection")
	}

	TarantoolConnection, DBConnectionErr := tarantool.Connect(session.Host+session.Port, tarantool.Opts{
		User: session.User,
		Pass: session.Password,
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
