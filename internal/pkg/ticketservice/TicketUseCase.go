//go:generate mockgen -source TicketUseCase.go -destination mock/ticketUC_mock.go -package mock
package ticketservice

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type UseCase interface {
	BuyTicket(ticket *models.TicketInput, userID interface{}) error
	GetHallScheduleTickets(scheduleID string) (*[]models.TicketPlace, error)
	GetSimpleTicket(userID uint64, ticketID string) (*models.Ticket, error)
	GetUserTickets(userID uint64) (*[]models.Ticket, error)
}
