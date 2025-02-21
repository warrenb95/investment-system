package repository

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/warrenb95/investment-system/internal/domain/models"
)

func (r *PostgresRepository) Invest(ctx context.Context, customerID string, investments ...models.Investment) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(&models.Customer{ID: customerID}).OnConflict("DO NOTHING").Insert()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting customer into pg store")
			return fmt.Errorf("inserting customer: %w", err)
		}

		_, err = r.db.Model(&investments).Table("isa_investments").Insert()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting investments into pg store")
			return fmt.Errorf("inserting investments: %w", err)
		}

		return nil
	})
}

func (r *PostgresRepository) ListInvestments(ctx context.Context, customerID string) ([]models.Investment, error) {
	var investments []models.Investment
	return investments, r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		err := r.db.Model(&investments).Table("isa_investments").Where("customer_id = ?", customerID).Select()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("listing customer investments from pg store")
			return fmt.Errorf("listing investments: %w", err)
		}

		return nil
	})
}
