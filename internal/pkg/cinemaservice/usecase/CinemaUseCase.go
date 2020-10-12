package usecase

import (
	"backend/internal/pkg/cinemaservice"
	"backend/internal/pkg/models"
)

type CinemaUseCase struct {
	DBConn cinemaservice.Repository
}

func NewCinemaUseCase(dbConn cinemaservice.Repository) *CinemaUseCase {
	return &CinemaUseCase{
		DBConn: dbConn,
	}
}

func (t *CinemaUseCase) CreateCinema(cinema *models.Cinema) error {
	return t.DBConn.CreateCinema(cinema)
}

func (t *CinemaUseCase) GetCinema(name *string) (*models.Cinema, error) {
	return t.DBConn.GetCinema(name)
}

func (t *CinemaUseCase) GetCinemaList(limit, page int) (*[]models.Cinema, error) {
	return t.DBConn.GetCinemaList(limit, page)
}