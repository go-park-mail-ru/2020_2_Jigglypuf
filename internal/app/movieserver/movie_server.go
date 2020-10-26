package movieserver

import (
	"backend/internal/pkg/authentication"
	movieConfig "backend/internal/pkg/movieservice"
	movieDelivery "backend/internal/pkg/movieservice/delivery"
	movieRepository "backend/internal/pkg/movieservice/repository"
	movieUseCase "backend/internal/pkg/movieservice/usecase"
	"database/sql"
	"errors"
	"github.com/julienschmidt/httprouter"
)

type MovieService struct {
	MovieRepository movieConfig.MovieRepository
	MovieUseCase    *movieUseCase.MovieUseCase
	MovieDelivery   *movieDelivery.MovieHandler
	MovieRouter     *httprouter.Router
}

func configureMovieRouter(handler *movieDelivery.MovieHandler) *httprouter.Router {
	movieRouter := httprouter.New()

	movieRouter.GET(movieConfig.URLPattern+":id/", handler.GetMovie)
	movieRouter.GET(movieConfig.URLPattern, handler.GetMovieList)
	movieRouter.POST(movieConfig.URLPattern+":id/rate/", handler.RateMovie)

	return movieRouter
}

func Start(connection *sql.DB, authRep authentication.AuthRepository) (*MovieService, error) {
	if connection == nil {
		return nil, errors.New("no database connection")
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
