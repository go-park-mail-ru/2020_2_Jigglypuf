package ticketservice

import (
	"backend/internal/pkg/models"
)

type Repository interface {
	CreateTicket(ticket *models.TicketInput) error
	GetSimpleTicket(ticketID uint64, Login string) (*models.Ticket, error)
	GetUserTickets(Login string) (*[]models.Ticket, error)
	GetHallTickets(scheduleID uint64) (*[]models.TicketPlace, error)
}
