package ticketservice

import (
	"backend/internal/pkg/authentication/interfaces"
	"backend/internal/pkg/hallservice"
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"backend/internal/pkg/ticketservice"
	"backend/internal/pkg/ticketservice/delivery"
	"backend/internal/pkg/ticketservice/repository"
	"backend/internal/pkg/ticketservice/usecase"
	"database/sql"
	"fmt"
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
	router.HandleFunc("/ticket"+schedule.URLPattern+fmt.Sprintf("{%s:[0-9]+}/", ticketservice.ScheduleIDName),
		handler.GetHallScheduleTickets).Methods("GET")

	return router
}

func Start(connection *sql.DB, authRep interfaces.AuthRepository, hallRep hallservice.Repository) (*TicketService, error) {
	if connection == nil || authRep == nil || hallRep == nil {
		return nil, models.ErrFooArgsMismatch
	}
	rep := repository.NewTicketSQLRepository(connection)
	uc := usecase.NewTicketUseCase(rep, authRep, hallRep)
	handler := delivery.NewTicketDelivery(uc)
	router := configureAPI(handler)

	return &TicketService{
		rep,
		uc,
		handler,
		router,
	}, nil
}
