package backend

type Cinema struct{
	Id uint64
	Name string
	Address string
}

type SearchCinema struct{
	Name string
}

type GetCinemaList struct{
	Limit int
	Page int
}
