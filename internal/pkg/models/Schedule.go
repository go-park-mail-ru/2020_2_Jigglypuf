package models

import "time"

type Schedule struct{
	ID uint64
	MovieID uint64
	CinemaID uint64
	HallID	uint64
	PremierTime time.Time
}