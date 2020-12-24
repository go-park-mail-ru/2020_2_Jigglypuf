package ticketservice

import (
	"database/sql"
	"fmt"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/globalconfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice/usecase"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils/configurator"
	"github.com/gorilla/mux"
)

type TicketService struct {
	Repository ticketservice.Repository
	UseCase    ticketservice.UseCase
	Handler    *delivery.TicketDelivery
	Router     *mux.Router
}

func configureAPI(handler *delivery.TicketDelivery) *mux.Router {
	handle := mux.NewRouter()

	handle.HandleFunc(globalconfig.TicketURLPattern+"buy/", handler.BuyTicket).Methods("POST")
	handle.HandleFunc(globalconfig.TicketURLPattern, handler.GetUserTickets).Methods("GET")
	handle.HandleFunc(globalconfig.TicketURLPattern+fmt.Sprintf("{%s:[0-9]+}/", ticketservice.TicketIDQuery),
		handler.GetUsersSimpleTicket).Methods("GET")
	handle.HandleFunc(globalconfig.TicketScheduleURLPattern+fmt.Sprintf("{%s:[0-9]+}/", ticketservice.ScheduleIDName),
		handler.GetHallScheduleTickets).Methods("GET")
	handle.HandleFunc(globalconfig.QRCodeTicketURLPattern+fmt.Sprintf("{%s:[0-9A-Za-z]+}/", ticketservice.TicketTransactionPathName),
		handler.GetTicketByCode)

	return handle
}

func Start(connection *sql.DB, auth authService.AuthenticationServiceClient, hallRep hallservice.Repository, scheduleRep schedule.TimeTableRepository, config *configurator.MailConfig) (*TicketService, error) {
	if connection == nil || auth == nil || hallRep == nil {
		return nil, models.ErrFooArgsMismatch
	}
	rep := repository.NewTicketSQLRepository(connection)
	uc := usecase.NewTicketUseCase(rep, auth, hallRep, scheduleRep, config.Mail, config.Password, config.SMTPServerDomain, config.SMTPServerPort)
	handler := delivery.NewTicketDelivery(uc)
	handle := configureAPI(handler)

	return &TicketService{
		rep,
		uc,
		handler,
		handle,
	}, nil
}
