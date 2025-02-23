package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/warrenb95/investment-system/internal/domain/models"
)

func TestInvest(t *testing.T) {
	tests := map[string]struct {
		storedFunds []models.Fund

		customerID string
		investReq  []models.Investment
	}{
		"successfully store single investment": {
			storedFunds: []models.Fund{
				{ID: "8c6619bd-6c36-47af-9e7f-434c17f410c9", Name: "fund 1", Description: "fund 1"},
			},

			customerID: "5eb8608c-b092-4f2b-8a71-763522e0240f",
			investReq: []models.Investment{
				{
					ID:         "4f148b45-596a-4dcb-a096-60f91e7cad6e",
					CustomerID: "5eb8608c-b092-4f2b-8a71-763522e0240f",
					FundID:     "8c6619bd-6c36-47af-9e7f-434c17f410c9",
					Amount:     10,
				},
			},
		},
		"successfully store multiply investment": {
			storedFunds: []models.Fund{
				{ID: "e37ec08e-d360-4a2a-8571-d1a38905b1f0", Name: "fund 1", Description: "fund 1"},
				{ID: "bd243380-c8ae-4b65-80df-052c265a341a", Name: "fund 2", Description: "fund 2"},
			},

			customerID: "5eb8608c-b092-4f2b-8a71-763522e0240f",
			investReq: []models.Investment{
				{
					ID:         "2c8d8f27-76d4-4e0c-bb64-803d2aca842d",
					CustomerID: "5eb8608c-b092-4f2b-8a71-763522e0240f",
					FundID:     "e37ec08e-d360-4a2a-8571-d1a38905b1f0",
					Amount:     10,
				},
				{
					ID:         "ec604c70-22fd-458c-8875-f183b2ec4874",
					CustomerID: "5eb8608c-b092-4f2b-8a71-763522e0240f",
					FundID:     "e37ec08e-d360-4a2a-8571-d1a38905b1f0",
					Amount:     10,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, f := range test.storedFunds {
				err := testDB.CreateFund(context.Background(), &f)
				require.NoError(t, err, "creating fund")
			}

			err := testDB.CreateCustomer(context.Background(), &models.Customer{ID: test.customerID})
			require.NoError(t, err, "creating customer")

			err = testDB.Invest(context.Background(), test.customerID, test.investReq...)
			require.NoError(t, err, "investing")

			listResp, err := testDB.ListInvestments(context.Background(), test.customerID)
			require.NoError(t, err, "listing funds")

			assert.ElementsMatch(t, test.investReq, listResp, "list response")
		})
	}
}
