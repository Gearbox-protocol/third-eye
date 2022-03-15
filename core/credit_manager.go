package core

import (
	"math/big"
)

func (CreditManagerState) TableName() string {
	return "credit_managers"
}

type CreditManagerState struct {
	CreditManagerData
	Address           string            `gorm:"primaryKey" json:"address"`
	IsWETH            bool              `gorm:"is_weth"`
	PoolAddress       string            `gorm:"column:pool_address" json:"pool"`
	UnderlyingToken   string            `gorm:"column:underlying_token" json:"underlyingToken"`
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
	BlockNum            int64   `gorm:"column:block_num;primaryKey"`
	CreditManager       string  `gorm:"column:credit_manager;primaryKey"`
	MinAmount           *BigInt `gorm:"column:min_amount"`
	MaxAmount           *BigInt `gorm:"column:max_amount"`
	MaxLeverage         *BigInt `gorm:"column:max_leverage"`
	FeeInterest         *BigInt `gorm:"column:fee_interest"`
	FeeLiquidation      *BigInt `gorm:"column:fee_liquidation"`
	LiquidationDiscount *BigInt `gorm:"column:liq_discount"`
}

func NewParameters() *Parameters {
	return &Parameters{
		MinAmount:           (*BigInt)(big.NewInt(0)),
		MaxAmount:           (*BigInt)(big.NewInt(0)),
		MaxLeverage:         (*BigInt)(big.NewInt(0)),
		FeeInterest:         (*BigInt)(big.NewInt(0)),
		FeeLiquidation:      (*BigInt)(big.NewInt(0)),
		LiquidationDiscount: (*BigInt)(big.NewInt(0)),
	}
}

func (old *Parameters) Diff(new *Parameters) *Json {
	obj := Json{}
	obj["minAmount"] = []*BigInt{old.MinAmount, new.MinAmount}
	obj["maxAmount"] = []*BigInt{old.MaxAmount, new.MaxAmount}
	obj["maxLeverage"] = []*BigInt{old.MaxLeverage, new.MaxLeverage}
	obj["feeInterest"] = []*BigInt{old.FeeInterest, new.FeeInterest}
	obj["feeLiquidation"] = []*BigInt{old.FeeLiquidation, new.FeeLiquidation}
	obj["LiquidationDiscount"] = []*BigInt{old.LiquidationDiscount, new.LiquidationDiscount}
	return &obj
}

type FastCheckParams struct {
	BlockNum        int64   `gorm:"column:block_num;primaryKey"`
	CreditManager   string  `gorm:"column:credit_manager;primaryKey"`
	ChiThreshold    *BigInt `gorm:"column:chi_threshold"`
	HFCheckInterval *BigInt `gorm:"column:hf_checkinterval"`
}

func NewFastCheckParams() *FastCheckParams {
	return &FastCheckParams{
		ChiThreshold:    (*BigInt)(big.NewInt(0)),
		HFCheckInterval: (*BigInt)(big.NewInt(0)),
	}
}
func (old *FastCheckParams) Diff(new *FastCheckParams) *Json {
	obj := Json{}
	obj["chiThreshold"] = []*BigInt{old.ChiThreshold, new.ChiThreshold}
	obj["fastDelay"] = []*BigInt{old.HFCheckInterval, new.HFCheckInterval}
	return &obj
}

func (FastCheckParams) TableName() string {
	return "fast_check_params"
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

type TransferAccountAllowed struct {
	From        string `gorm:"column:sender"`
	To          string `gorm:"column:receiver"`
	Allowed     bool   `gorm:"column:allowed"`
	BlockNumber int64  `gorm:"column:block_num;primaryKey"`
	LogId       int64  `gorm:"column:log_id;primaryKey"`
}
