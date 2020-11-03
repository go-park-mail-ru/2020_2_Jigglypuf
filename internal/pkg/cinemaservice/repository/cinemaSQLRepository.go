package repository

import (
	"backend/internal/pkg/models"
	"database/sql"
	"errors"
	"log"
)

type CinemaSQLRepository struct {
	DBConnection *sql.DB
}

func NewCinemaSQLRepository(connection *sql.DB) *CinemaSQLRepository {
	return &CinemaSQLRepository{
		DBConnection: connection,
	}
}

func (t *CinemaSQLRepository) CreateCinema(cinema *models.Cinema) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	ScanErr := t.DBConnection.QueryRow("INSERT INTO cinema (CinemaName, Address) VALUES($1,$2) RETURNING ID", cinema.Name, cinema.Address).Scan(&cinema.ID)
	if ScanErr != nil {
		log.Println(ScanErr)
		return errors.New("service not available")
	}
	return nil
}

func (t *CinemaSQLRepository) GetCinema(id uint64) (*models.Cinema, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	result := t.DBConnection.QueryRow("SELECT ID, CinemaName, Address, Hall_count FROM cinema WHERE ID = $1", id)
	rowsErr := result.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	resultCinema := new(models.Cinema)
	scanErr := result.Scan(&resultCinema.ID, &resultCinema.Name, &resultCinema.Address, &resultCinema.HallCount)
	if scanErr != nil {
		log.Println(scanErr)
		return nil, scanErr
	}

	return resultCinema, nil
}

func (t *CinemaSQLRepository) GetCinemaList(limit, page int) (*[]models.Cinema, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	resultList, DBErr := t.DBConnection.Query("SELECT ID,CinemaName,Address,Hall_count FROM cinema ORDER BY ID,CinemaName LIMIT $1 OFFSET $2", limit, page*limit)
	if DBErr != nil {
		log.Println(DBErr)
		return nil, DBErr
	}
	defer resultList.Close()

	rowsErr := resultList.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}
	cinemaList := make([]models.Cinema, 0)
	for resultList.Next() {
		cinemaItem := new(models.Cinema)
		cinemaItemScanError := resultList.Scan(&cinemaItem.ID, &cinemaItem.Name, &cinemaItem.Address, &cinemaItem.HallCount)
		if cinemaItemScanError != nil {
			log.Println(cinemaItemScanError)
			return nil, cinemaItemScanError
		}
		cinemaList = append(cinemaList, *cinemaItem)
	}

	return &cinemaList, nil
}

func (t *CinemaSQLRepository) UpdateCinema(cinema *models.Cinema) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("UPDATE cinema SET CinemaName = $1, Address = $2 WHERE ID = $3", cinema.Name, cinema.Address, cinema.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *CinemaSQLRepository) DeleteCinema(cinema *models.Cinema) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("DELETE FROM cinema WHERE ID = $1", cinema.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}
