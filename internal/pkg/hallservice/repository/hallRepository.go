package repository

import (
	"database/sql"
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"log"
)

type SQLRepository struct {
	DBConnection *sql.DB
}

func NewHallSQLRepository(connection *sql.DB) *SQLRepository {
	return &SQLRepository{
		connection,
	}
}

func (t *SQLRepository) CheckAvailability(hallID uint64, place *models.TicketPlace) (bool, error) {
	if t.DBConnection == nil {
		return false, models.ErrFooInternalDBErr
	}

	SQLResult := t.DBConnection.QueryRow("select Place_amount from cinema_hall, jsonb_array_elements(hall_params->'levels') WITH ORDINALITY arr(item_object, position) where arr.item_object->>'place' = $1 AND arr.item_object->>'row' = $2 AND ID = $3",
		place.Place, place.Row, hallID)
	if SQLResult == nil || SQLResult.Err() != nil {
		return false, models.ErrFooPlaceAlreadyBusy
	}

	return true, nil
}

func (t *SQLRepository) GetHallStructure(hallID uint64) (*models.CinemaHall, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooInternalDBErr
	}

	SQLResult := t.DBConnection.QueryRow("SELECT Place_amount,Hall_params FROM cinema_hall WHERE ID = $1", hallID)
	if SQLResult == nil || SQLResult.Err() != nil {
		return nil, models.ErrFooIncorrectSQLQuery
	}

	hallItem := new(models.CinemaHall)
	placeConfig := ""
	hallItem.ID = hallID
	ScanErr := SQLResult.Scan(&hallItem.PlaceAmount, &placeConfig)
	if ScanErr != nil {
		return nil, models.ErrFooIncorrectSQLQuery
	}
	decodeErr := json.Unmarshal([]byte(placeConfig), &hallItem.PlaceConfig)
	if decodeErr != nil {
		return nil, models.ErrFooCastErr
	}

	log.Println("rep:", hallItem.PlaceConfig)
	return hallItem, nil
}
