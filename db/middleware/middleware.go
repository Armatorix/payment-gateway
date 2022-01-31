package middleware

import (
	"github.com/Armatorix/payment-gateway/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthMiddleware(sqlDB *db.DB) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, secret string, c echo.Context) (bool, error) {
		return sqlDB.VerifySecret(username, secret)
	})
}
