package hallService

import "backend/internal/pkg/models"

type UseCase interface{
	CheckAvailability(hallID string, place *models.TicketPlace) (bool,error)
	GetHallStructure(HallID string)(*models.CinemaHall, error)
}
