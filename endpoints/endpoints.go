package endpoints

import (
	"net/http"

	"github.com/Armatorix/payment-gateway/db"
	"github.com/Armatorix/payment-gateway/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	respSuccess = "success"
	respFailed  = "failure"
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

func ErrorRespMiddleware() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := hf(c)
			if err == nil {
				return nil
			}
			switch {
			// TODO: define all specific cases with messages
			default:
				return c.JSON(http.StatusInternalServerError,
					errorResp{
						Status:  respFailed,
						Message: "unexpected error",
					})
			}
		}
	}
}

type authorizeReq struct {
	CardHolder string `json:"card_holder" validate:"required,alphanum"`
	CardNumber string `json:"card_number" validate:"required,luhn"`
	Expiration struct {
		Year  int64 `json:"year"`
		Month int64 `json:"month"`
	} `json:"card_expiration" validate:"required,not_expired"`
	CVV int64 `json:"card_cvv" validate:"required,gte=100,lte=999"`
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
	creditCard, err := h.DB.SaveCreditCard(model.CreditCard{
		ExpiryMonth: req.Expiration.Month,
		ExpiryYear:  req.Expiration.Year,
		Number:      req.CardNumber,
		Holder:      req.CardHolder,
		CVV:         req.CVV,
	})
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
	// TODO
	return nil
}
func (h *Handler) Void(c echo.Context) error {
	// TODO
	return nil
}

func (h *Handler) Refund(c echo.Context) error {
	// TODO
	return nil
}
