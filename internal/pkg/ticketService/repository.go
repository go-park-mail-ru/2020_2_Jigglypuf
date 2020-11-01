package ticketService

import (
	"backend/internal/pkg/models"
)

type Repository interface{
	CreateTicket(ticket *models.Ticket) error
	GetSimpleTicket(ticketID uint64, username string)(*models.Ticket,error)
	GetUserTickets(username string)(*[]models.Ticket, error)
	GetHallTickets(scheduleID uint64)(*[]models.TicketPlace, error)
}