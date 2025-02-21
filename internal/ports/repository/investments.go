package repository

import (
	"context"

	"github.com/warrenb95/investment-system/internal/domain/models"
)

type FundStore interface {
	CreateFund(ctx context.Context, fund *models.Fund) error
	ListFunds(ctx context.Context) ([]models.Fund, error)
}
