package core

import (
	"encoding/json"
	"github.com/Gearbox-protocol/third-eye/log"
	"math/big"
)

type Debt struct {
	Id                              int64   `gorm:"primaryKey;column:id" json:"-"`
	BlockNumber                     int64   `gorm:"column:block_num" json:"blockNum"`
	SessionId                       string  `gorm:"column:session_id" json:"sessionId"`
	HealthFactor                    *BigInt `gorm:"-" json:"-"`
	TotalValueBI                    *BigInt `gorm:"-" json:"-"`
	BorrowedAmountPlusInterestBI    *BigInt `gorm:"-" json:"-"`
	CalHealthFactor                 *BigInt `gorm:"column:cal_health_factor" json:"calHealthFactor"`
	CalTotalValueBI                 *BigInt `gorm:"column:cal_total_value_bi" json:"calTotalValue"`
	CalBorrowedAmountPlusInterestBI *BigInt `gorm:"column:cal_borrowed_amt_with_interest_bi" json:"calBorrowedAmountWithInterest"`
	CalThresholdValueBI             *BigInt `gorm:"column:cal_threshold_value_bi" json:"calThresholdValue"`
	AmountToPoolBI                  *BigInt `gorm:"-" json:"-"`
	ProfitInUSD                     float64 `gorm:"column:profit_usd" json:"profitUSD"`
	CollateralInUSD                 float64 `gorm:"column:collateral_usd" json:"collateralUSD"`
	CollateralInUnderlying          float64 `gorm:"column:collateral_underlying" json:"collateralUnderlying"`
	ProfitInUnderlying              float64 `gorm:"column:profit_underlying" json:"profitUnderlying"`
	// field not present in current_debts
	TotalValueInUSD                 float64 `gorm:"column:total_value_usd" json:"totalValueInUSD"`
}

type CurrentDebt struct {
	SessionId                       string  `gorm:"column:session_id;primaryKey" json:"sessionId"`
	BlockNumber                     int64   `gorm:"column:block_num" json:"blockNum"`
	CalHealthFactor                 *BigInt `gorm:"column:cal_health_factor" json:"calHealthFactor"`
	CalTotalValue                   float64 `gorm:"column:cal_total_value" json:"calTotalValue"`
	CalTotalValueBI                 *BigInt `gorm:"column:cal_total_value_bi" json:"-"`
	CalBorrowedAmountPlusInterest   float64 `gorm:"column:cal_borrowed_amt_with_interest" json:"calBorrowedAmountPlusInterest"`
	CalBorrowedAmountPlusInterestBI *BigInt `gorm:"column:cal_borrowed_amt_with_interest_bi" json:"-"`
	CalThresholdValue               float64 `gorm:"column:cal_threshold_value" json:"calThresholdValue"`
	CalThresholdValueBI             *BigInt `gorm:"column:cal_threshold_value_bi" json:"-"`
	AmountToPoolBI                  *BigInt `gorm:"column:amount_to_pool_bi" json:"-"`
	AmountToPool                    float64 `gorm:"column:amount_to_pool" json:"amountToPool"`
	ProfitInUSD                     float64 `gorm:"column:profit_usd" json:"profitUSD"`
	ProfitInUnderlying              float64 `gorm:"column:profit_underlying" json:"profitUnderlying"`
	CollateralInUSD                 float64 `gorm:"column:collateral_usd" json:"collateralUSD"`
	CollateralInUnderlying          float64 `gorm:"column:collateral_underlying" json:"collateralUnderlying"`
}

func (CurrentDebt) TableName() string {
	return "current_debts"
}

type DebtSync struct {
	LastCalculatedAt int64 `gorm:"column:last_calculated_at"`
	FieldSet         bool  `gorm:"column:field_set;primaryKey"`
}

func (DebtSync) TableName() string {
	return "debt_sync"
}

type TokenDetails struct {
	Price             *big.Int
	Decimals          int8
	TokenLiqThreshold *BigInt `json:"tokenLiqThreshold"`
	Symbol            string  `json:"symbol"`
	Version           int16   `json:"version"`
}
type DebtProfile struct {
	*Debt                  `json:"debt"`
	*CreditSessionSnapshot `json:"css"`
	RPCBalances            *JsonBalance            `json:"rpcBalances"`
	Tokens                 map[string]TokenDetails `json:"tokens"`
	UnderlyingDecimals     int8                    `json:"underlyingDecimals"`
	*CumIndexAndUToken     `json:"poolDetails"`
}

type CumIndexAndUToken struct {
	CumulativeIndex *big.Int
	Token           string
	Decimals        int8
	Symbol          string
	PriceInETH      *big.Int
	PriceInUSD      *big.Int
}

func (c *CumIndexAndUToken) GetPrice(version int16) *big.Int {
	switch version {
	case 1:
		return c.PriceInETH
	case 2:
		return c.PriceInUSD
	}
	return nil
}

func (debt *DebtProfile) Json() []byte {
	str, err := json.Marshal(debt)
	if err != nil {
		log.Fatal(err)
	}
	return str
}

type ProfileTable struct {
	Profile string `gorm:"column:profile"`
}

func (ProfileTable) TableName() string {
	return "profiles"
}

type DebtEngineI interface {
	Clear()
	ProcessBackLogs()
	CalculateDebtAndClear(to int64)
	CalCurrentDebts(to int64)
	CalculateDebt()
	GetDebts() Json
}

type LiquidableAccount struct {
	SessionId            string `gorm:"primaryKey;column:session_id"`
	BlockNum             int64  `gorm:"column:block_num"`
	NotifiedIfLiquidable bool   `gorm:"column:notified_if_liquidable"`
	Updated              bool   `gorm:"-"`
}
