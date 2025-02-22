package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/warrenb95/investment-system/internal/domain/models"
)

func (s *InvestmentsService) CreateCustomer(ctx context.Context) (string, error) {
	customerID := uuid.NewString()
	// TODO: could add retry logic here if there's an ID colision.
	return customerID, s.customerStore.CreateCustomer(ctx, &models.Customer{ID: customerID})
}
