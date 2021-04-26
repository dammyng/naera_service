package receiver

import (
	"encoding/json"
	"fmt"
	"shared/amqp/events"

	"github.com/streadway/amqp"
)

type amqpEventListener struct {
	connection *amqp.Connection
	queue      string
}

func (a *amqpEventListener) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
	return err
}

func NewEventEventListener(conn *amqp.Connection, queue string) (events.EventListener, error) {
	listener := &amqpEventListener{
		connection: conn,
		queue:      queue,
	}

	err := listener.setup()
	if err != nil {
		return nil, err
	}
	return listener, err
}

func (a *amqpEventListener) Listen(exchange string, eventNames ...string) (<-chan events.Event, <-chan error, error) {

	channel, err := a.connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	for _, eventName := range eventNames {
		if err := channel.QueueBind(a.queue, eventName, exchange, false, nil); err != nil {
			return nil, nil, err
		}
	}

	// Consume message from channel - in out case all queues bonded to it as above
	msgs, err := channel.Consume(a.queue, "", false, false, true, false, nil)

	if err != nil {
		return nil, nil, err
	}

	cevents := make(chan events.Event)
	errors := make(chan error)

	go func() {
		for msg := range msgs {
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errors <- fmt.Errorf("msg did not contain a header name")
				msg.Nack(false, false)
				continue
			}

			eventName, ok := rawEventName.(string)

			if !ok {
				errors <- fmt.Errorf("x-event-name header is not string but %T", rawEventName)
				msg.Nack(false, false)
				continue
			}

			var event events.Event
			switch eventName {
			case "user.created":
				event = new(events.UserCreatedEvent)
			case "user.passwordresetrequest":
				event = new(events.PasswordResetRequest)
			case "user.resendemailvalidation":
				event = new(events.ResendEmailEvent)
			default:
				errors <- fmt.Errorf("event type %s i unknown", eventName)
				continue
			}
			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errors <- err
				continue
			}
			cevents <- event
			msg.Ack(false)

		}
	}()
	return cevents, errors, nil
}
