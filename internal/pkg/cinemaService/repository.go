package cinemaService

import (
	"backend/internal/pkg/models"
)

type CinemaRepository interface{
	CreateCinema( cinema *models.Cinema ) error
	GetCinema( name *string )( *models.Cinema, error )
	GetCinemaList(limit, page int)( *[]models.Cinema, error )
	UpdateCinema( cinema *models.Cinema ) error
	DeleteCinema( cinema *models.Cinema ) error
}