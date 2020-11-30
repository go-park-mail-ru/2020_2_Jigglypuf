package ticketservice

import (
	"database/sql"
	"fmt"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice/usecase"
	"github.com/gorilla/mux"
)

type TicketService struct {
	Repository ticketservice.Repository
	UseCase    ticketservice.UseCase
	Handler    *delivery.TicketDelivery
	Router     *mux.Router
}

func configureAPI(handler *delivery.TicketDelivery) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(ticketservice.URLPattern+"buy/", handler.BuyTicket).Methods("POST")
	router.HandleFunc(ticketservice.URLPattern, handler.GetUserTickets).Methods("GET")
	router.HandleFunc(ticketservice.URLPattern+fmt.Sprintf("{%s:[0-9]+}/", ticketservice.TicketIDQuery),
		handler.GetUsersSimpleTicket).Methods("GET")
	router.HandleFunc(ticketservice.ScheduleURLPattern+fmt.Sprintf("{%s:[0-9]+}/", ticketservice.ScheduleIDName),
		handler.GetHallScheduleTickets).Methods("GET")

	return router
}

func Start(connection *sql.DB, auth authService.AuthenticationServiceClient, hallRep hallservice.Repository, scheduleRep schedule.TimeTableRepository) (*TicketService, error) {
	if connection == nil || auth == nil || hallRep == nil {
		return nil, models.ErrFooArgsMismatch
	}
	rep := repository.NewTicketSQLRepository(connection)
	uc := usecase.NewTicketUseCase(rep, auth, hallRep, scheduleRep)
	handler := delivery.NewTicketDelivery(uc)
	router := configureAPI(handler)

	return &TicketService{
		rep,
		uc,
		handler,
		router,
	}, nil
}
