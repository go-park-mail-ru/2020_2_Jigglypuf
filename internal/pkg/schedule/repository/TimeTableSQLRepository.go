package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"log"
)

type ScheduleSQLRepository struct {
	DBConnection *sql.DB
}

func NewScheduleSQLRepository(connection *sql.DB) *ScheduleSQLRepository {
	return &ScheduleSQLRepository{
		DBConnection: connection,
	}
}

func (t *ScheduleSQLRepository) GetMovieCinemaSchedule(movieID, cinemaID uint64, date string) (*[]models.Schedule, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	DBRows, DBErr := t.DBConnection.Query("SELECT ID,Movie_id,Cinema_ID,Hall_ID,Premiere_time,Cost FROM schedule "+
		"WHERE Movie_ID = $1 AND Cinema_ID = $2 AND DATE(Premiere_time) = DATE($3)", movieID, cinemaID, date)
	if DBErr != nil || DBRows == nil || DBRows.Err() != nil {
		log.Println(DBErr)
		return nil, models.ErrFooInternalDBErr
	}
	scheduleList := make([]models.Schedule, 0)
	scheduleItem := new(models.Schedule)
	for DBRows.Next() {
		ScanErr := DBRows.Scan(&scheduleItem.ID, &scheduleItem.MovieID, &scheduleItem.CinemaID, &scheduleItem.HallID,
			&scheduleItem.PremierTime, &scheduleItem.Cost)
		if ScanErr != nil {
			log.Println(ScanErr)
			return nil, models.ErrFooNoDBConnection
		}
		scheduleList = append(scheduleList, *scheduleItem)
	}

	return &scheduleList, nil
}

func (t *ScheduleSQLRepository) GetMovieSchedule(movieID uint64, date string) (*[]models.Schedule, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	DBRows, DBErr := t.DBConnection.Query("SELECT ID,Movie_id,Cinema_ID,Hall_ID,Premiere_time, Cost FROM schedule "+
		"WHERE Movie_ID = $1 AND DATE(Premiere_time) = DATE($2)", movieID, date)
	if DBErr != nil || DBRows != nil && DBRows.Err() != nil {
		log.Println(DBErr)
		return nil, models.ErrFooInternalDBErr
	}
	scheduleList := make([]models.Schedule, 0)
	scheduleItem := new(models.Schedule)
	for DBRows.Next() {
		ScanErr := DBRows.Scan(&scheduleItem.ID, &scheduleItem.MovieID, &scheduleItem.CinemaID, &scheduleItem.HallID,
			&scheduleItem.PremierTime, &scheduleItem.Cost)
		if ScanErr != nil {
			log.Println(ScanErr)
			return nil, models.ErrFooNoDBConnection
		}
		scheduleList = append(scheduleList, *scheduleItem)
	}

	return &scheduleList, nil
}

func (t *ScheduleSQLRepository) GetScheduleHallID(scheduleID uint64) (uint64, error) {
	if t.DBConnection == nil {
		return 0, models.ErrFooNoDBConnection
	}
	var HallID uint64 = 0
	ScanErr := t.DBConnection.QueryRow("SELECT hall_id from schedule where id = $1", scheduleID).Scan(&HallID)
	if ScanErr != nil {
		return 0, models.ErrFooIncorrectInputInfo
	}
	return HallID, nil
}

func (t *ScheduleSQLRepository) GetSchedule(scheduleID uint64) (*models.Schedule, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	resultItem := new(models.Schedule)
	ScanErr := t.DBConnection.QueryRow("SELECT id, movie_id, cinema_id, hall_id, premiere_time, cost from schedule where id = $1", scheduleID).Scan(&resultItem.ID,
		&resultItem.MovieID, &resultItem.CinemaID, &resultItem.HallID, &resultItem.PremierTime, &resultItem.Cost)
	if ScanErr != nil {
		return nil, models.ErrFooIncorrectInputInfo
	}

	return resultItem, nil
}
