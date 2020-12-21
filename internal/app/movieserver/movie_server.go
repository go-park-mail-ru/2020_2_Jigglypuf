package movieserver

import (
	"database/sql"
	"fmt"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/globalConfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	movieConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	movieDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/delivery"
	movieRepository "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/repository"
	movieUseCase "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/usecase"
	"github.com/gorilla/mux"
)

type MovieService struct {
	MovieRepository movieConfig.MovieRepository
	MovieUseCase    *movieUseCase.MovieUseCase
	MovieDelivery   *movieDelivery.MovieHandler
	MovieRouter     *mux.Router
}

func configureMovieRouter(handler *movieDelivery.MovieHandler) *mux.Router {
	movieRouter := mux.NewRouter()
	movieRouter.HandleFunc(globalConfig.MovieURLPattern, handler.GetMovieList)
	movieRouter.HandleFunc(globalConfig.MovieURLPattern+"rate/", handler.RateMovie)
	movieRouter.HandleFunc(globalConfig.MovieURLPattern+"actual/", handler.GetActualMovies)
	movieRouter.HandleFunc(globalConfig.MovieURLPattern+fmt.Sprintf("{%s:[0-9]+}/", movieConfig.MovieIDQuery), handler.GetMovie)

	return movieRouter
}

func Start(connection *sql.DB, auth authService.AuthenticationServiceClient) (*MovieService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	movieRep := movieRepository.NewMovieSQLRepository(connection)
	movieUC := movieUseCase.NewMovieUseCase(movieRep, auth)
	movieHandler := movieDelivery.NewMovieHandler(movieUC)

	movieRouter := configureMovieRouter(movieHandler)

	return &MovieService{
		movieRep,
		movieUC,
		movieHandler,
		movieRouter,
	}, nil
}
