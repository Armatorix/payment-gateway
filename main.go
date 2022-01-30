package main

import (
	"fmt"
	"log"

	"github.com/Armatorix/payment-gateway/config"
	"github.com/Armatorix/payment-gateway/db"
	dbmiddleware "github.com/Armatorix/payment-gateway/db/middleware"
	"github.com/Armatorix/payment-gateway/endpoints"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type V struct {
	Validator *validator.Validate
}

func (cv *V) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Logger.SetLevel(cfg.Server.LogLevel)
	e.Validator = &V{validator.New()}
	e.Use(middleware.CORS())
	e.GET("/public/health-check", endpoints.Healthcheck)
	sqlDB := db.Connect(cfg.DB.DSN)
	withTransaction := e.Group("",
		dbmiddleware.AuthMiddleware(sqlDB),
		dbmiddleware.SerializableTxMiddleware(sqlDB))

	withTransaction.POST("/authorize", endpoints.Authorize)
	withTransaction.POST("/capture", endpoints.Capture)
	withTransaction.POST("/void", endpoints.Void)
	withTransaction.POST("/refund", endpoints.Refund)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Server.Port)).Error())
}
