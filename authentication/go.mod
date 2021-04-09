module authentication

replace shared => ../shared

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/twinj/uuid v1.0.0
	gopkg.in/stretchr/testify.v1 v1.2.2 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.7
	shared v0.0.0-00010101000000-000000000000
)
