package delivery

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"encoding/json"
	"net/http"
)

type ScheduleDelivery struct {
	UseCase schedule.TimeTableUseCase
}

func NewScheduleDelivery(useCase schedule.TimeTableUseCase) *ScheduleDelivery {
	return &ScheduleDelivery{
		useCase,
	}
}

// Schedule godoc
// @Summary Get movie schedule
// @Description Returns movie schedule by getting movie id, cinema id and day(date) in format schedule.TimeStandard
// @ID movie-schedule-id
// @Param movie_id query int true "movie_id"
// @Param cinema_id query int true "cinema_id"
// @Param date query string true "date"
// @Success 200 {array} models.Schedule
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Failure 500 {object} models.ServerResponse "internal error"
// @Router /schedule/ [get]
func (t *ScheduleDelivery) GetMovieSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	movieID := r.URL.Query().Get(schedule.MovieIDQueryParamName)
	cinemaID := r.URL.Query().Get(schedule.CinemaIDQueryParamName)
	date := r.URL.Query().Get(schedule.DateQueryParamName)
	resultList, paramErr := t.UseCase.GetMovieSchedule(movieID, cinemaID, date)
	if paramErr != nil {
		models.BadBodyHTTPResponse(&w, paramErr)
		return
	}

	outputBuf, castErr := json.Marshal(resultList)
	if castErr != nil {
		models.InternalErrorHTTPResponse(&w)
		return
	}

	_, _ = w.Write(outputBuf)
}
