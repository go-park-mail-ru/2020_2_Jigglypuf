package usecase

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"strconv"
	"time"
)

type ScheduleUseCase struct{
	ScheduleRepository schedule.TimeTableRepository
}

func NewTimeTableUseCase(repository schedule.TimeTableRepository) *ScheduleUseCase{
	return &ScheduleUseCase{
		repository,
	}
}

func (t *ScheduleUseCase) GetMovieSchedule(MovieID, CinemaID string, date string)(*[]models.Schedule, error){
	castedMovieID,castErr := strconv.Atoi(MovieID)
	if castErr != nil{
		return nil,models.ErrFooCastErr
	}
	dateTime, castErr := time.Parse("2006-09-02",date)
	if castErr != nil{
		dateTime = time.Now()
	}
	castedCinemaID, castErr := strconv.Atoi(CinemaID)
	if castErr != nil{
		return t.ScheduleRepository.GetMovieSchedule(uint64(castedMovieID),dateTime)
	}
	return t.ScheduleRepository.GetMovieCinemaSchedule(uint64(castedMovieID),uint64(castedCinemaID),dateTime)
}