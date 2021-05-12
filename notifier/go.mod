module notifier

replace shared => ../shared

go 1.14

require (
	github.com/joho/godotenv v1.3.0
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/sendgrid/rest v2.6.3+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.8.0+incompatible
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/twinj/uuid v1.0.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	gopkg.in/stretchr/testify.v1 v1.2.2 // indirect
	shared v0.0.0-00010101000000-000000000000
)
