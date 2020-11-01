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

func (t *MovieUseCase) GetMovie(id uint64, isAuth bool, userID uint64) (*models.Movie, error) {
	movie, err := t.DBConn.GetMovie(id)
	if err != nil{
		return nil,err
	}
	if isAuth{
		rating, ratingErr := t.DBConn.GetRating(userID, id)
		if ratingErr == nil{
			movie.PersonalRating = rating
		}
	}
	return movie,nil
}

func (t *MovieUseCase) GetMovieList(limit, page int) (*[]models.MovieList, error) {
	return t.DBConn.GetMovieList(limit, page)
}

func (t *MovieUseCase) CreateMovie(movie *models.Movie) error {
	return t.DBConn.CreateMovie(movie)
}

func (t *MovieUseCase) UpdateMovie(movie *models.Movie) error {
	return t.DBConn.UpdateMovie(movie)
}

func (t *MovieUseCase) RateMovie(user *models.User, id uint64, rating int64) error {
	personalRatingErr := t.DBConn.RateMovie(user, id, rating)
	if personalRatingErr != nil {
		return personalRatingErr
	}

	movieRatingErr := t.DBConn.UpdateMovieRating(id, rating)
	return movieRatingErr
}

func (t *MovieUseCase) GetRating(user *models.User, id uint64) (int64, error) {
	return t.DBConn.GetRating(user.ID, id)
}

func (t *MovieUseCase) GetMoviesInCinema(limit, page int) (*[]models.MovieList, error) {
	return t.DBConn.GetMoviesInCinema(limit, page)
}
