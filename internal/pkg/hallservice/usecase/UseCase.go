package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"strconv"
)

type HallUseCase struct {
	repository hallservice.Repository
}

func NewHallUseCase(repository hallservice.Repository) *HallUseCase {
	return &HallUseCase{
		repository,
	}
}

func (t *HallUseCase) CheckAvailability(hallID string, place *models.TicketPlace) (bool, error) {
	castedHallID, castErr := strconv.Atoi(hallID)
	if castErr != nil || hallID == "" || place == nil {
		return false, models.ErrFooIncorrectInputInfo
	}

	return t.repository.CheckAvailability(uint64(castedHallID), place)
}

func (t *HallUseCase) GetHallStructure(hallID string) (*models.CinemaHall, error) {
	castedHallID, castErr := strconv.Atoi(hallID)
	if castErr != nil || hallID == "" {
		return nil, models.ErrFooIncorrectInputInfo
	}

	return t.repository.GetHallStructure(uint64(castedHallID))
}
