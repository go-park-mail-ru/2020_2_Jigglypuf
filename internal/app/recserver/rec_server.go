package recserver

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/globalconfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/usecase"
	"github.com/gorilla/mux"
)

type RecommendationService struct {
	RecommendationDelivery   *delivery.RecommendationDelivery
	RecommendationUseCase    *usecase.ContentBasedRecommender
	RecommendationRepository *repository.RecommendationSystemRepository
	RecommendationRouter     *mux.Router
}

func configureRecommendationRouter(handler *delivery.RecommendationDelivery) *mux.Router {
	handle := mux.NewRouter()

	handle.HandleFunc(globalconfig.RecommendationsURLPattern, handler.GetRecommendedMovieList)

	return handle
}

func Start(connection *sql.DB, movieRep movieservice.MovieRepository) (*RecommendationService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	recommendationRep := repository.NewRecommendationRepository(connection)
	recommendationUC, err := usecase.NewContentBasedRecommender(recommendationRep, movieRep)
	if err != nil{
		return nil, err
	}
	recommendationDelivery := delivery.NewRecommendationDelivery(recommendationUC)
	handler := configureRecommendationRouter(recommendationDelivery)

	return &RecommendationService{
		RecommendationDelivery:   recommendationDelivery,
		RecommendationUseCase:    recommendationUC,
		RecommendationRepository: recommendationRep,
		RecommendationRouter:     handler,
	}, nil
}
