package model

import "github.com/google/uuid"

type ActionType string

const (
	ActiveAT   = ActionType("active")
	VoidedAT   = ActionType("voided")
	RefundedAT = ActionType("refunded")
)

type CreditCard struct {
	MerchantID uuid.UUID

	ID          uuid.UUID
	Holder      string
	CardNumber  string
	CVV         int32
	ExpiryMonth int64
	ExpiryYear  int64
	Amount      int64
	Currency    string
}

type Merchant struct {
	ID uuid.UUID

	Name      string
	APISecret string
}

type CardAction struct {
	CardID uuid.UUID

	ID         uuid.UUID
	ActionType ActionType
	Amount     int64
}
