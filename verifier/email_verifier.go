package verifier

import (
	"email-verification/model"
	"email-verification/repository"
	"email-verification/senderservice"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const verificationCodeSize = 6
const verificationCodeValidityDuration = time.Minute * 2

type EmailVerifier struct {
	adminEmailAddress       model.EmailAddress
	emailSenderService      senderservice.EmailSenderService
	verificationCodeStorage repository.VerificationCodeStorage
}

func NewEmailVerifier(
	adminEmailAddress model.EmailAddress,
	emailSenderService senderservice.EmailSenderService,
	verificationCodeStorage repository.VerificationCodeStorage,
) *EmailVerifier {
	ev := new(EmailVerifier)
	ev.adminEmailAddress = adminEmailAddress
	ev.emailSenderService = emailSenderService
	ev.verificationCodeStorage = verificationCodeStorage
	return ev
}

func getVerificationCode(userEmail model.EmailAddress) string {
	verificationCode := make([]string, verificationCodeSize)
	for i := 0; i < verificationCodeSize; i++ {
		verificationCode[i] = strconv.Itoa(rand.Intn(10))
	}
	return strings.Join(verificationCode, "")
}

func (ev *EmailVerifier) InitVerification(userEmail model.EmailAddress) {
	code := getVerificationCode(userEmail)
	expiresAt := time.Now().Add(verificationCodeValidityDuration)

	ev.verificationCodeStorage.SaveVerificationData(
		string(userEmail),
		model.VerificationData{
			Code:      code,
			ExpiresAt: expiresAt,
		},
	)

	emailModel := model.Email{
		From:    ev.adminEmailAddress,
		To:      userEmail,
		Subject: "Verify your email at the Layer8 service",
		Content: model.EmailContentData{
			Username: generateUsername(userEmail),
			Code:     code,
		},
		//Content: fmt.Sprintf("Dear user, your verification code for Layer8 is %s", code),
	}

	err := ev.emailSenderService.SendEmail(emailModel)
	if err != nil {
		log.Fatalf("an error while sending a verification email occurred: %e", err)
	}
}

func (ev *EmailVerifier) VerifyCode(email model.EmailAddress, code string) error {
	verificationData := ev.verificationCodeStorage.GetVerificationData(string(email))

	if verificationData.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf(
			"the verification code is expired. Please try to run the verification process again",
		)
	}

	if code != verificationData.Code {
		return fmt.Errorf(
			"invalid verification code, expected %s, got %s",
			verificationData.Code,
			code,
		)
	}

	return nil
}

func generateUsername(email model.EmailAddress) string {
	return strings.Split(string(email), "@")[0]
}
