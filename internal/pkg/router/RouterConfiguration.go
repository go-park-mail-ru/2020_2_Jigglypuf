package router

import (
	"database/sql"
	"fmt"
	_ "github.com/go-park-mail-ru/2020_2_Jigglypuf/docs"
	cinemaService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cinemaserver"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cookieserver"
	hallService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/hallserver"
	movieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/movieserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/recserver"
	scheduleService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/scheduleserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/ticketservice"
	authDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/delivery"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/csrf"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/logger"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/delivery"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/middleware"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"sync"
	"time"
)

type RouterStruct struct {
	CinemaService         *cinemaService.CinemaService
	MovieService          *movieService.MovieService
	CookieService         *cookieService.CookieService
	ScheduleService       *scheduleService.ScheduleService
	TicketService         *ticketservice.TicketService
	HallService           *hallService.HallService
	CsrfMiddleware        *csrf.HashCSRFToken
	AuthServiceClient     *authDelivery.UserHandler
	ProfileServiceClient  *profileDelivery.ProfileHandler
	RecommendationService *recserver.RecommendationService
}

func ConfigureHandlers(cookieDBConnection *tarantool.Connection, mainDBConnection *sql.DB, authClient authService.AuthenticationServiceClient, profileClient profileService.ProfileServiceClient) (*RouterStruct, error) {
	mutex := sync.RWMutex{}
	NewCookieService, cookieErr := cookieService.Start(cookieDBConnection)
	if cookieErr != nil {
		log.Println("No Tarantool Cookie DB connection")
		return nil, cookieErr
	}

	newHallService, hallErr := hallService.Start(mainDBConnection)
	newCinemaService, cinemaErr := cinemaService.Start(mainDBConnection)
	newMovieService, movieErr := movieService.Start(mainDBConnection, authClient)
	newScheduleService, scheduleErr := scheduleService.Start(mainDBConnection)
	if scheduleErr != nil || hallErr != nil {
		log.Println(scheduleErr, hallErr)
		return nil, models.ErrFooInitFail
	}

	newTicketService, ticketErr := ticketservice.Start(mainDBConnection, authClient, newHallService.Repository, newScheduleService.Repository)
	newHashCSRFMiddleware, csrfErr := csrf.NewHashCSRFToken(models.RandStringRunes(7), time.Hour*24)
	if cinemaErr != nil || movieErr != nil || ticketErr != nil || csrfErr != nil {
		log.Println(models.ErrFooInitFail)
		return nil, models.ErrFooInitFail
	}

	recommendationService, recErr := recserver.Start(mainDBConnection, &mutex, time.Minute*10)
	if recErr != nil {
		return nil, models.ErrFooInitFail
	}

	authHandler := authDelivery.NewUserHandler(authClient)
	profileHandler := profileDelivery.NewProfileHandler(profileClient)
	return &RouterStruct{
		AuthServiceClient:     authHandler,
		ProfileServiceClient:  profileHandler,
		CinemaService:         newCinemaService,
		MovieService:          newMovieService,
		CookieService:         NewCookieService,
		ScheduleService:       newScheduleService,
		TicketService:         newTicketService,
		HallService:           newHallService,
		CsrfMiddleware:        newHashCSRFMiddleware,
		RecommendationService: recommendationService,
	}, nil
}

func configureAuthRouter(authHandler *authDelivery.UserHandler) *httprouter.Router {
	authAPIHandler := httprouter.New()

	authAPIHandler.POST(utils.AuthURLPattern+"register/", authHandler.RegisterHandler)
	authAPIHandler.POST(utils.AuthURLPattern+"login/", authHandler.AuthHandler)
	authAPIHandler.POST(utils.AuthURLPattern+"logout/", authHandler.SignOutHandler)

	return authAPIHandler
}

func configureProfileRouter(handler *profileDelivery.ProfileHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET(utils.ProfileURLPattern, handler.GetProfile)
	router.PUT(utils.ProfileURLPattern, handler.UpdateProfile)

	return router
}


func ConfigureRouter(application *RouterStruct) http.Handler {
	handler := mux.NewRouter()
	fmt.Println(application.MovieService.MovieRouter)
	handler.Handle(utils.MovieURLPattern, application.MovieService.MovieRouter)
	handler.Handle(utils.CinemaURLPattern, application.CinemaService.CinemaRouter)
	handler.Handle(utils.AuthURLPattern, configureAuthRouter(application.AuthServiceClient))
	handler.Handle(utils.ProfileURLPattern, configureProfileRouter(application.ProfileServiceClient))
	handler.Handle(utils.ScheduleURLPattern, application.ScheduleService.Router)
	handler.Handle(utils.HallURLPattern, application.HallService.Router)
	handler.Handle(utils.TicketURLPattern, application.TicketService.Router)
	handler.Handle(utils.RecommendationsURLPattern, application.RecommendationService.RecommendationRouter)
	handler.HandleFunc(utils.CSRFURLPattern, application.CsrfMiddleware.GenerateCSRFToken)

	handler.HandleFunc(utils.MediaURLPattern, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI, http.StatusMovedPermanently)
	})
	handler.HandleFunc(utils.DocsURLPattern, httpSwagger.WrapHandler)
	middlewareHandler := application.CsrfMiddleware.CSRFMiddleware(handler)
	middlewareHandler = middleware.CookieMiddleware(middlewareHandler, application.CookieService.CookieDelivery)
	middlewareHandler = cors.MiddlewareCORS(middlewareHandler)
	middlewareHandler = logger.AccessLogMiddleware(middlewareHandler)
	return middlewareHandler
}
