package core

import (
	"math/big"
)

const (
	Active = iota
	Closed
	Repaid
	Liquidated
)

type (
	CreditSession struct {
		ID               string            `gorm:"primaryKey" json:"id"`
		Name             string            `gorm:"column:name"`
		Background       string            `gorm:"column:background"`
		Status           int               `json:"status"`
		Borrower         string            `json:"borrower"`
		CreditManager    string            `json:"creditManager"`
		Account          string            `json:"account"`
		Since            int64             `json:"since"`
		ClosedAt         int64             `json:"closedAt"`
		InitialAmount    *BigInt           `json:"initialAmount"`
		BorrowedAmount   *BigInt           `json:"borrowedAmount"`
		Profit           *BigInt           `json:"profit"`
		ProfitPercentage float64           `gorm:"column:profit_percent" json:"profitPercentage"`
		TotalValue       *BigInt           `gorm:"column:total_value" json:"totalValue"`
		HealthFactor     int64               `gorm:"column:health_factor" json:"healthFactor"`
		Score            float64           `json:"score"`
	}

	CreditAccountData struct {
		Address                    string
		Borrower                   string
		InUse                      bool
		CreditManager              string
		Kind                       string
		UnderlyingToken            string
		BorrowedAmountPlusInterest *big.Int
		TotalValue                 *big.Int
		HealthFactor               int64
		BorrowRate                 *big.Int
	}

	CreditAccountDataExtended struct {
		CreditAccountData
		RepayAmount           *big.Int
		LiquidationAmount     *big.Int
		BorrowedAmount        *big.Int
		СumulativeIndexAtOpen *big.Int
		Since                 int64
	}
	CreditSessionSnapshot struct {
		ID               int64            `gorm:"primaryKey" json:"id"`
		BlockNum         int64            `gorm:"column:block_num"`
		SessionId        string           `gorm:"column:session_id"`
		BorrowedAmountBI *BigInt          `gorm:"column:borrowed_amount_bi"`
		BorrowedAmount   float64          `gorm:"column:borrowed_amount"`
		TotalValueBI     *BigInt          `gorm:"column:total_value_bi"`
		TotalValue       float64          `gorm:"column:total_value"`
		Balances         Hstore           `gorm:"column:balances"`
	}
)
