package models

type Cinema struct {
	ID      uint64
	Name    string
	Address string
	HallCount int
	AuthorID uint64
}

type SearchCinema struct {
	Name string
}

type GetCinemaList struct {
	Limit int
	Page  int
}
