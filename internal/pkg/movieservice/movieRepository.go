//go:generate mockgen -source movieRepository.go -destination mock/MovieRep_mock.go -package mock
package movieservice

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type MovieRepository interface {
	CreateMovie(*models.Movie) error
	UpdateMovie(*models.Movie) error
	GetMovie(id uint64) (*models.Movie, error)
	GetMovieList(limit, page int) (*[]models.MovieList, error)
	RateMovie(user *models.User, id uint64, rating int64) error
	GetRating(userID uint64, id uint64) (int64, error)
	UpdateMovieRating(movieID uint64, ratingScore int64) error
	GetMoviesInCinema(limit, page int, date string, allTime bool) (*[]models.MovieList, error)
	GetAllMovies() ([]models.MovieList, error)
}
