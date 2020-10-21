package movie_server

import(
	"backend/internal/pkg/authentication"
	movieRepository "backend/internal/pkg/movieservice/repository"
	movieUseCase "backend/internal/pkg/movieservice/usecase"
	movieDelivery "backend/internal/pkg/movieservice/delivery"
	movieConfig "backend/internal/pkg/movieservice"
	"github.com/julienschmidt/httprouter"
	"sync"
)

type MovieService struct{
	MovieRepository *movieRepository.MovieRepository
	MovieUseCase *movieUseCase.MovieUseCase
	MovieDelivery *movieDelivery.MovieHandler
	MovieRouter *httprouter.Router
}

func configureMovieRouter(handler *movieDelivery.MovieHandler) *httprouter.Router{
	movieRouter := httprouter.New()

	movieRouter.GET(movieConfig.UrlPattern + ":id/", handler.GetMovie)
	movieRouter.GET(movieConfig.UrlPattern, handler.GetMovieList)
	movieRouter.POST(movieConfig.UrlPattern + ":id/rate/", handler.RateMovie)

	return movieRouter
}


func Start(mutex *sync.RWMutex, authRep authentication.AuthRepository) *MovieService{
	movieRep := movieRepository.NewMovieRepository(mutex)
	movieUC := movieUseCase.NewMovieUseCase(movieRep)
	movieHandler := movieDelivery.NewMovieHandler(movieUC,authRep)

	movieRouter := configureMovieRouter(movieHandler)

	return &MovieService{
		movieRep,
		movieUC,
		movieHandler,
		movieRouter,
	}
}
