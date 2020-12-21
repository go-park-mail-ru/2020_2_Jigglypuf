package cinemaserver

import (
	"database/sql"
	cinemaConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice"
	cinemaDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/delivery"
	cinemaRepository "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/repository"
	cinemaUseCase "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/usecase"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/globalconfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
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
	cinemaAPIRouter.GET(globalconfig.CinemaURLPattern, handler.GetCinemaList)
	cinemaAPIRouter.GET(globalconfig.CinemaURLPattern+":id/", handler.GetCinema)

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
