package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"log"
	"strings"
)

type SQLRepository struct {
	DBConnection *sql.DB
}

func NewTicketSQLRepository(connection *sql.DB) *SQLRepository {
	return &SQLRepository{
		connection,
	}
}

func (t *SQLRepository) GetUserTickets(login string) (*[]models.Ticket, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	SQLResult, SQLErr := t.DBConnection.Query("SELECT ID,User_login,schedule_id,transaction_date,row,place FROM ticket WHERE User_login = $1", login)
	if SQLErr != nil || SQLResult == nil || SQLResult.Err() != nil {
		log.Println(SQLErr)
		return nil, models.ErrFooInternalDBErr
	}
	defer func() {
		_ = SQLResult.Close()
	}()

	ticketList := make([]models.Ticket, 0)
	ticketItem := new(models.Ticket)
	for SQLResult.Next() {
		ScanErr := SQLResult.Scan(&ticketItem.ID, &ticketItem.Login, &ticketItem.Schedule.ID,
			&ticketItem.TransactionDate, &ticketItem.PlaceField.Row, &ticketItem.PlaceField.Place)
		if ScanErr != nil {
			log.Println(ScanErr)
			return nil, models.ErrFooInternalDBErr
		}
		ticketList = append(ticketList, *ticketItem)
	}
	return &ticketList, nil
}

func (t *SQLRepository) GetSimpleTicket(ticketID uint64, login string) (*models.Ticket, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	SQLResult := t.DBConnection.QueryRow("SELECT ID,User_login,schedule_id,transaction_date,row,place FROM ticket WHERE ID = $1 AND User_login = $2", ticketID, login)
	if SQLResult == nil || SQLResult.Err() != nil {
		return nil, models.ErrFooInternalDBErr
	}
	ticketItem := new(models.Ticket)
	ScanErr := SQLResult.Scan(&ticketItem.ID, &ticketItem.Login, &ticketItem.Schedule.ID,
		&ticketItem.TransactionDate, &ticketItem.PlaceField.Row, &ticketItem.PlaceField.Place)
	if ScanErr != nil {
		log.Println(ScanErr)
		return nil, models.ErrFooInternalDBErr
	}
	return ticketItem, nil
}

func (t *SQLRepository) GetHallTickets(scheduleID uint64) (*[]models.TicketPlace, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	SQLResult, SQLErr := t.DBConnection.Query("SELECT t.row,t.place FROM ticket t WHERE t.schedule_id = $1",
		scheduleID)
	if SQLErr != nil || SQLResult == nil || SQLResult.Err() != nil {
		return nil, models.ErrFooIncorrectSQLQuery
	}
	placeList := make([]models.TicketPlace, 0)
	placeItem := new(models.TicketPlace)
	for SQLResult.Next() {
		ScanErr := SQLResult.Scan(&placeItem.Row, &placeItem.Place)
		if ScanErr != nil {
			return nil, models.ErrFooInternalDBErr
		}
		placeList = append(placeList, *placeItem)
	}
	return &placeList, nil
}

func (t *SQLRepository) GetTicketByTransaction(transaction string) (*models.TicketInfo, error){
	if t.DBConnection == nil{
		return nil, models.ErrFooNoDBConnection
	}
	query := "SELECT User_login,m.moviename,s.premiere_time,row,place FROM ticket JOIN schedule s on s.id = ticket.schedule_id JOIN movie m on m.id = s.movie_id WHERE transaction = $1"
	resp := new(models.TicketInfo)
	ScanErr := t.DBConnection.QueryRow(query, transaction).Scan(&resp.UserLogin, &resp.MovieName,
		&resp.PremiereTime, &resp.Row, &resp.Place)
	if ScanErr != nil{
		log.Println(ScanErr)
		return nil, models.ErrFooInternalDBErr
	}

	return resp, nil
}

func (t *SQLRepository) CreateTicket(ticket *models.TicketInput) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}
	var ID uint64 = 0
	insertQuery := "INSERT INTO ticket (User_login,schedule_id,row,place, transaction) VALUES %s"
	query, args := bulkInsert(ticket, insertQuery)
	query += " RETURNING ID"
	ScanErr := t.DBConnection.QueryRow(query,
		*args...).Scan(&ID)
	if ScanErr != nil {
		return models.ErrFooIncorrectSQLQuery
	}

	return nil
}

func bulkInsert(rows *models.TicketInput, query string) (string, *[]interface{}) {
	ValueStrings := make([]interface{}, 0)
	QueryStrings := make([]string, 0)
	i := 0
	for index, val := range rows.PlaceField {
		QueryStrings = append(QueryStrings, fmt.Sprintf("($%d, $%d, $%d,$%d,$%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		ValueStrings = append(ValueStrings, rows.Login, rows.ScheduleID, val.Row, val.Place, rows.Transaction[index])
		i++
	}
	smtp := fmt.Sprintf(query, strings.Join(QueryStrings, ","))
	return smtp, &ValueStrings
}
