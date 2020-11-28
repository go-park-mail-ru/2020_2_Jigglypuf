package recserver

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	config "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/usecase"
	"github.com/gorilla/mux"
	"sync"
	"time"
)

type RecommendationService struct {
	RecommendationDelivery   *delivery.RecommendationDelivery
	RecommendationUseCase    *usecase.RecommendationSystemUseCase
	RecommendationRepository *repository.RecommendationSystemRepository
	RecommendationRouter     *mux.Router
}

func configureRecommendationRouter(handler *delivery.RecommendationDelivery) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(config.URLPattern, handler.GetRecommendedMovieList)

	return router
}

func Start(connection *sql.DB, mutex *sync.RWMutex, sleepTime time.Duration) (*RecommendationService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	recommendationRep := repository.NewRecommendationRepository(connection)
	recommendationUC := usecase.NewRecommendationSystemUseCase(recommendationRep, sleepTime, mutex)
	recommendationDelivery := delivery.NewRecommendationDelivery(recommendationUC)
	router := configureRecommendationRouter(recommendationDelivery)

	return &RecommendationService{
		RecommendationDelivery:   recommendationDelivery,
		RecommendationUseCase:    recommendationUC,
		RecommendationRepository: recommendationRep,
		RecommendationRouter:     router,
	}, nil
}
