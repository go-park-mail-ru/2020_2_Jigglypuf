package cinemaserver

import (
	cinemaConfig "backend/internal/pkg/cinemaservice"
	cinemaDelivery "backend/internal/pkg/cinemaservice/delivery"
	cinemaRepository "backend/internal/pkg/cinemaservice/repository"
	cinemaUseCase "backend/internal/pkg/cinemaservice/usecase"
	"backend/internal/pkg/models"
	"database/sql"
	"github.com/julienschmidt/httprouter"
)

type CinemaService struct {
	CinemaRepository cinemaConfig.Repository
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

func Start(connection *sql.DB) (*CinemaService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	cinemaRep := cinemaRepository.NewCinemaSQLRepository(connection)
	cinemaUC := cinemaUseCase.NewCinemaUseCase(cinemaRep)
	cinemaHandler := cinemaDelivery.NewCinemaHandler(cinemaUC)

	cinemaRouter := configureCinemaRouter(cinemaHandler)

	return &CinemaService{
		cinemaRep,
		cinemaUC,
		cinemaHandler,
		cinemaRouter,
	}, nil
}
