package models

import "time"

type Fund struct {
	ID          string
	Name        string
	Description string
}

type Customer struct {
	ID string
}

type Investment struct {
	ID         string
	CustomerID string
	FundID     string  `json:"fund_id"`
	Amount     float64 `json:"amount"`
	CreatedAt  time.Time
}
