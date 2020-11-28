package delivery

import(
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"net/http"
)

type RecommendationDelivery struct{
	recommendationUseCase recommendation.UseCase
}


func NewRecommendationDelivery(useCase recommendation.UseCase) *RecommendationDelivery{
	return &RecommendationDelivery{
		useCase,
	}
}


func (t *RecommendationDelivery) GetRecommendedMovieList(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		models.BadMethodHTTPResponse(&w)
		return
	}

	isAuth, authErr := r.Context().Value(session.ContextIsAuthName).(bool)
	userID, userErr := r.Context().Value(session.ContextUserIDName).(uint64)
	var movieList *[]models.Movie = nil
	var movieErr error = nil
	if !authErr || !userErr || !isAuth{
		movieList, movieErr = t.recommendationUseCase.GetPopularMovies()
		if movieErr != nil{
			models.InternalErrorHTTPResponse(&w)
			return
		}
	}

	movieList, movieErr = t.recommendationUseCase.GetRecommendedMovieList(userID)
	if movieErr != nil{
		models.InternalErrorHTTPResponse(&w)
		return
	}

	outputBuf, _ := json.Marshal(movieList)

	_,_ = w.Write(outputBuf)
}