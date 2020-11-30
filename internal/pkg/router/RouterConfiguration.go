package router

import (
	"database/sql"
	cinemaService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cinemaserver"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/cookieserver"
	hallService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/hallserver"
	movieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/movieserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/recserver"
	scheduleService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/scheduleserver"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/ticketservice"
	authConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	authDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/delivery"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/csrf"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/logger"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	profileDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/delivery"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/middleware"
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
	cinemaService         *cinemaService.CinemaService
	movieService          *movieService.MovieService
	cookieService         *cookieService.CookieService
	scheduleService       *scheduleService.ScheduleService
	ticketService         *ticketservice.TicketService
	hallService           *hallService.HallService
	csrfMiddleware        *csrf.HashCSRFToken
	authServiceClient     *authDelivery.UserHandler
	profileServiceClient  *profileDelivery.ProfileHandler
	recommendationService *recserver.RecommendationService
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
		authServiceClient:     authHandler,
		profileServiceClient:  profileHandler,
		cinemaService:         newCinemaService,
		movieService:          newMovieService,
		cookieService:         NewCookieService,
		scheduleService:       newScheduleService,
		ticketService:         newTicketService,
		hallService:           newHallService,
		csrfMiddleware:        newHashCSRFMiddleware,
		recommendationService: recommendationService,
	}, nil
}

func configureAuthRouter(authHandler *authDelivery.UserHandler) *httprouter.Router {
	authAPIHandler := httprouter.New()

	authAPIHandler.POST(authConfig.URLPattern+"register/", authHandler.RegisterHandler)
	authAPIHandler.POST(authConfig.URLPattern+"login/", authHandler.AuthHandler)
	authAPIHandler.POST(authConfig.URLPattern+"logout/", authHandler.SignOutHandler)

	return authAPIHandler
}

func configureProfileRouter(handler *profileDelivery.ProfileHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET(profileConfig.URLPattern, handler.GetProfile)
	router.PUT(profileConfig.URLPattern, handler.UpdateProfile)

	return router
}


func ConfigureRouter(application *RouterStruct) http.Handler {
	handler := mux.NewRouter()
	handler = handler.PathPrefix("/api").Subrouter()

	handler.Handle(MovieURLPattern, application.movieService.MovieRouter)
	handler.Handle(CinemaURLPattern, application.cinemaService.CinemaRouter)
	handler.Handle(AuthURLPattern, configureAuthRouter(application.authServiceClient))
	handler.Handle(ProfileURLPattern, configureProfileRouter(application.profileServiceClient))
	handler.Handle(ScheduleURLPattern, application.scheduleService.Router)
	handler.Handle(HallURLPattern, application.hallService.Router)
	handler.Handle(TicketURLPattern, application.ticketService.Router)
	handler.Handle(RecommendationsURLPattern, application.recommendationService.RecommendationRouter)
	handler.HandleFunc(CSRFURLPattern, application.csrfMiddleware.GenerateCSRFToken)

	handler.HandleFunc(MediaURLPattern, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI, http.StatusMovedPermanently)
	})
	handler.HandleFunc(DocsURLPattern, httpSwagger.WrapHandler)
	middlewareHandler := application.csrfMiddleware.CSRFMiddleware(handler)
	middlewareHandler = middleware.CookieMiddleware(middlewareHandler, application.cookieService.CookieDelivery)
	middlewareHandler = cors.MiddlewareCORS(middlewareHandler)
	middlewareHandler = logger.AccessLogMiddleware(middlewareHandler)
	return middlewareHandler
}
