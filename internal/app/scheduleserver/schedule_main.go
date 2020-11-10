package scheduleserver

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule/usecase"
	"database/sql"
	"github.com/gorilla/mux"
)

type ScheduleService struct {
	Delivery   *delivery.ScheduleDelivery
	UseCase    schedule.TimeTableUseCase
	Repository schedule.TimeTableRepository
	Router     *mux.Router
}

func configureRouter(handler *delivery.ScheduleDelivery) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(schedule.URLPattern, handler.GetMovieSchedule)
	return router
}

func Start(connection *sql.DB) (*ScheduleService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	rep := repository.NewScheduleSQLRepository(connection)
	useCase := usecase.NewTimeTableUseCase(rep)
	handler := delivery.NewScheduleDelivery(useCase)
	router := configureRouter(handler)

	return &ScheduleService{
		handler,
		useCase,
		rep,
		router,
	}, nil
}
