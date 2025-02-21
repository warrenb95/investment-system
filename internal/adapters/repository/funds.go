package repository

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/warrenb95/investment-system/internal/domain/models"
)

func (r *PostgresRepository) CreateFund(ctx context.Context, fund *models.Fund) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(fund).OnConflict("DO NOTHING").Insert()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting fund into pg store")
			return fmt.Errorf("inserting fund: %w", err)
		}

		return nil
	})
}

func (r *PostgresRepository) ListFunds(ctx context.Context) ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Model(&funds).Select()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("listing funds from pg store")
		return nil, fmt.Errorf("listing fund: %w", err)
	}

	return funds, nil
}
