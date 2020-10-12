package cinemaservice

import "backend/internal/pkg/models"

type UseCase interface {
	CreateCinema(*models.Cinema) error
	GetCinema(name *string) (*models.Cinema, error)
	GetCinemaList(limit, page int) (*[]models.Cinema, error)
}
