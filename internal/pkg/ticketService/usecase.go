package ticketService

import "backend/internal/pkg/models"

type UseCase interface{
	BuyTicket(ticket *models.Ticket, userID uint64) error
	GetHallScheduleTickets(scheduleID string)(*[]models.TicketPlace,error)
	GetSimpleTicket(userID uint64, ticketID string)(*models.Ticket, error)
	GetUserTickets(userID uint64) (*[]models.Ticket, error)
}