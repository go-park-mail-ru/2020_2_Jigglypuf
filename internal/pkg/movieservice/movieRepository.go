package movieservice

import (
	"backend/internal/pkg/models"
)

type MovieRepository interface {
	CreateMovie(*models.Movie) error
	UpdateMovie(*models.Movie) error
	GetMovie(id uint64) (*models.Movie, error)
	GetMovieList(limit, page int) (*[]models.Movie, error)
	RateMovie(user *models.User, id uint64, rating int64) error
	GetRating(user *models.User, id uint64) (int64, error)
	UpdateMovieRating(movieID uint64, ratingScore int64) error
	GetMoviesInCinema(limit, page int) (*[]models.Movie, error)
}
