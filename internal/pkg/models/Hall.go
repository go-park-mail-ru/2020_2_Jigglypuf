package models

type HallConfig struct {
	Levels []HallPlace
}

type HallPlace struct {
	Place int
	Row   int
}

type CinemaHall struct {
	ID          uint64
	PlaceAmount int
	PlaceConfig HallConfig
}
