package cinema_server

import(
	cinemaRepository "backend/internal/pkg/cinemaservice/repository"
	cinemaUseCase "backend/internal/pkg/cinemaservice/usecase"
	cinemaDelivery "backend/internal/pkg/cinemaservice/delivery"
	cinemaConfig "backend/internal/pkg/cinemaservice"
	"github.com/julienschmidt/httprouter"
	"sync"
)

type CinemaService struct{
	CinemaRepository *cinemaRepository.CinemaRepository
	CinemaUseCase *cinemaUseCase.CinemaUseCase
	CinemaDelivery *cinemaDelivery.CinemaHandler
	CinemaRouter *httprouter.Router
}


func configureCinemaRouter(handler *cinemaDelivery.CinemaHandler) *httprouter.Router{
	cinemaApiRouter := httprouter.New()
	cinemaApiRouter.GET(cinemaConfig.UrlPattern, handler.GetCinemaList)
	cinemaApiRouter.GET(cinemaConfig.UrlPattern + ":id/", handler.GetCinema)

	return cinemaApiRouter
}

func Start(mutex *sync.RWMutex) *CinemaService{
	cinemaRep := cinemaRepository.NewCinemaRepository(mutex)
	cinemaUC := cinemaUseCase.NewCinemaUseCase(cinemaRep)
	cinemaHandler := cinemaDelivery.NewCinemaHandler(cinemaUC)

	cinemaRouter := configureCinemaRouter(cinemaHandler)

	return &CinemaService{
		cinemaRep,
		cinemaUC,
		cinemaHandler,
		cinemaRouter,
	}
}