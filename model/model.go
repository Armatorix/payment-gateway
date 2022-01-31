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
	ID          uuid.UUID
	Data        string
	Number      string
	CVV         int64
	ExpiryMonth int64
	ExpiryYear  int64
	Amount      int64
	Currency    string
	State       State
}
