package models

type Cinema struct {
	ID      uint64
	Name    string
	Address string
}

type SearchCinema struct {
	Name string
}

type GetCinemaList struct {
	Limit int
	Page  int
}
