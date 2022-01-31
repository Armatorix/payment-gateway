package model

import "github.com/google/uuid"

// TODO models
// type  struct

type State string

const (
	ActiveState = State("active")
	VoidedState = State("voided")
)

type CreditCard struct {
	MerchantID uuid.UUID

	ID          uuid.UUID
	Holder      string
	Number      string
	CVV         int64
	ExpiryMonth int64
	ExpiryYear  int64
	Amount      int64
	Currency    string
	State       State
}

type Merchant struct {
	ID        uuid.UUID
	Name      string
	APISecret string
}
