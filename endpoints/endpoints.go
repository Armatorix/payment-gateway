package endpoints

import (
	"net/http"

	"github.com/Armatorix/payment-gateway/db"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	respSuccess = "success"
	respFailed  = "failed"
)

var statusOK = map[string]string{"status": "ok"}

func Healthcheck(c echo.Context) error {
	c.Logger().Info("healthcheck")
	return c.JSON(http.StatusOK, statusOK)
}

type Handler struct {
	DB *db.DB
}

type errorResp struct {
	Status  string
	Message string
}
type authorizeReq struct {
	CreditCardData string `json:"card_data" validate:"required,alphanum"`
	CardNumber     int64  `json:"card_number" validate:"required,luhn"`
}
type authorizeResp struct {
	AuthorizationID uuid.UUID
	Status          string
	AmountAvailable int64
	Currency        string
}

func (h *Handler) Authorize(c echo.Context) error {
	var req authorizeReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}
	creditCard, err := h.DB.SaveCreditCard()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, authorizeResp{
		AuthorizationID: creditCard.ID,
		Status:          respSuccess,
		AmountAvailable: creditCard.Amount,
		Currency:        creditCard.Currency,
	})
}

func (h *Handler) Capture(c echo.Context) error {
	return nil
}
func (h *Handler) Void(c echo.Context) error {
	return nil
}

func (h *Handler) Refund(c echo.Context) error {
	return nil
}
