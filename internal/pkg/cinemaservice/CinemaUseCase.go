//go:generate mockgen -source CinemaUseCase.go -destination mock/CinemaUC_mock.go -package mock
package cinemaservice

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type UseCase interface {
	CreateCinema(*models.Cinema) error
	GetCinema(id uint64) (*models.Cinema, error)
	GetCinemaList(limit, page int) (*[]models.Cinema, error)
}
