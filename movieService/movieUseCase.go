package movieService

import "models"

type MovieUseCase interface{
	GetMovie(name string)(*models.Movie, error)
	GetMovieList(limit, page int)(*[]models.Movie, error)
	CreateMovie(movie *models.Movie)error
	UpdateMovie(movie *models.Movie)error
}