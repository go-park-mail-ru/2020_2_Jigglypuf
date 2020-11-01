package ticketService

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/hallService"
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"backend/internal/pkg/ticketService"
	"backend/internal/pkg/ticketService/delivery"
	"backend/internal/pkg/ticketService/repository"
	"backend/internal/pkg/ticketService/usecase"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
)

type TicketService struct{
	Repository ticketService.Repository
	UseCase ticketService.UseCase
	Handler *delivery.TicketDelivery
	Router *mux.Router
}

func configureAPI(handler *delivery.TicketDelivery) *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc(ticketService.URLPattern + "buy/", handler.BuyTicket).Methods("POST")
	router.HandleFunc(ticketService.URLPattern, handler.GetUserTickets).Methods("GET")
	router.HandleFunc(ticketService.URLPattern + fmt.Sprintf("{%s:[0-9]+}/", ticketService.TicketIDQuery),
		handler.GetUsersSimpleTicket).Methods("GET")
	router.HandleFunc("/ticket" + schedule.URLPattern + fmt.Sprintf("{%s:[0-9]+}/", ticketService.ScheduleIDName),
		handler.GetHallScheduleTickets).Methods("GET")

	return router
}

func Start(connection *sql.DB, authRep authentication.AuthRepository, hallRep hallService.Repository) (*TicketService,error){
	if connection == nil || authRep == nil || hallRep == nil{
		return nil, models.ErrFooArgsMismatch
	}
	rep := repository.NewTicketSQLRepository(connection)
	uc := usecase.NewTicketUseCase(rep,authRep, hallRep)
	handler := delivery.NewTicketDelivery(uc)
	router := configureAPI(handler)

	return &TicketService{
		rep,
		uc,
		handler,
		router,
	}, nil
}
