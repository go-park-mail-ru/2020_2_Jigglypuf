package router

import (
	"database/sql"
	"fmt"
	cinemaService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cinemaserver"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cookieserver"
	hallService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/hallserver"
	movieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/movieserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/recserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/replyserver"
	scheduleService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/scheduleserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/ticketservice"
	authDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/delivery"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/globalConfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/csrf"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/monitoring"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/delivery"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/middleware"
	ticketservice2 "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"sync"
	"time"
)

type RoutingConfig struct {
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
	ReplyService          *replyserver.ReplyService
}

func ConfigureHandlers(cookieDBConnection *tarantool.Connection, mainDBConnection *sql.DB, authClient authService.AuthenticationServiceClient, profileClient profileService.ProfileServiceClient) (*RoutingConfig, error) {
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

	replyService, replyErr := replyserver.Start(profileClient, mainDBConnection)
	if replyErr != nil {
		return nil, models.ErrFooInitFail
	}

	authHandler := authDelivery.NewUserHandler(authClient)
	profileHandler := profileDelivery.NewProfileHandler(profileClient)
	return &RoutingConfig{
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
		ReplyService:          replyService,
	}, nil
}

func configureAuthRouter(authHandler *authDelivery.UserHandler) *httprouter.Router {
	authAPIHandler := httprouter.New()

	authAPIHandler.POST(globalConfig.AuthURLPattern+"register/", authHandler.RegisterHandler)
	authAPIHandler.POST(globalConfig.AuthURLPattern+"login/", authHandler.AuthHandler)
	authAPIHandler.POST(globalConfig.AuthURLPattern+"logout/", authHandler.SignOutHandler)

	return authAPIHandler
}

func configureProfileRouter(handler *profileDelivery.ProfileHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET(globalConfig.ProfileURLPattern, handler.GetProfile)
	router.PUT(globalConfig.ProfileURLPattern, handler.UpdateProfile)

	return router
}

func ConfigureRouter(application *RoutingConfig) http.Handler {
	handler := mux.NewRouter()

	handler.Handle(globalConfig.MovieURLPattern, application.MovieService.MovieRouter)
	handler.Handle(globalConfig.CinemaURLPattern, application.CinemaService.CinemaRouter)
	handler.Handle(globalConfig.AuthURLPattern, configureAuthRouter(application.AuthServiceClient))
	handler.Handle(globalConfig.ProfileURLPattern, configureProfileRouter(application.ProfileServiceClient))
	handler.Handle(globalConfig.ScheduleURLPattern, application.ScheduleService.Router)
	handler.Handle(globalConfig.HallURLPattern, application.HallService.Router)
	handler.Handle(globalConfig.TicketURLPattern, application.TicketService.Router)
	handler.Handle(globalConfig.RecommendationsURLPattern, application.RecommendationService.RecommendationRouter)
	handler.Handle(globalConfig.ReplyURLPattern, application.ReplyService.ReplyRouter)
	handler.HandleFunc(globalConfig.QRCodeTicketURLPattern+fmt.Sprintf("{%s:[0-9A-Za-z]+}/", ticketservice2.TicketTransactionPathName),
		application.TicketService.Handler.GetTicketByCode)
	handler.HandleFunc(globalConfig.CSRFURLPattern, application.CsrfMiddleware.GenerateCSRFToken)

	handler.HandleFunc(globalConfig.MediaURLPattern, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI, http.StatusMovedPermanently)
	})

	handler.HandleFunc(globalConfig.DocsURLPattern, httpSwagger.WrapHandler)
	handler.Handle("/metrics/", promhttp.Handler())

	middlewareHandler := application.CsrfMiddleware.CSRFMiddleware(handler)
	middlewareHandler = middleware.CookieMiddleware(middlewareHandler, application.CookieService.CookieDelivery)
	middlewareHandler = cors.MiddlewareCORS(middlewareHandler)
	middlewareHandler = monitoring.AccessLogMiddleware(middlewareHandler)
	return middlewareHandler
}
