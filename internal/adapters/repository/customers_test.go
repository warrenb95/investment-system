package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/warrenb95/investment-system/internal/domain/models"
)

func TestCreateCustomer(t *testing.T) {
	tests := map[string]struct {
		customerReq *models.Customer
	}{
		"successfully create customer": {
			customerReq: &models.Customer{ID: "f69c5605-e1db-4bae-ac39-6704a6fb9556"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := testDB.CreateCustomer(context.Background(), test.customerReq)
			require.NoError(t, err, "creating customer")

			customersResp, err := testDB.ListCustomer(context.Background())
			require.NoError(t, err, "listing customers")

			require.Len(t, customersResp, 1, "customers list response")
			assert.Equal(t, test.customerReq.ID, customersResp[0].ID, "customer ID")
		})
	}
}
