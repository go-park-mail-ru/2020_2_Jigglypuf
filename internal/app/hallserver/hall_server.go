package hallserver

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/globalconfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice/usecase"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/gorilla/mux"
)

type HallService struct {
	Repository hallservice.Repository
	UseCase    hallservice.UseCase
	Delivery   *delivery.HallDelivery
	Router     *mux.Router
}

func configureAPI(handler *delivery.HallDelivery) *mux.Router {
	handle := mux.NewRouter()

	handle.HandleFunc(globalconfig.HallURLPattern+fmt.Sprintf("{%s:[0-9]+}/", hallservice.HallIDPathName), handler.GetHallStructure)
	return handle
}

func Start(connection *sql.DB) (*HallService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	rep := repository.NewHallSQLRepository(connection)
	uc := usecase.NewHallUseCase(rep)
	handler := delivery.NewHallDelivery(uc)
	handle := configureAPI(handler)
	return &HallService{
		rep,
		uc,
		handler,
		handle,
	}, nil
}
