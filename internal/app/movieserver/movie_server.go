package movieserver

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/models"
	movieConfig "backend/internal/pkg/movieservice"
	movieDelivery "backend/internal/pkg/movieservice/delivery"
	movieRepository "backend/internal/pkg/movieservice/repository"
	movieUseCase "backend/internal/pkg/movieservice/usecase"
	"github.com/gorilla/mux"
	"database/sql"
)

type MovieService struct {
	MovieRepository movieConfig.MovieRepository
	MovieUseCase    *movieUseCase.MovieUseCase
	MovieDelivery   *movieDelivery.MovieHandler
	MovieRouter     *mux.Router
}

func configureMovieRouter(handler *movieDelivery.MovieHandler) *mux.Router {
	movieRouter := mux.NewRouter()
	movieRouter.HandleFunc(movieConfig.URLPattern, handler.GetMovieList)
	movieRouter.HandleFunc(movieConfig.URLPattern+"rate/", handler.RateMovie)
	movieRouter.HandleFunc(movieConfig.URLPattern+"actual/", handler.GetMoviesInCinema)
	movieRouter.HandleFunc(movieConfig.URLPattern+"{id:[0-9]+}/", handler.GetMovie)

	return movieRouter
}

func Start(connection *sql.DB, authRep authentication.AuthRepository) (*MovieService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	movieRep := movieRepository.NewMovieSQLRepository(connection)
	movieUC := movieUseCase.NewMovieUseCase(movieRep)
	movieHandler := movieDelivery.NewMovieHandler(movieUC, authRep)

	movieRouter := configureMovieRouter(movieHandler)

	return &MovieService{
		movieRep,
		movieUC,
		movieHandler,
		movieRouter,
	}, nil
}
