//go:generate mockgen -source TicketRepository.go -destination mock/TicketRep_mock.go -package mock
package ticketservice

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type Repository interface {
	CreateTicket(ticket *models.TicketInput) error
	GetSimpleTicket(ticketID uint64, Login string) (*models.Ticket, error)
	GetUserTickets(Login string) (*[]models.Ticket, error)
	GetHallTickets(scheduleID uint64) (*[]models.TicketPlace, error)
	GetTicketByTransaction(transaction string) (*models.TicketInfo, error)
}
