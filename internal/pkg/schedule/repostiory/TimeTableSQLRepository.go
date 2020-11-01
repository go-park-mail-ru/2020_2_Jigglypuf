package repostiory

import (
	"backend/internal/pkg/models"
	"database/sql"
	"log"
)

type ScheduleSQLRepository struct{
	DBConnection *sql.DB
}

func NewScheduleSQLRepository(connection *sql.DB) *ScheduleSQLRepository{
	return &ScheduleSQLRepository{
		DBConnection: connection,
	}
}

func (t *ScheduleSQLRepository) GetMovieCinemaSchedule (MovieID, CinemaID uint64, date string)(*[]models.Schedule, error){
	if t.DBConnection == nil{
		return nil,models.ErrFooNoDBConnection
	}

	DBRows, DBErr := t.DBConnection.Query("SELECT ID,Movie_id,Cinema_ID,Hall_ID,Premiere_time FROM schedule " +
		"WHERE Movie_ID = $1 AND Cinema_ID = $2 AND DATE(Premiere_time) = DATE($3)", MovieID,CinemaID,date)
	if DBErr != nil || DBRows == nil || DBRows.Err() != nil{
		log.Println(DBErr)
		return nil,models.ErrFooInternalDBErr
	}
	scheduleList := make([]models.Schedule, 0)
	scheduleItem := new(models.Schedule)
	for DBRows.Next(){
		ScanErr := DBRows.Scan(&scheduleItem.ID, &scheduleItem.MovieID, &scheduleItem.CinemaID, &scheduleItem.HallID, &scheduleItem.PremierTime)
		if ScanErr != nil{
			log.Println(ScanErr)
			return nil,models.ErrFooNoDBConnection
		}
		scheduleList = append(scheduleList,*scheduleItem)
	}

	return &scheduleList,nil
}

func (t *ScheduleSQLRepository) GetMovieSchedule(MovieID uint64, date string)(*[]models.Schedule, error){
	if t.DBConnection == nil{
		return nil,models.ErrFooNoDBConnection
	}

	DBRows, DBErr := t.DBConnection.Query("SELECT ID,Movie_id,Cinema_ID,Hall_ID,Premiere_time FROM schedule " +
		"WHERE Movie_ID = $1 AND DATE(Premiere_time) = DATE($2)", MovieID,date)
	if DBErr != nil || DBRows != nil && DBRows.Err() != nil{
		log.Println(DBErr)
		return nil,models.ErrFooInternalDBErr
	}
	scheduleList := make([]models.Schedule, 0)
	scheduleItem := new(models.Schedule)
	for DBRows.Next(){
		ScanErr := DBRows.Scan(&scheduleItem.ID, &scheduleItem.MovieID, &scheduleItem.CinemaID, &scheduleItem.HallID, &scheduleItem.PremierTime)
		if ScanErr != nil{
			log.Println(ScanErr)
			return nil,models.ErrFooNoDBConnection
		}
		scheduleList = append(scheduleList,*scheduleItem)
	}

	return &scheduleList,nil
}