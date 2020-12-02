package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/promconfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"net/http"
)

type RecommendationDelivery struct {
	recommendationUseCase recommendation.UseCase
}

func NewRecommendationDelivery(useCase recommendation.UseCase) *RecommendationDelivery {
	return &RecommendationDelivery{
		useCase,
	}
}

// recommendations godoc
// @Summary recommendations
// @Description get user recommendations
// @ID get-user-recommendations
// @Success 200 {array} models.Movie
// @Failure 405 {object} models.ServerResponse
// @Failure 500 {object} models.ServerResponse
// @Router /api/recommendations/ [get]
func (t *RecommendationDelivery) GetRecommendedMovieList(w http.ResponseWriter, r *http.Request) {
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w, promconfig.GetRecommendedMovieList, &status)

	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	isAuth, authErr := r.Context().Value(session.ContextIsAuthName).(bool)
	userID, userErr := r.Context().Value(session.ContextUserIDName).(uint64)
	if !authErr || !userErr || !isAuth {
		movieList, movieErr := t.recommendationUseCase.GetPopularMovies()
		if movieErr != nil {
			models.InternalErrorHTTPResponse(&w)
			return
		}
		outputBuf, _ := json.Marshal(movieList)
		_, _ = w.Write(outputBuf)
		return
	}

	movieList, movieErr := t.recommendationUseCase.GetRecommendedMovieList(userID)
	if movieErr != nil {
		models.InternalErrorHTTPResponse(&w)
		return
	}

	status = promconfig.StatusSuccess
	outputBuf, _ := json.Marshal(movieList)

	_, _ = w.Write(outputBuf)
}
