//go:generate mockgen -source CinemaRepository.go -destination mock/CinemaRep_mock.go -package mock
package cinemaservice

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type Repository interface {
	CreateCinema(cinema *models.Cinema) error
	GetCinema(id uint64) (*models.Cinema, error)
	GetCinemaList(limit, page int) (*[]models.Cinema, error)
	UpdateCinema(cinema *models.Cinema) error
	DeleteCinema(cinema *models.Cinema) error
}
