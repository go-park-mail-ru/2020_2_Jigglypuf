package schedule

import "backend/internal/pkg/models"

type TimeTableUseCase interface {
	GetMovieSchedule(MovieID string, CinemaID string, date string) (*[]models.Schedule, error)
}
