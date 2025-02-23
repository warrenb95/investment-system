package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/warrenb95/investment-system/internal/domain/models"
	"github.com/warrenb95/investment-system/internal/ports/repository"
)

type InvestmentsService struct {
	logger          *logrus.Logger
	fundStore       repository.FundStore
	investmentStore repository.InvestmentStore
	customerStore   repository.CustomerStore
}

func NewInvestmentsService(
	logger *logrus.Logger,
	fundStore repository.FundStore,
	investmentsStore repository.InvestmentStore,
	customerStore repository.CustomerStore,
) *InvestmentsService {
	return &InvestmentsService{
		logger:          logger,
		fundStore:       fundStore,
		investmentStore: investmentsStore,
		customerStore:   customerStore,
	}
}

func (s *InvestmentsService) LoadFunds(ctx context.Context, reader io.Reader) error {
	logger := s.logger.WithContext(ctx)

	allBytes, err := io.ReadAll(reader)
	if err != nil {
		logger.WithError(err).Error("failed to read all while loading funds")
		return fmt.Errorf("reading all: %w", err)
	}

	var funds []models.Fund
	err = json.Unmarshal(allBytes, &funds)
	if err != nil {
		logger.WithError(err).Error("unmarshalling funds from reader")
		return fmt.Errorf("unmarshalling funds from reader: %w", err)
	}

	for _, f := range funds {
		err := s.fundStore.CreateFund(ctx, &f)
		if err != nil {
			logger.WithError(err).Warn("failed to create fund in store while loading funds from reader")
		}
	}

	return nil
}

func (s *InvestmentsService) ListFunds(ctx context.Context) ([]*models.Fund, error) {
	// Can add business logic here if needed...

	return s.fundStore.ListFunds(ctx)
}

func (s *InvestmentsService) Invest(ctx context.Context, customerID string, investments ...models.Investment) error {
	if len(investments) == 0 {
		s.logger.Error("no investments in invest request")
		return models.ErrEmptyInvestments
	}

	for i, investment := range investments {
		investment.ID = uuid.NewString()
		investment.CustomerID = customerID
		investments[i] = investment
	}

	return s.investmentStore.Invest(ctx, customerID, investments...)
}

func (s *InvestmentsService) ListInvestments(ctx context.Context, customerID string) ([]*models.Investment, error) {
	// Can add business logic here if needed...

	return s.investmentStore.ListInvestments(ctx, customerID)
}
