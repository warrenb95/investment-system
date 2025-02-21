package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	ports "github.com/warrenb95/investment-system/internal/ports/http"
)

func ListFunds(s ports.InvestmentsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		funds, err := s.ListFunds(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to list funds"})
		}

		return c.JSON(http.StatusOK, map[string]any{"funds": funds})
	}
}
