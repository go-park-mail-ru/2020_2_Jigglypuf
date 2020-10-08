package backend

import (
	cinemaService "backend/cinemaService"
	models "backend/models"
)

type CinemaUseCase struct{
	DBConn cinemaService.CinemaRepository
}

func NewCinemaUseCase(dbConn cinemaService.CinemaRepository) *CinemaUseCase{
	return &CinemaUseCase{
		DBConn: dbConn,
	}
}


func (t *CinemaUseCase) CreateCinema(cinema *models.Cinema) error{
	return t.DBConn.CreateCinema(cinema)
}


func (t *CinemaUseCase) GetCinema(name *string)(*models.Cinema, error){
	return t.DBConn.GetCinema(name)
}


func (t *CinemaUseCase) GetCinemaList(limit, page int)(*[]models.Cinema, error){
	return t.DBConn.GetCinemaList(limit, page)
}