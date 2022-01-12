package core

import (
	"encoding/json"
	"github.com/Gearbox-protocol/third-eye/log"
	"math/big"
)

type Debt struct {
	Id                              int64   `gorm:"primaryKey;column:id"`
	BlockNumber                     int64   `gorm:"column:block_num"`
	SessionId                       string  `gorm:"column:session_id"`
	HealthFactor                    *BigInt `gorm:"-"`
	TotalValueBI                    *BigInt `gorm:"-"`
	BorrowedAmountPlusInterestBI    *BigInt `gorm:"-"`
	CalHealthFactor                 *BigInt `gorm:"column:cal_health_factor"`
	CalTotalValueBI                 *BigInt `gorm:"column:cal_total_value"`
	CalBorrowedAmountPlusInterestBI *BigInt `gorm:"column:cal_borrowed_amt_with_interest"`
	CalThresholdValueBI             *BigInt `gorm:"column:cal_threshold_value"`
	AmountToPoolBI                  *BigInt `gorm:"-"`
	ProfitInUSDBI                   *BigInt `gorm:"column:profit_usd_bi"`
	TotalValueInUSDBI               *BigInt `gorm:"column:total_value_usd_bi"`
	CollateralInUSDBI               *BigInt `gorm:"column:collateral_usd_bi"`
}

type CurrentDebt struct {
	SessionId                       string  `gorm:"column:session_id;primaryKey"`
	BlockNumber                     int64   `gorm:"column:block_num"`
	CalHealthFactor                 *BigInt `gorm:"column:cal_health_factor"`
	CalTotalValue                   float64 `gorm:"column:cal_total_value"`
	CalTotalValueBI                 *BigInt `gorm:"column:cal_total_value_bi"`
	CalBorrowedAmountPlusInterest   float64 `gorm:"column:cal_borrowed_amt_with_interest"`
	CalBorrowedAmountPlusInterestBI *BigInt `gorm:"column:cal_borrowed_amt_with_interest_bi"`
	CalThresholdValue               float64 `gorm:"column:cal_threshold_value"`
	CalThresholdValueBI             *BigInt `gorm:"column:cal_threshold_value_bi"`
	AmountToPoolBI                  *BigInt `gorm:"column:amount_to_pool_bi"`
	AmountToPool                    float64 `gorm:"column:amount_to_pool"`
	ProfitInUSDBI                   *BigInt `gorm:"column:profit_usd_bi"`
	CollateralInUSDBI               *BigInt `gorm:"column:collateral_usd_bi"`
}

func (CurrentDebt) TableName() string {
	return "current_debts"
}

type DebtSync struct {
	LastCalculatedAt int64 `gorm:"last_calculated_at"`
}

func (DebtSync) TableName() string {
	return "debt_sync"
}

type TokenDetails struct {
	Price             *big.Int
	Decimals          int8
	TokenLiqThreshold *BigInt `json:"tokenLiqThreshold"`
	Symbol            string  `json:"symbol"`
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
}

type LiquidableAccount struct {
	SessionId            string `gorm:"primaryKey;column:session_id"`
	BlockNum             int64  `gorm:"column:block_num"`
	NotifiedIfLiquidable bool   `gorm:"column:notified_if_liquidable"`
	Updated              bool   `gorm:"-"`
}
