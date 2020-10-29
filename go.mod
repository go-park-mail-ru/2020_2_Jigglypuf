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
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/go-openapi/spec v0.19.11 // indirect
	github.com/go-openapi/swag v0.19.11 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.8.0
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.9
	github.com/tarantool/go-tarantool v0.0.0-20200816172506-a535b8e0224a
	golang.org/x/net v0.0.0-20201024042810-be3efd7ff127 // indirect
	golang.org/x/tools v0.0.0-20201023174141-c8cfbd0f21e6 // indirect
)
