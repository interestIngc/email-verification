package main

import (
	"email-verification/input"
	"email-verification/model"
	"email-verification/repository"
	"email-verification/senderservice"
	"email-verification/verifier"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	configFile, e := os.Open("config.json")
	if e != nil {
		log.Fatalf("Could not open the config file to read from: %e", e)
	}
	configBytes, e := io.ReadAll(configFile)
	if e != nil {
		log.Fatalf("Error while reading from the config file: %e", e)
	}

	var config input.Config
	e = json.Unmarshal(configBytes, &config)
	if e != nil {
		log.Fatalf("Error while unmarshalling bytes from the config json file: %e", e)
	}

	adminEmail := model.EmailAddress(fmt.Sprintf("layer8@%s", config.TestDomain))
	userEmail := model.EmailAddress("iamveronikanikina@gmail.com")

	emailSenderService := senderservice.NewMailerSendService(
		config.ApiKey,
		config.TemplateId,
	)
	verificationCodeStorage := repository.NewInMemoryCodeStorage()

	emailVerifier := verifier.NewEmailVerifier(
		adminEmail,
		emailSenderService,
		verificationCodeStorage,
	)

	fmt.Printf("Starting verification of email address %s\n", userEmail)
	fmt.Println()

	emailVerifier.InitVerification(userEmail)

	fmt.Println()
	fmt.Println("Please input the verification code received:")

	var code string
	_, err := fmt.Scanln(&code)
	if err != nil {
		log.Fatalf("Error while reading the input verification code occurred, %e", err)
	}

	err = emailVerifier.VerifyCode(userEmail, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations! Your email address was successfully verified!")
}
