//go:generate mockgen -source HallUsecase.go -destination mock/HallUC_mock.go -package mock
package hallservice

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type UseCase interface {
	CheckAvailability(hallID string, place *models.TicketPlace) (bool, error)
	GetHallStructure(hallID string) (*models.CinemaHall, error)
}
