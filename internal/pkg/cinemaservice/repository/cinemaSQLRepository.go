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
		return errors.New("no database connection")
	}

	result, DBErr := t.DBConnection.Exec("INSERT INTO cinema (`CinemaName`, `Address`) VALUES(?,?)", cinema.Name, cinema.Address)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	cinemaID, resultErr := result.LastInsertId()
	if resultErr != nil {
		log.Println(resultErr)
		return resultErr
	}

	cinema.ID = uint64(cinemaID)
	return nil
}

func (t *CinemaSQLRepository) GetCinema(id uint64) (*models.Cinema, error) {
	if t.DBConnection == nil {
		return nil, errors.New("no database connection")
	}

	result, DBErr := t.DBConnection.Query("SELECT ID, CinemaName, Address FROM cinema WHERE ID = ?", id)
	if DBErr != nil {
		log.Println(DBErr)
		return nil, DBErr
	}
	rowsErr := result.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	resultCinema := new(models.Cinema)
	scanErr := result.Scan(&resultCinema.ID, &resultCinema.Name, &resultCinema.Address)
	if scanErr != nil {
		log.Println(scanErr)
		return nil, scanErr
	}

	return resultCinema, nil
}

func (t *CinemaSQLRepository) GetCinemaList(limit, page int) (*[]models.Cinema, error) {
	if t.DBConnection == nil {
		return nil, errors.New("no database connection")
	}

	resultList, DBErr := t.DBConnection.Query("SELECT ID,CinemaName,Address FROM cinema ORDER BY ID,CinemaName LIMIT ? OFFSET ?", limit, page*limit)
	if DBErr != nil {
		log.Println(DBErr)
		return nil, DBErr
	}
	rowsErr := resultList.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}
	cinemaList := make([]models.Cinema, 1)
	for resultList.Next() {
		cinemaItem := new(models.Cinema)
		cinemaItemScanError := resultList.Scan(&cinemaItem.ID, &cinemaItem.Name, &cinemaItem.Address)
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
		return errors.New("no database connection")
	}

	_, DBErr := t.DBConnection.Exec("UPDATE cinema SET CinemaName = ?, Address = ?, WHERE ID = ?", cinema.Name, cinema.Address, cinema.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *CinemaSQLRepository) DeleteCinema(cinema *models.Cinema) error {
	if t.DBConnection == nil {
		return errors.New("no database connection")
	}

	_, DBErr := t.DBConnection.Exec("DELETE FROM cinema WHERE ID = ?", cinema.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}
