package repository

import (
	"backend/internal/pkg/models"
	"database/sql"
	"log"
)

type SQLRepository struct{
	DBConnection *sql.DB
}

func NewTicketSQLRepository(connection *sql.DB)*SQLRepository{
	return &SQLRepository{
		connection,
	}
}

func (t *SQLRepository) GetUserTickets(username string)(*[]models.Ticket, error){
	if t.DBConnection == nil{
		return nil,models.ErrFooNoDBConnection
	}

	SQLResult, SQLErr := t.DBConnection.Query("SELECT ID,User_login,schedule_id,transaction_date,row,place FROM ticket WHERE User_login = $1", username)
	if SQLErr != nil || SQLResult == nil || SQLResult.Err() != nil{
		log.Println(SQLErr)
		return nil,models.ErrFooInternalDBErr
	}
	defer SQLResult.Close()

	ticketList := make([]models.Ticket,0)
	ticketItem := new(models.Ticket)
	for SQLResult.Next(){
		ScanErr := SQLResult.Scan(&ticketItem.ID, &ticketItem.Username, &ticketItem.Schedule.ID,
			&ticketItem.TransactionDate, &ticketItem.PlaceField.Row, &ticketItem.PlaceField.Place)
		if ScanErr != nil{
			log.Println(ScanErr)
			return nil,models.ErrFooInternalDBErr
		}
		ticketList = append(ticketList, *ticketItem)
	}
	return &ticketList, nil
}

func (t *SQLRepository) GetSimpleTicket(ticketID uint64, username string)(*models.Ticket,error){
	if t.DBConnection == nil{
		return nil,models.ErrFooNoDBConnection
	}

	SQLResult := t.DBConnection.QueryRow("SELECT ID,User_login,schedule_id,transaction_date,row,place FROM ticket WHERE ID = $1 AND User_login = $2", ticketID, username)
	if SQLResult == nil || SQLResult.Err() != nil{
		return nil,models.ErrFooInternalDBErr
	}
	ticketItem := new(models.Ticket)
	ScanErr := SQLResult.Scan(&ticketItem.ID, &ticketItem.Username, &ticketItem.Schedule.ID,
		&ticketItem.TransactionDate, &ticketItem.PlaceField.Row, &ticketItem.PlaceField.Place)
	if ScanErr != nil{
		log.Println(ScanErr)
		return nil,models.ErrFooInternalDBErr
	}
	return ticketItem, nil
}

func (t *SQLRepository) GetHallTickets(scheduleID uint64)(*[]models.TicketPlace, error){
	if t.DBConnection == nil{
		return nil,models.ErrFooNoDBConnection
	}

	SQLResult,SQLErr := t.DBConnection.Query("SELECT t.row,t.place FROM ticket t WHERE t.schedule_id = $1",
		scheduleID)
	if SQLErr != nil || SQLResult == nil || SQLResult.Err() != nil{
		return nil,models.ErrFooIncorrectSQLQuery
	}
	placeList := make([]models.TicketPlace, 0)
	placeItem := new(models.TicketPlace)
	for SQLResult.Next(){
		ScanErr := SQLResult.Scan(&placeItem.Row, &placeItem.Place)
		if ScanErr != nil{
			return nil,models.ErrFooInternalDBErr
		}
		placeList = append(placeList, *placeItem)
	}
	return &placeList, nil
}


func (t *SQLRepository) CreateTicket(ticket *models.Ticket) error{
	if t.DBConnection == nil{
		return models.ErrFooNoDBConnection
	}

	ScanErr := t.DBConnection.QueryRow("INSERT INTO ticket (User_login,schedule_id,row,place) VALUES($1,$2,$3,$4) RETURNING ID",
		ticket.Username, ticket.Schedule.ID, ticket.PlaceField.Row, ticket.PlaceField.Place).Scan(&ticket.ID)
	if ScanErr != nil{
		return models.ErrFooIncorrectSQLQuery
	}

	return nil
}
