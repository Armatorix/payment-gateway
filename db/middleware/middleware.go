package middleware

import (
	"github.com/Armatorix/payment-gateway/db"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const MerchantIDKey = "merchantId"

func AuthMiddleware(sqlDB *db.DB) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, secret string, c echo.Context) (bool, error) {
		exists, mID, err := sqlDB.VerifySecret(c.Request().Context(), username, secret)
		if mID != uuid.Nil {
			c.Set(MerchantIDKey, mID)
		}
		return exists, err
	})
}
