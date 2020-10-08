module backend

go 1.15

replace authentication => ./authentication

replace cinemaService => ./cinemaService

replace cookie => ./cookie

replace models => ./models

replace movieService => ./movieService

replace profile => ./profile

replace server => ./server

require (
	github.com/stretchr/testify v1.6.1
	golang.org/x/tools v0.0.0-20201008025239-9df69603baec // indirect
)
