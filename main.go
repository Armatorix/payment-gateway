package main

import (
	"fmt"
	"log"

	"github.com/Armatorix/payment-gateway/config"
	"github.com/Armatorix/payment-gateway/db"
	dbmiddleware "github.com/Armatorix/payment-gateway/db/middleware"
	"github.com/Armatorix/payment-gateway/endpoints"
	"github.com/Armatorix/payment-gateway/validator"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	prometheus.NewPrometheus("payment-gateway", nil).Use(e)
	e.Use(middleware.Logger())
	e.Logger.SetLevel(cfg.Server.LogLevel)

	if e.Validator, err = validator.New(); err != nil {
		log.Fatal(err)
	}

	e.Use(endpoints.ErrorRespMiddleware,
		middleware.CORS())
	e.GET("/public/health-check", endpoints.Healthcheck)

	sqlDB := db.Connect(cfg.DB.DSN)
	if err = sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	authed := e.Group("", dbmiddleware.AuthMiddleware(sqlDB))

	h := &endpoints.Handler{DB: sqlDB}
	authed.POST("/authorize", h.Authorize)
	authed.POST("/capture", h.Capture)
	authed.POST("/void", h.Void)
	authed.POST("/refund", h.Refund)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Server.Port)).Error())
}
