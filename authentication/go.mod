module authentication

replace shared => ../shared

go 1.14

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/copier v0.2.9
	github.com/joho/godotenv v1.3.0
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/twinj/uuid v1.0.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/stretchr/testify.v1 v1.2.2 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.7
	shared v0.0.0-00010101000000-000000000000
)
