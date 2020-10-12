package usecase

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/movieservice"
)

type MovieUseCase struct {
	DBConn movieservice.MovieRepository
}

func NewMovieUseCase(rep movieservice.MovieRepository) *MovieUseCase {
	return &MovieUseCase{
		DBConn: rep,
	}
}

func (t *MovieUseCase) GetMovie(name string) (*models.Movie, error) {
	return t.DBConn.GetMovie(name)
}

func (t *MovieUseCase) GetMovieList(limit, page int) (*[]models.Movie, error) {
	return t.DBConn.GetMovieList(limit, page)
}

func (t *MovieUseCase) CreateMovie(movie *models.Movie) error {
	return t.DBConn.CreateMovie(movie)
}

func (t *MovieUseCase) UpdateMovie(movie *models.Movie) error {
	return t.DBConn.UpdateMovie(movie)
}

func (t *MovieUseCase) RateMovie(user *models.User, name string, rating int64) error {
	return t.DBConn.RateMovie(user, name, rating)
}

func (t *MovieUseCase) GetRating(user *models.User, name string) (int64, error) {
	return t.DBConn.GetRating(user, name)
}
