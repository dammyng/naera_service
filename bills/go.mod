module bills

replace shared => ../shared

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/twinj/uuid v1.0.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/stretchr/testify.v1 v1.2.2
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.8
	shared v0.0.0-00010101000000-000000000000

)
