package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"notifier/billoperations"
	"notifier/models"
	"strings"
	"time"
	"github.com/twinj/uuid"

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

	billsListener, err := receiver.NewEventEventListener(connection, "xx")
	if err != nil {
		log.Fatalf("receiver listenner error %v", err.Error())
	}
	go ProcessEvents(billsListener)
	c := make(chan int)
	<-c
}

func ProcessEvents(eventListener events.EventListener) error {
	JWTkey := "CreWwdvOO3pclP3ZFqZUbsDZYL0HyWoU"
	received, errors, err := eventListener.Listen("NaeraExchange", "buy.airtime")
	if err != nil {
		log.Fatalf("event listenner error %v", err.Error())
	}
	for {

		select {

		case evt := <-received:
			log.Printf("got event %s ", evt.EventName())

			switch e := evt.(type) {
			case *events.ServiceAirTimeEvent:
				request := models.ServiceRequestPayload{
					Country:    "NG",
					Customer:   e.Phone,
					Amount:     e.Amount,
					Recurrence: "ONCE",
					Type:       "AIRTIME",
					Reference:  e.ID,
				}
				_request, _ := json.Marshal(&request)

				_, err := billoperations.ServiceTransaction(string(_request))
				order := &billoperations.Order{
					TransactionId: e.Transaction,
					CreatedAt:     time.Now().Unix(),
					Amount:        float32(e.Amount),
					Id:            uuid.NewV4().String(),
					Title:         "Buy Airtime" + " " + e.Phone,
					Charged:       true,
					Fulfilled:     true,
				}
				if err != nil {
					order.Fulfilled = false
				}
				_body, _ := json.Marshal(&order)
				body := string(_body)
				orderCreateURL, _ := url.Parse(e.OrderURL)

				createOrderReq := &http.Request{
					Method: "POST",
					URL:    orderCreateURL,
					Header: map[string][]string{
						"Content-Type":  {"application/json"},
						"Authorization": {"Bearer " + JWTkey},
					},
					Body: ioutil.NopCloser(strings.NewReader(body)),
				}
				_, err = billoperations.HttpReq(createOrderReq)
				if err != nil {
					log.Println("")
				}
			default:
				log.Printf("unknown event: %t", e)
			}
		case err = <-errors:
			log.Printf(" received error while processing msg: %s", err)
		}

	}
}
