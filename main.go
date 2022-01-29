package main

import (
	"fmt"
	"log"

	"github.com/Armatorix/payment-gateway/config"
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
		log.Fatalf("failed env init: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Logger.SetLevel(cfg.Server.LogLevel)
	e.Validator = &V{validator.New()}
	e.Use(middleware.CORS())
	e.GET("/public/health-check", endpoints.Healthcheck)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Server.Port)).Error())
}
