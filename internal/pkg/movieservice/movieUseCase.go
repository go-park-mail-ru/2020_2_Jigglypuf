//go:generate mockgen -source movieUseCase.go -destination mock/MovieUC_mock.go -package mock
package movieservice

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type MovieUseCase interface {
	GetMovie(id uint64, isAuth bool, userID uint64) (*models.Movie, error)
	GetMovieList(limit, page int) (*[]models.MovieList, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovie(movie *models.Movie) error
	RateMovie(userID uint64, id uint64, rating int64) error
	GetRating(user *models.User, id uint64) (int64, error)
	GetActualMovies(limit, page int, date []string) (*[]models.MovieList, error)
}
