package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/warrenb95/investment-system/internal/domain/models"
)

func TestListFunds(t *testing.T) {
	tests := map[string]struct {
		storedFunds []models.Fund
	}{
		"successfully list funds from the store": {
			storedFunds: []models.Fund{
				{ID: "065d1a8e-65a1-4d15-8e2f-7551a7f4b574", Name: "test 1", Description: "list test 1"},
				{ID: "88ff2cf3-a004-4537-9de1-18c66278b9c6", Name: "test 2", Description: "list test 1"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, f := range test.storedFunds {
				err := testDB.CreateFund(context.Background(), &f)
				require.NoError(t, err, "creating fund")
			}

			listResp, err := testDB.ListFunds(context.Background())
			require.NoError(t, err, "listing funds")

			assert.ElementsMatch(t, test.storedFunds, listResp, "list response")
		})
	}
}
