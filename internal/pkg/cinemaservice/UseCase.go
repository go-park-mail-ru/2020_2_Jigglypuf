//go:generate mockgen -source UseCase.go -destination mock/CinemaUC_mock.go -package mock
package cinemaservice

import "backend/internal/pkg/models"

type UseCase interface {
	CreateCinema(*models.Cinema) error
	GetCinema(id uint64) (*models.Cinema, error)
	GetCinemaList(limit, page int) (*[]models.Cinema, error)
}
