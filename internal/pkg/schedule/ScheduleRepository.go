//go:generate mockgen -source ScheduleRepository.go -destination mock/ScheduleRep_mock.go -package mock
package schedule

import (
	"backend/internal/pkg/models"
)

type TimeTableRepository interface {
	GetMovieSchedule(MovieID uint64, date string) (*[]models.Schedule, error)
	GetMovieCinemaSchedule(MovieID, CinemaID uint64, date string) (*[]models.Schedule, error)
}