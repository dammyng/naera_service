module notifier

replace shared => ../shared

go 1.14

require (
	github.com/sendgrid/rest v2.6.3+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.8.0+incompatible
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	shared v0.0.0-00010101000000-000000000000
)
