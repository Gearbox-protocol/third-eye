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
		ID                     string       `gorm:"primaryKey" json:"sessionId"`
		Status                 int          `json:"status"`
		Borrower               string       `json:"borrower"`
		CreditManager          string       `json:"creditManager"`
		Account                string       `json:"account"`
		Since                  int64        `json:"since"`
		ClosedAt               int64        `json:"closedAt"`
		InitialAmount          *BigInt      `json:"initialAmount"`
		BorrowedAmount         *BigInt      `json:"borrowedAmount"`
		Balances               *JsonBalance `gorm:"column:balances"`
		RemainingFunds         *BigInt      `gorm:"column:remaining_funds"`
		CollateralInUSD        float64      `gorm:"<-:false;column:collateral_usd"`
		CollateralInUnderlying float64      `gorm:"<-:false;column:collateral_underlying"`
		HealthFactor           *BigInt      `gorm:"column:health_factor" json:"healthFactor"`
		IsDirty                bool         `gorm:"-"`
		Liquidator             string       `gorm:"liquidator"`
		Version                int16        `gorm:"version"`
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
		ID                     int64        `gorm:"primaryKey;autoincrement:true" json:"-"`
		BlockNum               int64        `gorm:"column:block_num" json:"blockNum"`
		SessionId              string       `gorm:"column:session_id" json:"sessionId"`
		BorrowedAmountBI       *BigInt      `gorm:"column:borrowed_amount_bi" json:"borrowedAmountBI"`
		BorrowedAmount         float64      `gorm:"column:borrowed_amount" json:"borrowedAmount"`
		TotalValueBI           *BigInt      `gorm:"column:total_value_bi" json:"totalValueBI"`
		TotalValue             float64      `gorm:"column:total_value" json:"totalValue"`
		Balances               *JsonBalance `gorm:"column:balances" json:"balance"`
		Borrower               string       `gorm:"column:borrower" json:"borrower"`
		CollateralInUSD        float64      `gorm:"column:collateral_usd" json:"collateralInUSD"`
		CollateralInUnderlying float64      `gorm:"column:collateral_underlying" json:"collateralInUnderlying"`
		СumulativeIndexAtOpen  *BigInt      `gorm:"column:cumulative_index" json:"cumulativeIndexAtOpen"`
		HealthFactor           *BigInt      `gorm:"column:health_factor" json:"healthFactor"`
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
