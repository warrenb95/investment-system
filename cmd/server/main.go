package main

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	httpadapter "github.com/warrenb95/investment-system/internal/adapters/http"
	"github.com/warrenb95/investment-system/internal/adapters/repository"
	"github.com/warrenb95/investment-system/internal/domain/services"
)

//go:embed funds.json
var fundsJSON []byte

// Database connection setup
func connectDB(logger *logrus.Logger) *pg.DB {
	user := os.Getenv("DB_USER")
	if user == "" {
		logger.Fatal("empty DB_USER env var")
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		logger.Fatal("empty DB_PASS env var")
	}
	database := os.Getenv("DB_NAME")
	if database == "" {
		logger.Fatal("empty DB_NAME env var")
	}

	opts := &pg.Options{
		Addr:     "postgres:5432",
		User:     user,
		Password: pass,
		Database: database,
	}
	return pg.Connect(opts)
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))

	db := connectDB(logger)
	defer db.Close()

	pgStore, err := repository.NewPostgresRepository(db, logger)
	if err != nil {
		logger.WithError(err).Fatal("creating new repository")
	}

	s := services.NewInvestmentsService(logger, pgStore, pgStore, pgStore)
	err = s.LoadFunds(context.Background(), bytes.NewReader(fundsJSON))
	if err != nil {
		logger.WithError(err).Fatal("loading funds from funds json")
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

	e.GET("/api/v1/retail/funds", httpadapter.ListFunds(s))

	e.POST("/api/v1/retail/customers", httpadapter.CreateCustomer(s))

	e.POST("/api/v1/retail/isa-investments/:customer_id", httpadapter.Invest(s))
	e.GET("/api/v1/retail/isa-investments/:customer_id", httpadapter.ListInvestments(s))

	e.GET("/api/v1/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Graceful shutdown
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("server forced to shutdown")
	}

	logger.Info("server shutdown")
}
