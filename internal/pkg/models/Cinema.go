package models

import "time"

type Cinema struct {
	ID        uint64
	Name      string
	Address   string
	HallCount int
	PathToAvatar string
	AuthorID  uint64
}

type SearchCinema struct {
	Name string
}

type GetCinemaList struct {
	Limit int
	Page  int
}

type CinemaHall struct {
	ID          uint64
	PlaceAmount int
	PlaceConfig string
}

type Ticket struct {
	ID              uint64
	Login           string `validate:"required,email"`
	Schedule        Schedule
	TransactionDate time.Time
	PlaceField      TicketPlace
}

type TicketInput struct {
	Login      string `validate:"required,email"`
	ScheduleID uint64
	HallID     uint64
	PlaceField TicketPlace
}

type TicketPlace struct {
	Row   int
	Place int
}
