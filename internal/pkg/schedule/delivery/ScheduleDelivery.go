package delivery

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"encoding/json"
	"net/http"
)

type ScheduleDelivery struct{
	UseCase schedule.TimeTableUseCase
}

func NewScheduleDelivery(useCase schedule.TimeTableUseCase) *ScheduleDelivery{
	return &ScheduleDelivery{
		useCase,
	}
}

func (t *ScheduleDelivery) GetMovieSchedule(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		models.BadMethodHTTPResponse(&w)
		return
	}

	movieID := r.URL.Query().Get(schedule.MovieIDQueryParamName)
	cinemaID := r.URL.Query().Get(schedule.CinemaIDQueryParamName)
	date := r.URL.Query().Get(schedule.DateQueryParamName)
	resultList, paramErr := t.UseCase.GetMovieSchedule(movieID,cinemaID,date)
	if paramErr != nil{
		models.BadBodyHTTPResponse(&w, paramErr)
		return
	}

	outputBuf, castErr := json.Marshal(resultList)
	if castErr != nil{
		models.InteralErrorHttpResponse(&w)
		return
	}

	_, _ = w.Write(outputBuf)
	return
}