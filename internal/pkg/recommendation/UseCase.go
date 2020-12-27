package recommendation

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type UseCase interface {
	MakeRecommendation(userID uint64, limit int) (*models.Recommendation, error)
	GetPopularMovies()(*models.Recommendation, error)
}
