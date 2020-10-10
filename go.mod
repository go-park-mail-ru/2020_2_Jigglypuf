module backend

go 1.15

replace authentication => ./internal/pkg/authentication

replace cinemaService => ./internal/pkg/cinemaService

replace cookie => ./internal/pkg/cookie

replace models => ./models

replace movieService => ./internal/pkg/movieService

replace profile => ./internal/pkg/profile

replace server => ./server

require github.com/stretchr/testify v1.6.1
