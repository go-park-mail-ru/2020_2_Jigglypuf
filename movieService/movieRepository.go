package movieService

import(
	"backend/models"
)

type MovieRepository interface{
	CreateMovie(*models.Movie)error
	UpdateMovie(*models.Movie) error
	GetMovie(name string)(*models.Movie, error)
	GetMovieList(limit, page int)(*[]models.Movie, error)
	RateMovie(user *models.User,name string, rating int64)error
	GetRating(user *models.User, name string)(int64, error)
}