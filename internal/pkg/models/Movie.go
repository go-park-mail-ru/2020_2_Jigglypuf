package models

type Movie struct {
	ID                 uint64 `repository:"ID"`
	Name               string
	Description        string
	GenreList          []string
	Duration           int
	Producer           string
	Country            string
	ReleaseYear        int
	AgeGroup           int
	Rating             float64
	RatingCount        int
	PersonalRating     int64
	ActorList          []Actor
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
	GenreList          []string
	Duration           int
	Producer           string
	Country            string
	ReleaseYear        int
	AgeGroup           int
	Rating             float64
	RatingCount        int
	ActorList          []Actor
	PathToAvatar       string
	PathToSliderAvatar string
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
