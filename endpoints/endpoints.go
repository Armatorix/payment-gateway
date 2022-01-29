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
