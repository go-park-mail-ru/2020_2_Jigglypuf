package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
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

func (t *CinemaUseCase) GetCinema(name uint64) (*models.Cinema, error) {
	return t.DBConn.GetCinema(name)
}

func (t *CinemaUseCase) GetCinemaList(limit, page int) (*[]models.Cinema, error) {
	page -= 1
	if page < 0 || limit < 0{
		return nil,models.ErrFooIncorrectInputInfo
	}
	return t.DBConn.GetCinemaList(limit, page)
}
