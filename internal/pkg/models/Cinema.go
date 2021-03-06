package models

import "time"

type Cinema struct {
	ID           uint64
	Name         string
	Address      string
	HallCount    int
	PathToAvatar string
	AuthorID     uint64
}

type SearchCinema struct {
	Name string `json:"name"`
}

type GetCinemaList struct {
	Limit int
	Page  int
}

type Ticket struct {
	ID              uint64
	Login           string `validate:"required,email"`
	Schedule        Schedule
	QRPath          string
	TransactionDate time.Time
	PlaceField      TicketPlace
}

type TicketInput struct {
	Login       string        `validate:"required,email" json:"login"`
	ScheduleID  uint64        `json:"scheduleID"`
	PlaceField  []TicketPlace `json:"placeField"`
	Transaction []string      `json:"-"`
}

type TicketPlace struct {
	Row   int `json:"row"`
	Place int `json:"place"`
}

type TicketInfo struct {
	UserLogin    string
	MovieName    string
	PremiereTime time.Time
	Row          int
	Place        int
}
