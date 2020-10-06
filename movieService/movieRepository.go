package movieService

import(
	"models"
)

type MovieRepository interface{
	CreateMovie(*models.Movie)error
	UpdateMovie(*models.Movie) error
	GetMovie(name string)(*models.Movie, error)
	GetMovieList(limit, page int)(*[]models.Movie, error)
}