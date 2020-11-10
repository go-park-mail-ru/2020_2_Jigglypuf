//go:generate mockgen -source HallRepository.go -destination mock/HallRep_mock.go -package mock
package hallservice

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type Repository interface {
	CheckAvailability(hallID uint64, place *models.TicketPlace) (bool, error)
	GetHallStructure(hallID uint64) (*models.CinemaHall, error)
}
