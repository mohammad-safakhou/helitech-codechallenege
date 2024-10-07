package rest

import (
	"codechallenge/config"
	"codechallenge/logger"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func Start() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	registerRoutes(e)

	go func() {
		if err := e.Start(config.AppConfig.General.Host + ":" + config.AppConfig.General.Listen); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Fatal("shutting down server: " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of X seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := e.Shutdown(context.TODO()); err != nil {
		logger.Logger.Fatal(err)
	}
}
