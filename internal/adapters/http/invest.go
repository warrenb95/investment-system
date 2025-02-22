package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/warrenb95/investment-system/internal/domain/models"
	ports "github.com/warrenb95/investment-system/internal/ports/http"
)

type investRequest struct {
	Investments []*models.Investment
}

func Invest(s ports.InvestmentsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID := c.Param("customer_id")
		if customerID == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request", "details": "customer_id is missing"})
		}

		var req investRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload", "details": err.Error()})
		}

		err := s.Invest(context.Background(), customerID, req.Investments...)
		if err != nil {
			if errors.Is(err, models.ErrEmptyInvestments) {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request", "details": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to invest", "details": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]any{"message": "Investments successfully made."})
	}
}

func ListInvestments(s ports.InvestmentsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID := c.Param("customer_id")
		if customerID == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request", "details": "customer_id is missing"})
		}

		investments, err := s.ListInvestments(c.Request().Context(), customerID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to list investments"})
		}

		return c.JSON(http.StatusOK, map[string]any{"investments": investments})
	}
}
