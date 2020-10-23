module backend

go 1.15

replace authentication => ./internal/pkg/authentication

replace cinemaService => ./internal/pkg/cinemaservice

replace cookie => ./internal/pkg/middleware/cookie

replace models => ./models

replace movieService => ./internal/pkg/movieservice

replace profile => ./internal/pkg/profile

replace server => ./server

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/stretchr/testify v1.6.1
	github.com/tarantool/go-tarantool v0.0.0-20200816172506-a535b8e0224a
)
