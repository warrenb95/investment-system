package repository

import (
	"context"

	"github.com/warrenb95/investment-system/internal/domain/models"
)

type FundStore interface {
	CreateFund(ctx context.Context, fund *models.Fund) error
	ListFunds(ctx context.Context) ([]models.Fund, error)
}

type InvestmentStore interface {
	CreateInvestment(ctx context.Context, customerID string, investments ...*models.Investment) error
}
