//go:generate mockgen -source Repository.go -destination mock/RecRepositoryMock.go -package mock
package recommendation

import (
	set "github.com/deckarep/golang-set"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type Repository interface {
	GetMovieRatingsDataset() (*[]models.RecommendationDataFrame, error)
	GetRecommendedMovieList(set *set.Set) (*[]models.Movie, error)
	GetPopularMovies() (*[]models.Movie, error)
}
