package auth

import (
	"log"
	//"notifier/mailer"
	"os"

	"shared/amqp/events"
	"shared/amqp/receiver"
	"shared/amqp/sender"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial(os.Getenv(""))
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}
	defer connection.Close()

	//Ensure Exchange exist my creating an emitter
	_ , err = sender.NewAmqpEventEmitter(connection, "NaeraExchange")
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}

	authLister, err := receiver.NewEventEventListener(connection, "NaeraAuth")
	go ProcessEvents(authLister)
	c := make(chan int)
	<-c
}

func ProcessEvents(eventListener events.EventListener) error {
	received, errors, err := eventListener.Listen("NaeraExchange", "user.created")
	if err != nil {
		log.Fatalf("event listenner error %v", err.Error())
	}
	for {
		select {
		case evt := <-received:
			// log
			switch e := evt.(type) {
			case *events.UserCreatedEvent:
				// Send sign up email
				//mailer.SendMail()

			default:
				log.Printf("unknown event: %t", e)
			}
		case err = <-errors:
			log.Printf(" received error while processing msg: %s", err)
		}
	}
}
