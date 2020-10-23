package cinemaserver

import (
	cinemaConfig "backend/internal/pkg/cinemaservice"
	cinemaDelivery "backend/internal/pkg/cinemaservice/delivery"
	cinemaRepository "backend/internal/pkg/cinemaservice/repository"
	cinemaUseCase "backend/internal/pkg/cinemaservice/usecase"
	"github.com/julienschmidt/httprouter"
	"sync"
)

type CinemaService struct {
	CinemaRepository *cinemaRepository.CinemaRepository
	CinemaUseCase    *cinemaUseCase.CinemaUseCase
	CinemaDelivery   *cinemaDelivery.CinemaHandler
	CinemaRouter     *httprouter.Router
}

func configureCinemaRouter(handler *cinemaDelivery.CinemaHandler) *httprouter.Router {
	cinemaAPIRouter := httprouter.New()
	cinemaAPIRouter.GET(cinemaConfig.URLPattern, handler.GetCinemaList)
	cinemaAPIRouter.GET(cinemaConfig.URLPattern+":id/", handler.GetCinema)

	return cinemaAPIRouter
}

func Start(mutex *sync.RWMutex) *CinemaService {
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
