package movieservice

import "backend/internal/pkg/models"

type MovieUseCase interface {
	GetMovie(id uint64) (*models.Movie, error)
	GetMovieList(limit, page int) (*[]models.Movie, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovie(movie *models.Movie) error
	RateMovie(user *models.User, id uint64, rating int64) error
	GetRating(user *models.User, id uint64) (int64, error)
}
