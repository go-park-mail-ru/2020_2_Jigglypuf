package schedule

import (
	"backend/internal/pkg/models"
	"time"
)

type TimeTableRepository interface{
	GetMovieSchedule(MovieID uint64, date time.Time)(*[]models.Schedule, error)
	GetMovieCinemaSchedule (MovieID, CinemaID uint64, date time.Time)(*[]models.Schedule, error)
}