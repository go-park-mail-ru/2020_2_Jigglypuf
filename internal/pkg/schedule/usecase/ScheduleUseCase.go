package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

type ScheduleUseCase struct {
	validator          *validator.Validate
	ScheduleRepository schedule.TimeTableRepository
}

func NewTimeTableUseCase(repository schedule.TimeTableRepository) *ScheduleUseCase {
	return &ScheduleUseCase{
		validator.New(),
		repository,
	}
}

func (t *ScheduleUseCase) GetMovieSchedule(movieID, cinemaID string, date string) (*[]models.Schedule, error) {
	castedMovieID, castErr := strconv.Atoi(movieID)
	if castErr != nil {
		return nil, models.ErrFooCastErr
	}
	_, castErr = time.Parse(schedule.TimeStandard, date)
	if castErr != nil {
		date = time.Now().Format(schedule.TimeStandard)
	}
	castedCinemaID, castErr := strconv.Atoi(cinemaID)
	if castErr != nil {
		return t.ScheduleRepository.GetMovieSchedule(uint64(castedMovieID), date)
	}
	return t.ScheduleRepository.GetMovieCinemaSchedule(uint64(castedMovieID), uint64(castedCinemaID), date)
}

func (t *ScheduleUseCase) GetSchedule(scheduleID string) (*models.Schedule, error) {
	castedScheduleID, castErr := strconv.Atoi(scheduleID)
	if castErr != nil {
		return nil, models.ErrFooCastErr
	}

	return t.ScheduleRepository.GetSchedule(uint64(castedScheduleID))
}
