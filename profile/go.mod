module profile

go 1.15

replace models => ../models

replace authentication => ../authentication

require (
	authentication v0.0.0-00010101000000-000000000000
	github.com/lamoda/gonkey v1.1.0
	github.com/stretchr/testify v1.6.1
	models v0.0.0-00010101000000-000000000000
)
