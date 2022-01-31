package endpoints

import (
	"log"
	"net/http"
	"strings"

	"github.com/Armatorix/payment-gateway/db"
	"github.com/Armatorix/payment-gateway/db/middleware"
	"github.com/Armatorix/payment-gateway/model"
	"github.com/Armatorix/payment-gateway/x/xtime"
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

func ErrorRespMiddleware(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := hf(c)
		if err == nil {
			return nil
		}
		log.Println(err)
		switch errMsg := strings.ToLower(err.Error()); {
		case strings.Contains(errMsg, "unauthorized"):
			return c.JSON(http.StatusUnauthorized,
				errorResp{
					Status:  respFailed,
					Message: "basic auth failed",
				})
		case strings.Contains(errMsg, "field validation"):
			return c.JSON(http.StatusBadRequest,
				errorResp{
					Status:  respFailed,
					Message: "fields with improper values",
				})
		default:
			return c.JSON(http.StatusInternalServerError,
				errorResp{
					Status:  respFailed,
					Message: "unexpected error",
				})
		}
	}
}

type authorizeReq struct {
	CardHolder string          `json:"card_holder" validate:"required"`
	CardNumber string          `json:"card_number" validate:"required,luhn"`
	Expiration xtime.MonthYear `json:"card_expiration" validate:"required"`
	CVV        int32           `json:"card_cvv" validate:"required,gte=100,lte=999"`
	Amount     int64           `json:"amount" validate:"required,gte=1"`
	Currency   string          `json:"currency" validate:"required"`
}

type authorizeResp struct {
	AuthorizationID uuid.UUID `json:"authorization_id"`
	Status          string    `json:"status"`
	AmountAvailable int64     `json:"amount_available"`
	Currency        string    `json:"currency"`
}

func (h *Handler) Authorize(c echo.Context) error {
	var req authorizeReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	merchantID := c.Get(middleware.MerchantIDKey).(uuid.UUID)

	creditCard, err := h.DB.SaveCreditCard(c.Request().Context(), model.CreditCard{
		MerchantID:  merchantID,
		ExpiryMonth: req.Expiration.Month,
		ExpiryYear:  req.Expiration.Year,
		CardNumber:  req.CardNumber,
		Holder:      req.CardHolder,
		CVV:         req.CVV,
		Amount:      req.Amount,
		Currency:    req.Currency,
	})
	if err != nil {
		return err
	}
	log.Println(creditCard)
	return c.JSON(http.StatusOK, authorizeResp{
		AuthorizationID: creditCard.ID,
		Status:          respSuccess,
		AmountAvailable: creditCard.Amount,
		Currency:        creditCard.Currency,
	})
}

func (h *Handler) Capture(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

func (h *Handler) Void(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

func (h *Handler) Refund(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
