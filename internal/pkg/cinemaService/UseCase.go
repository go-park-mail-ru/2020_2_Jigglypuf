package cinemaService

import "backend/internal/pkg/models"

type CinemaUseCase interface{
	CreateCinema(*models.Cinema) error
	GetCinema(name *string)(*models.Cinema, error)
	GetCinemaList(limit, page int)(*[]models.Cinema, error)
}