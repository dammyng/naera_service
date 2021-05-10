package service

import (
	"log"

	"shared/amqp/events"
	"shared/amqp/receiver"
	"shared/amqp/sender"

	"github.com/streadway/amqp"
)

func StartServiceProcessListener(AMQP_HOST, Exchange, Queue string) {
	connection, err := amqp.Dial(AMQP_HOST)
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}
	defer connection.Close()

	//Ensure Exchange exist my creating an emitter
	_, err = sender.NewAmqpEventEmitter(connection, Exchange)
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}

	billsListener, err := receiver.NewEventEventListener(connection, Queue)
	go ProcessEvents(billsListener)
	c := make(chan int)
	<-c
}

func ProcessEvents(eventListener events.EventListener) error {
	received, errors, err := eventListener.Listen("NaeraExchange", "buy.airtime")
	if err != nil {
		log.Fatalf("event listenner error %v", err.Error())
	}
	for {

		select {

		case evt := <-received:
			log.Printf("got event %s ", evt.EventName())

			// log
			switch e := evt.(type) {
			case *events.ServiceAirTimeEvent:
			default:
				log.Printf("unknown event: %t", e)
			}
		case err = <-errors:
			log.Printf(" received error while processing msg: %s", err)
		}

	}
}
