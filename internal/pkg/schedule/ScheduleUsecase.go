//go:generate mockgen -source ScheduleUsecase.go -destination mock/ScheduleUC_mock.go -package mock
package schedule

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type TimeTableUseCase interface {
	GetMovieSchedule(MovieID string, CinemaID string, date string) (*[]models.Schedule, error)
}
