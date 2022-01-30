package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var statusOK = map[string]string{"status": "ok"}

func Healthcheck(c echo.Context) error {
	c.Logger().Info("healthcheck")
	return c.JSON(http.StatusOK, statusOK)
}

func Authorize(c echo.Context) error {
	return nil
}

func Capture(c echo.Context) error {
	return nil
}
func Void(c echo.Context) error {
	return nil
}

func Refund(c echo.Context) error {
	return nil
}
