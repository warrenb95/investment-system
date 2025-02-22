package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	ports "github.com/warrenb95/investment-system/internal/ports/http"
)

func CreateCustomer(s ports.CustomerService) echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID, err := s.CreateCustomer(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create customer"})
		}

		return c.JSON(http.StatusCreated, map[string]string{"id": customerID})
	}
}
