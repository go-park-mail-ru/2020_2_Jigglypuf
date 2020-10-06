module movieService

go 1.15

replace models => ../models

replace authentication => ../authentication

require (
	authentication v0.0.0-00010101000000-000000000000
	models v0.0.0-00010101000000-000000000000
)
