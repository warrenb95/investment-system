package repository

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/warrenb95/investment-system/internal/domain/models"
)

func (r *PostgresRepository) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(customer).OnConflict("DO NOTHING").Insert() // TODO: add & return error if collision.
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting customer into pg store")
			return fmt.Errorf("inserting customer: %w", err)
		}

		return nil
	})
}

func (r *PostgresRepository) ListCustomer(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Model(&customers).Select()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("listing customers from pg store")
		return nil, fmt.Errorf("listing customers: %w", err)
	}

	return customers, nil
}
