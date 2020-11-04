package models

type Movie struct {
	ID             uint64 `repository:"ID"`
	Name           string
	Description    string
	Genre          string
	Duration       int
	Producer       string
	Country        string
	ReleaseYear    int
	AgeGroup       int
	Rating         float64
	RatingCount    int
	PersonalRating int64
	PathToAvatar   string
}

type MovieList struct {
	ID           uint64 `repository:"ID"`
	Name         string
	Description  string
	Genre        string
	Duration     int
	Producer     string
	Country      string
	ReleaseYear  int
	AgeGroup     int
	Rating       float64
	RatingCount  int
	PathToAvatar string
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
