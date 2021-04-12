package mailer

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailMessage struct {
	Subject string
	To      string
	Text    string
	Html    string
}

func SendMail(msg EmailMessage, sender, key string) {
	log.Println("send mail called")

	from := mail.NewEmail("Naera", sender)
	subject := msg.Subject
	to := mail.NewEmail("Recipient", msg.To)
	plainTextContent := msg.Text
	htmlContent := msg.Html

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(message)
	
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}