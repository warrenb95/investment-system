package http

import (
	"context"

	"github.com/warrenb95/investment-system/internal/domain/models"
)

type InvestmentsService interface {
	ListFunds(ctx context.Context) ([]models.Fund, error)
	Invest(ctx context.Context, customerID string, investments ...*models.Investment) error
}
