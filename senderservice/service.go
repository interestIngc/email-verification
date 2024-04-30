package senderservice

import (
	"email-verification/model"
)

type EmailSenderService interface {
	SendEmail(email model.Email) error
}
