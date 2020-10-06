package usecase

import (
	"models"
	"movieService"
)

type MovieUseCase struct{
	DBConn movieService.MovieRepository
}

func NewMovieUseCase(rep movieService.MovieRepository) *MovieUseCase{
	return &MovieUseCase{
		DBConn: rep,
	}
}

func (t *MovieUseCase) GetMovie(name string)(*models.Movie, error){
	return t.DBConn.GetMovie(name)
}


func (t *MovieUseCase) GetMovieList(params *models.GetMovieList)(*[]models.Movie, error){
	return t.DBConn.GetMovieList(params.Limit, params.Page)
}


func (t *MovieUseCase) CreateMovie(movie *models.Movie)error{
	return t.DBConn.CreateMovie(movie)
}

func (t *MovieUseCase) UpdateMovie(movie *models.Movie)error{
	return t.DBConn.UpdateMovie(movie)
}