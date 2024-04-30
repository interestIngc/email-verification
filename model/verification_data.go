package model

import "time"

type VerificationData struct {
	Code      string
	ExpiresAt time.Time
}
