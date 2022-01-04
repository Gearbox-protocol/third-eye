package core

import (
	"math/big"
)

func (CreditManagerState) TableName() string {
	return "credit_managers"
}

type CreditManagerState struct {
	CreditManagerData
	Address           string            `gorm:"primaryKey"`
	IsWETH            bool              `gorm:"is_weth"`
	PoolAddress       string            `gorm:"column:pool_address"`
	UnderlyingToken   string            `gorm:"column:underlying_token"`
	MaxLeverageFactor int64             `gorm:"column:max_leverage"`
	MinAmount         *BigInt           `gorm:"column:min_amount"`
	MaxAmount         *BigInt           `gorm:"column:max_amount"`
	FeeInterest       int64             `gorm:"column:fee_interest"`
	Sessions          map[string]string `gorm:"-"`
}

type CreditManagerData struct {
	BorrowRateBI            *BigInt `gorm:"column:borrow_rate_bi" `
	BorrowRate              float64 `gorm:"column:borrow_rate"`
	AvailableLiquidityBI    *BigInt `gorm:"column:available_liquidity_bi"`
	AvailableLiquidity      float64 `gorm:"column:available_liquidity"`
	OpenedAccountsCount     int     `gorm:"column:opened_accounts_count"`
	TotalOpenedAccounts     int     `gorm:"column:total_opened_accounts"`
	TotalClosedAccounts     int     `gorm:"column:total_closed_accounts"`
	TotalRepaidAccounts     int     `gorm:"column:total_repaid_accounts"`
	TotalLiquidatedAccounts int     `gorm:"column:total_liquidated_accounts"`
	UniqueUsers             int     `gorm:"column:unique_users"`
	TotalBorrowed           float64 `gorm:"column:total_borrowed"`
	TotalBorrowedBI         *BigInt `gorm:"column:total_borrowed_bi"`
	CumulativeBorrowed      float64 `gorm:"column:cumulative_borrowed"`
	CumulativeBorrowedBI    *BigInt `gorm:"column:cumulative_borrowed_bi"`
	TotalRepaid             float64 `gorm:"column:total_repaid"`
	TotalRepaidBI           *BigInt `gorm:"column:total_repaid_bi"`
	TotalProfit             float64 `gorm:"column:total_profit"`
	TotalProfitBI           *BigInt `gorm:"column:total_profit_bi"`
	TotalLosses             float64 `gorm:"column:total_losses"`
	TotalLossesBI           *BigInt `gorm:"column:total_losses_bi"`
}

type Parameters struct {
	BlockNum            int64   `gorm:"column:block_num"`
	CreditManager       string  `gorm:"column:credit_manager"`
	MinAmount           *BigInt `gorm:"column:min_amount"`
	MaxAmount           *BigInt `gorm:"column:max_amount"`
	MaxLeverage         *BigInt `gorm:"column:max_leverage"`
	FeeInterest         *BigInt `gorm:"column:fee_interest"`
	FeeLiquidation      *BigInt `gorm:"column:fee_liquidation"`
	LiquidationDiscount *BigInt `gorm:"column:liq_discount"`
}

type CreditManagerUpdate struct {
	*CreditManagerData
	Address string `gorm:"primaryKey"`
}

func (CreditManagerUpdate) TableName() string {
	return "credit_managers"
}

type CreditManagerStat struct {
	*CreditManagerData
	ID       int64  `gorm:"primaryKey"`
	BlockNum int64  `gorm:"column:block_num"`
	Address  string `gorm:"column:credit_manager"`
}

type PnlOnRepay struct {
	Loss           *big.Int
	Profit         *big.Int
	BorrowedAmount *big.Int
}
