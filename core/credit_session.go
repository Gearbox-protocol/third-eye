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
		ID                     string       `gorm:"primaryKey" json:"id"`
		Status                 int          `json:"status"`
		Borrower               string       `json:"borrower"`
		CreditManager          string       `json:"creditManager"`
		Account                string       `json:"account"`
		Since                  int64        `json:"since"`
		ClosedAt               int64        `json:"closedAt"`
		InitialAmount          *BigInt      `json:"initialAmount"`
		BorrowedAmount         *BigInt      `json:"borrowedAmount"`
		Balances               *JsonBalance `gorm:"column:balances"`
		CollateralInUSD        *BigInt      `gorm:"column:collateral_usd"`
		RemainingFunds         *BigInt      `gorm:"column:remaining_funds"`
		CollateralInUnderlying *BigInt      `gorm:"column:collateral_underlying"`
		HealthFactor           *BigInt      `gorm:"column:health_factor" json:"healthFactor"`
		IsDirty                bool         `gorm:"-"`
		Liquidator             string       `gorm:"liquidator"`
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
		HealthFactor               *big.Int
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
		ID                     int64        `gorm:"primaryKey;autoincrement:true"`
		BlockNum               int64        `gorm:"column:block_num"`
		SessionId              string       `gorm:"column:session_id"`
		BorrowedAmountBI       *BigInt      `gorm:"column:borrowed_amount_bi"`
		BorrowedAmount         float64      `gorm:"column:borrowed_amount"`
		TotalValueBI           *BigInt      `gorm:"column:total_value_bi"`
		TotalValue             float64      `gorm:"column:total_value"`
		Balances               *JsonBalance `gorm:"column:balances"`
		Borrower               string       `gorm:"column:borrower"`
		CollateralInUSD        *BigInt      `gorm:"column:collateral_in_usd"`
		CollateralInUnderlying *BigInt      `gorm:"qcolumn:collateral_in_underlying"`
		СumulativeIndexAtOpen  *BigInt      `gorm:"column:cumulative_index"`
		HealthFactor           *BigInt      `gorm:"column:health_factor"`
	}
	CreditSessionUpdate struct {
		SessionId        string  `gorm:"column:id;primaryKey"`
		BorrowedAmountBI *BigInt `gorm:"column:borrowed_amount_bi"`
		TotalValueBI     *BigInt `gorm:"column:total_value_bi"`
		Borrower         string  `gorm:"column:borrower"`
		HealthFactor     *BigInt `gorm:"column:health_factor"`
	}
)

func (CreditSessionUpdate) TableName() string {
	return "credit_sessions"
}
