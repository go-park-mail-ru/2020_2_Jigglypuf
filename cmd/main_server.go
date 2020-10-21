package main

import (
	authService "backend/internal/app/auth_server"
	cinemaService "backend/internal/app/cinema_server"
	cookieService "backend/internal/app/cookie_server"
	movieService "backend/internal/app/movie_server"
	profileService "backend/internal/app/profile_server"
	authConfig "backend/internal/pkg/authentication"
	cinemaConfig "backend/internal/pkg/cinemaservice"
	"backend/internal/pkg/middleware/CORS"
	"backend/internal/pkg/middleware/cookie/middleware"
	movieConfig "backend/internal/pkg/movieservice"
	profileConfig "backend/internal/pkg/profile"
	"log"
	"net/http"
	"sync"
	"time"
)

type ServerStruct struct {
	authService    *authService.AuthService
	cinemaService  *cinemaService.CinemaService
	movieService   *movieService.MovieService
	profileService *profileService.ProfileService
	cookieService  *cookieService.CookieService
	httpServer     *http.Server
}


func configureAPI() *ServerStruct {
	mutex := sync.RWMutex{}
	NewCookieService := cookieService.Start(&mutex)
	newAuthService := authService.Start(&mutex, NewCookieService.CookieRepository)
	newCinemaService := cinemaService.Start(&mutex)
	newMovieService := movieService.Start(&mutex, newAuthService.AuthenticationRepository)
	newProfileService := profileService.Start(&mutex,newAuthService.AuthenticationRepository)

	return &ServerStruct{
		authService:    newAuthService,
		cinemaService:  newCinemaService,
		movieService:   newMovieService,
		profileService: newProfileService,
		cookieService: NewCookieService,
	}
}

func configureRouter(application *ServerStruct) http.Handler {
	handler := http.NewServeMux()

	handler.Handle(movieConfig.UrlPattern, application.movieService.MovieRouter)
	handler.Handle(cinemaConfig.UrlPattern, application.cinemaService.CinemaRouter)
	handler.Handle(authConfig.UrlPattern, application.authService.AuthRouter)
	handler.Handle(profileConfig.UrlPattern, application.profileService.ProfileRouter)



	handler.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI, http.StatusMovedPermanently)
	})
	middlewareHandler := middleware.CookieMiddleware(handler)
	middlewareHandler = CORS.MiddlewareCORS(middlewareHandler)

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

func main() {
	serverConfig := configureAPI()
	responseHandler := configureRouter(serverConfig)
	serverConfig.httpServer = configureServer("8080", responseHandler)

	err := serverConfig.httpServer.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
