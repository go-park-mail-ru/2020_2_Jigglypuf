module server

go 1.15

replace authentication => ../authentication

replace models => ../models

replace cookie => ../cookie

replace cinemaService => ../cinemaService

replace movieService => ../movieService

replace profile => ../profile

require (
	authentication v0.0.0-00010101000000-000000000000
	cinemaService v0.0.0-00010101000000-000000000000
	cookie v0.0.0-00010101000000-000000000000
	movieService v0.0.0-00010101000000-000000000000
	profile v0.0.0-00010101000000-000000000000
)
