//go:generate mockgen -source Repository.go -destination mock/RecRepositoryMock.go -package mock
package recommendation

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type Repository interface{
	GetMovieRatingsDataset() (*[]models.RecommendationDataFrame, error)
}
