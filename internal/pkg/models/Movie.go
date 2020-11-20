package models

import "encoding/json"

type Movie struct {
	ID                 uint64 `repository:"ID"`
	Name               string
	Description        string
	GenreList          GenreList
	Duration           int
	Producer           string
	Country            string
	ReleaseYear        int
	AgeGroup           int
	Rating             float64
	RatingCount        int
	PersonalRating     int64
	ActorList          ActorList
	PathToAvatar       string
	PathToSliderAvatar string
}

type Actor struct{
	ID uint64
	Name string
	Surname string
	Patronymic string
	Description string
}

type MovieList struct {
	ID                 uint64 `repository:"ID"`
	Name               string
	Description        string
	GenreList          GenreList
	Duration           int
	Producer           string
	Country            string
	ReleaseYear        int
	AgeGroup           int
	Rating             float64
	RatingCount        int
	ActorList          ActorList
	PathToAvatar       string
	PathToSliderAvatar string
}

type Genre struct{
	ID uint64
	Name string
}

type GenreList []Genre
type ActorList []Actor
func (t *GenreList) Scan(src interface{}) error{
	if val, ok := src.([]byte); ok{
		return json.Unmarshal(val, t)
	}
	return ErrFooCastErr
}

func (t *ActorList) Scan(src interface{})error{
	if val, ok := src.([]byte); ok{
		return json.Unmarshal(val, t)
	}
	return ErrFooCastErr
}

type RatingSet struct {
	MovieRating *Movie
	Rating      int64
}

type RatingInput struct {
	Rating    int64
	MovieName string
}

type GetMovieList struct {
	Limit int
	Page  int
}

type SearchMovie struct {
	Name string
}

type RateMovie struct {
	ID     uint64
	Rating int64
}
