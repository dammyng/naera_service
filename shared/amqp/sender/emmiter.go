package sender

import (
	"encoding/json"
	"naera_shared/amqp/events"

	"github.com/streadway/amqp"
)

type EventEmitter interface {
	Emit(event events.Event, exchange string) error
}

type amqpEventEmitter struct {
	conn *amqp.Connection
}

func NewAmqpEventEmitter(conn *amqp.Connection, exchange string) (EventEmitter, error) {
	emitter := &amqpEventEmitter{
		conn: conn,
	}

	err := emitter.declearExchange(exchange)
	if err != nil {
		return nil, err
	}
	return emitter, nil
}

func (ee *amqpEventEmitter) Emit(event events.Event, exchange string) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	channel, err := ee.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	msg := amqp.Publishing{
		Headers: amqp.Table{
			"x-event-name": event.EventName(),
		},
		ContentType: "application/json",
		Body:        jsonData,
	}
	err = channel.Publish(
		exchange,          // exchange
		event.EventName(), // routing key
		false,             // mandatory
		false,             // immediate
		msg,
	)
	return err
}

func (ee *amqpEventEmitter) declearExchange(exchange string) error {
	channel, err := ee.conn.Channel()

	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	return err
}
