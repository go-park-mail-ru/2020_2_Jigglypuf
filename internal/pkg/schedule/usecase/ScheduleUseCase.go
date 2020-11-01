package usecase

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

type ScheduleUseCase struct{
	validator *validator.Validate
	ScheduleRepository schedule.TimeTableRepository
}

func NewTimeTableUseCase(repository schedule.TimeTableRepository) *ScheduleUseCase{
	return &ScheduleUseCase{
		validator.New(),
		repository,
	}
}

func (t *ScheduleUseCase) GetMovieSchedule(MovieID, CinemaID string, date string)(*[]models.Schedule, error){
	castedMovieID,castErr := strconv.Atoi(MovieID)
	if castErr != nil{
		return nil,models.ErrFooCastErr
	}
	_, castErr = time.Parse(schedule.TimeStandard,date)
	if castErr != nil{
		date = time.Now().String()
	}
	castedCinemaID, castErr := strconv.Atoi(CinemaID)
	if castErr != nil{
		return t.ScheduleRepository.GetMovieSchedule(uint64(castedMovieID),date)
	}
	return t.ScheduleRepository.GetMovieCinemaSchedule(uint64(castedMovieID),uint64(castedCinemaID),date)
}