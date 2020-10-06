package movieService

import "models"

type MovieUseCase interface{
	GetMovie(name string)(*models.Movie, error)
	GetMovieList(limit, page int)(*[]models.Movie, error)
	CreateMovie(movie *models.Movie)error
	UpdateMovie(movie *models.Movie)error
	RateMovie(user *models.User,name string, rating int64)error
	GetRating(user *models.User, name string)(int64, error)
}