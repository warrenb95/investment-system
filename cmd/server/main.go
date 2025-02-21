package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/warrenb95/investment-system/internal/adapters/repository"
)

// Database connection setup
func connectDB() *pg.DB {
	opts := &pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "password",
		Database: "railway_db",
	}
	return pg.Connect(opts)
}

func main() {
	db := connectDB()
	defer db.Close()
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))

	_, err := repository.NewPostgresRepository(db, logger)
	if err != nil {
		logger.WithError(err).Fatal("Creating new repository")
	}

	// Echo instance
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logger.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Graceful shutdown
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("Server forced to shutdown")
	}

	logger.Info("Server shutdown")
}
