package repository

import "email-verification/model"

type InMemoryCodeStorage struct {
	emailToVerificationData map[string]model.VerificationData
}

func NewInMemoryCodeStorage() *InMemoryCodeStorage {
	cs := new(InMemoryCodeStorage)

	cs.emailToVerificationData = make(map[string]model.VerificationData)

	return cs
}

func (cs *InMemoryCodeStorage) SaveVerificationData(
	email string,
	data model.VerificationData,
) {
	cs.emailToVerificationData[email] = data
}

func (cs *InMemoryCodeStorage) GetVerificationData(
	email string,
) model.VerificationData {
	return cs.emailToVerificationData[email]
}
