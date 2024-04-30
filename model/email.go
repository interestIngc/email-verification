package model

type EmailAddress string

type Email struct {
	From    EmailAddress
	To      EmailAddress
	Subject string
	Content EmailContentData
}

type EmailContentData struct {
	Username string
	Code     string
}
