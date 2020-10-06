package models

type Movie struct{
	Id uint64
	Name string
	Description string
	PathToAvatar string
}


type GetMovieList struct{
	Limit int
	Page int
}

type SearchMovie struct{
	Name string
}