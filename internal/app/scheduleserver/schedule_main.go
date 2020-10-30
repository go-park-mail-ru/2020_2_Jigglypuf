package scheduleserver

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"backend/internal/pkg/schedule/delivery"
	"backend/internal/pkg/schedule/repostiory"
	"backend/internal/pkg/schedule/usecase"
	"database/sql"
	"github.com/gorilla/mux"
)

type ScheduleService struct{
	Delivery *delivery.ScheduleDelivery
	UseCase  schedule.TimeTableUseCase
	Repository schedule.TimeTableRepository
	Router *mux.Router
}

func configureRouter(handler *delivery.ScheduleDelivery) *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc(schedule.URLPattern, handler.GetMovieSchedule)
	return router
}

func Start(connection *sql.DB) (*ScheduleService,error){
	if connection == nil{
		return nil,models.ErrFooNoDBConnection
	}
	repository := repostiory.NewScheduleSQLRepository(connection)
	useCase := usecase.NewTimeTableUseCase(repository)
	handler := delivery.NewScheduleDelivery(useCase)
	router := configureRouter(handler)

	return &ScheduleService{
		handler,
		useCase,
		repository,
		router,
	}, nil
}