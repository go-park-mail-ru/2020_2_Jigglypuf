package models

type HallConfig struct {
	Levels []HallPlace `json:"levels"`
}

type HallPlace struct {
	Place int `json:"place"`
	Row   int `json:"row"`
}

type CinemaHall struct {
	ID          uint64
	PlaceAmount int
	PlaceConfig HallConfig
}
