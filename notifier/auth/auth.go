package auth

import (
	"bytes"
	"fmt"
	"log"
	"notifier/mailer"
	"os"
	"text/template"

	"shared/amqp/events"
	"shared/amqp/receiver"
	"shared/amqp/sender"

	"github.com/streadway/amqp"
)

func StartAuthenticationListener(AMQP_HOST, Exchange, Queue string) {
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

	authLister, err := receiver.NewEventEventListener(connection, Queue)
	go ProcessEvents(authLister)
	c := make(chan int)
	<-c
}

func ProcessEvents(eventListener events.EventListener) error {
	received, errors, err := eventListener.Listen("NaeraExchange", "user.created", "user.resendemailvalidation", "user.passwordresetrequest")
	if err != nil {
		log.Fatalf("event listenner error %v", err.Error())
	}
	for {

		select {

		case evt := <-received:
			log.Printf("got event %s ", evt.EventName())

			// log
			switch e := evt.(type) {
			case *events.UserCreatedEvent:
				//Send sign up email
				log.Println("New user Mail")
				subject := "Your Naera Pay Verification"
				textContent := fmt.Sprintf("You're on your way! Your email verification code is %s", e.Token)
				t := template.Must(template.New("email_confirm").Parse(`
				You're on your way! Your email verification code is {{.Token}}`))
				out := new(bytes.Buffer)
				data := struct {
					Token string
				}{
					e.Token,
				}
				err = t.Execute(out, data)
				if err != nil {
					log.Fatal(err)
				}

				htmlBytes := out.Bytes()
				htmlContent := string(htmlBytes)
				msg := mailer.EmailMessage{
					subject,
					e.Email,
					textContent,
					htmlContent,
				}

				mailer.SendMail(msg, os.Getenv("AlphaAdmin"), os.Getenv("SendGridKey"))
			case *events.ResendEmailEvent:
				//Send sign up email
				log.Println("Resend Email")
				subject := "Your Naera Pay Verification"
				textContent := fmt.Sprintf("You're on your way! Your email verification code is %s", e.Token)
				t := template.Must(template.New("email_confirm").Parse(`
					You're on your way! Your email verification code is {{.Token}}`))
				out := new(bytes.Buffer)
				data := struct {
					Token string
				}{
					e.Token,
				}
				err = t.Execute(out, data)
				if err != nil {
					log.Fatal(err)
				}

				htmlBytes := out.Bytes()
				htmlContent := string(htmlBytes)
				msg := mailer.EmailMessage{
					subject,
					e.Email,
					textContent,
					htmlContent,
				}

				mailer.SendMail(msg, os.Getenv("AlphaAdmin"), os.Getenv("SendGridKey"))

			case *events.PasswordResetRequest:
				//Send sign up email
				log.Println("New user Mail")
				subject := "Naera pay password reset"
				textContent := fmt.Sprintf("You recently request to reset your naera pay password. Continue with this token: %s", e.Token)
				t := template.Must(template.New("password_reset").Parse(`
				You recently request to reset your naera pay password. Continue with this token: {{.Token}}`))
				out := new(bytes.Buffer)
				data := struct {
					Token string
				}{
					e.Token,
				}
				err = t.Execute(out, data)
				if err != nil {
					log.Fatal(err)
				}

				htmlBytes := out.Bytes()
				htmlContent := string(htmlBytes)
				msg := mailer.EmailMessage{
					subject,
					e.Email,
					textContent,
					htmlContent,
				}

				mailer.SendMail(msg, os.Getenv("AlphaAdmin"), os.Getenv("SendGridKey"))
			default:
				log.Printf("unknown event: %t", e)
			}
		case err = <-errors:
			log.Printf(" received error while processing msg: %s", err)
		}
	}
}
