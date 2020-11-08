//go:generate mockgen -source ScheduleUsecase.go -destination mock/ScheduleUC_mock.go -package mock
package schedule

import "backend/internal/pkg/models"

type TimeTableUseCase interface {
	GetMovieSchedule(MovieID string, CinemaID string, date string) (*[]models.Schedule, error)
}
