//go:generate mockgen -source repository.go -destination mock/CinemaRep_mock.go -package mock
package cinemaservice

import (
	"backend/internal/pkg/models"
)

type Repository interface {
	CreateCinema(cinema *models.Cinema) error
	GetCinema(id uint64) (*models.Cinema, error)
	GetCinemaList(limit, page int) (*[]models.Cinema, error)
	UpdateCinema(cinema *models.Cinema) error
	DeleteCinema(cinema *models.Cinema) error
}
