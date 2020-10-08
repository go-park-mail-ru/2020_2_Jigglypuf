package backend

type Movie struct{
	Id uint64
	Name string
	Description string
	Rating float64
	PathToAvatar string
}

type RatingSet struct{
	MovieRating *Movie
	Rating int64
}

type RatingInput struct{
	Rating int64
	MovieName string
}

type GetMovieList struct{
	Limit int
	Page int
}

type SearchMovie struct{
	Name string
}

type RateMovie struct{
	Name string
	Rating int64
}