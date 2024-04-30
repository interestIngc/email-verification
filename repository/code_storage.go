package repository

import "email-verification/model"

type VerificationCodeStorage interface {
	SaveVerificationData(email string, data model.VerificationData)
	GetVerificationData(email string) model.VerificationData
}
