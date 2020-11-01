package hallServer

import (
	"backend/internal/pkg/hallService"
	"backend/internal/pkg/hallService/delivery"
	"backend/internal/pkg/hallService/repository"
	"backend/internal/pkg/hallService/usecase"
	"backend/internal/pkg/models"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
)

type HallService struct{
	Repository hallService.Repository
	UseCase hallService.UseCase
	Delivery *delivery.HallDelivery
	Router *mux.Router
}

func configureAPI(handler *delivery.HallDelivery)*mux.Router{
	router := mux.NewRouter()

	router.HandleFunc(hallService.URLPattern + fmt.Sprintf("{%s:[0-9]+}/",hallService.HallIDPathName), handler.GetHallStructure)
	return router
}

func Start(connection *sql.DB) (*HallService,error){
	if connection == nil{
		return nil,models.ErrFooNoDBConnection
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