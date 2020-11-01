package usecase

import (
	"backend/internal/pkg/hallService"
	"backend/internal/pkg/models"
	"strconv"
)

type HallUseCase struct{
	repository hallService.Repository
}

func NewHallUseCase(repository hallService.Repository) *HallUseCase{
	return &HallUseCase{
		repository,
	}
}

func (t *HallUseCase) CheckAvailability(hallID string, place *models.TicketPlace) (bool,error){
	castedHallID, castErr := strconv.Atoi(hallID)
	if castErr != nil || hallID == "" || place == nil{
		return false,models.ErrFooIncorrectInputInfo
	}

	return t.repository.CheckAvailability(uint64(castedHallID), place)
}

func (t *HallUseCase) GetHallStructure(HallID string)(*models.CinemaHall, error){
	castedHallID, castErr := strconv.Atoi(HallID)
	if castErr != nil || HallID == ""{
		return nil,models.ErrFooIncorrectInputInfo
	}

	return t.repository.GetHallStructure(uint64(castedHallID))
}