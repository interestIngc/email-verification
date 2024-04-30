package senderservice

import (
	"email-verification/model"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

type SendGridEmailSenderService struct{}

func (service *SendGridEmailSenderService) SendEmail(email model.Email) error {
	from := mail.NewEmail("Layer8", string(email.From))
	to := mail.NewEmail("User", string(email.To))
	sendgridEmail := mail.NewSingleEmailPlainText(from, email.Subject, to, "")

	client := sendgrid.NewSendClient(os.Getenv("twillio_key"))
	response, err := client.Send(sendgridEmail)
	if err != nil {
		return err
	}

	fmt.Printf("Email sent. Response from the API endpoint: %d", response.StatusCode)
	return nil
}
