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

type Actor struct {
	ID          uint64
	Name        string
	Surname     string
	Patronymic  string
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

type Recommendation struct {
	RatingMovieName string
	Movie *[]MovieList
}

type Genre struct {
	ID   uint64
	Name string
}

type GenreList []Genre
type ActorList []Actor

func (t *GenreList) Scan(src interface{}) error {
	if val, ok := src.([]byte); ok {
		return json.Unmarshal(val, t)
	}
	return ErrFooCastErr
}

func (t *GenreList) Join(cp string) string{
	result := ""
	for _, val := range *t{
		result += val.Name + cp
	}
	return result
}

func (t *ActorList) Scan(src interface{}) error {
	if val, ok := src.([]byte); ok {
		return json.Unmarshal(val, t)
	}
	return ErrFooCastErr
}

func (t *ActorList) Join(cp string) string{
	result := ""
	for _, val := range *t{
		result += val.Name + val.Surname + cp
	}
	return result
}

type RatingSet struct {
	MovieRating *Movie
	Rating      int64
}

type RatingInput struct {
	Rating    int64  `json:"rating"`
	MovieName string `json:"moviename"`
}

type GetMovieList struct {
	Limit int
	Page  int
}

type SearchMovie struct {
	Name string `json:"name"`
}

type RateMovie struct {
	ID     uint64 `json:"id"`
	Rating int64  `json:"rating"`
}

type RatingModel struct {
	ID uint64
	UserID uint64
	MovieID uint64
	MovieRating int64
}
