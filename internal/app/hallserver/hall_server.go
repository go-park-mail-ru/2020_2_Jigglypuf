package hallserver

import (
	"backend/internal/pkg/hallservice"
	"backend/internal/pkg/hallservice/delivery"
	"backend/internal/pkg/hallservice/repository"
	"backend/internal/pkg/hallservice/usecase"
	"backend/internal/pkg/models"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
)

type HallService struct {
	Repository hallservice.Repository
	UseCase    hallservice.UseCase
	Delivery   *delivery.HallDelivery
	Router     *mux.Router
}

func configureAPI(handler *delivery.HallDelivery) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(hallservice.URLPattern+fmt.Sprintf("{%s:[0-9]+}/", hallservice.HallIDPathName), handler.GetHallStructure)
	return router
}

func Start(connection *sql.DB) (*HallService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	rep := repository.NewHallSQLRepository(connection)
	uc := usecase.NewHallUseCase(rep)
	handler := delivery.NewHallDelivery(uc)
	router := configureAPI(handler)
	return &HallService{
		rep,
		uc,
		handler,
		router,
	}, nil
}
