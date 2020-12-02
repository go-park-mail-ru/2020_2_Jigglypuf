package recommendation

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type UseCase interface {
	GetRecommendedMovieList(uint64) (*[]models.Movie, error)
	GetPopularMovies() (*[]models.Movie, error)
}
