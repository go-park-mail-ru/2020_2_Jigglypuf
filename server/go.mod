module server

go 1.15

replace authentication => ../authentication

replace models => ../models

replace cookie => ../cookie

require (
	authentication v0.0.0-00010101000000-000000000000
	cookie v0.0.0-00010101000000-000000000000
)
