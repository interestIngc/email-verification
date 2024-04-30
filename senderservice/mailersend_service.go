package senderservice

import (
	"context"
	"email-verification/model"
	"fmt"
	"github.com/mailersend/mailersend-go"
	"time"
)

type MailerSendService struct {
	apiKey     string
	templateId string
}

func NewMailerSendService(apiKey string, templateId string) *MailerSendService {
	service := new(MailerSendService)

	service.apiKey = apiKey
	service.templateId = templateId

	return service
}

func (service *MailerSendService) SendEmail(email model.Email) error {
	client := mailersend.NewMailersend(service.apiKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	username := email.Content.Username

	from := mailersend.From{
		Name:  "Layer8",
		Email: string(email.From),
	}

	to := mailersend.Recipient{
		Name:  username,
		Email: string(email.To),
	}

	personalization := []mailersend.Personalization{
		{
			Email: string(email.To),
			Data: map[string]interface{}{
				"code": email.Content.Code,
				"user": username,
			},
		},
	}

	message := client.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients([]mailersend.Recipient{to})
	message.SetSubject(email.Subject)
	message.SetTemplateID(service.templateId)
	message.SetPersonalization(personalization)

	fmt.Println("MailerSendService:")
	fmt.Println("Sending an API request to MailerSend")

	response, err := client.Email.Send(ctx, message)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Response from the provider received. Status code %d\n",
		response.StatusCode,
	)

	return nil
}
