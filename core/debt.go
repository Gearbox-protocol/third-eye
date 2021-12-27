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
	HealthFactor                    *BigInt `gorm:"column:health_factor"`
	TotalValueBI                    *BigInt `gorm:"column:total_value"`
	BorrowedAmountPlusInterestBI    *BigInt `gorm:"column:borrowed_amt_with_interest"`
	CalHealthFactor                 *BigInt `gorm:"column:cal_health_factor"`
	CalTotalValueBI                 *BigInt `gorm:"column:cal_total_value"`
	CalBorrowedAmountPlusInterestBI *BigInt `gorm:"column:cal_borrowed_amt_with_interest"`
	CalThresholdValueBI             *BigInt `gorm:"column:cal_threshold_value"`
	ProfitBI                        *BigInt `gorm:"-"`
	LossBI                          *BigInt `gorm:"-"`
	RepayAmountBI                   *BigInt `gorm:"-"`
	LiqAmountBI                     *BigInt `gorm:"-"`
}

type CurrentDebt struct {
	SessionId                       string  `gorm:"column:session_id;primaryKey"`
	BlockNumber                     int64   `gorm:"column:block_num"`
	CalHealthFactor                 *BigInt `gorm:"column:cal_health_factor"`
	CalTotalValue                   float64 `gorm:"column:cal_total_value"`
	CalTotalValueBI                 *BigInt `gorm:"column:cal_total_value_bi"`
	Profit                          float64 `gorm:"column:profit"`
	Loss                            float64 `gorm:"column:loss"`
	RepayAmount                     float64 `gorm:"column:repay_amount"`
	LiqAmount                       float64 `gorm:"column:liq_amount"`
	CalBorrowedAmountPlusInterest   float64 `gorm:"column:cal_borrowed_amt_with_interest"`
	CalBorrowedAmountPlusInterestBI *BigInt `gorm:"column:cal_borrowed_amt_with_interest_bi"`
	CalThresholdValue               float64 `gorm:"column:cal_threshold_value"`
	CalThresholdValueBI             *BigInt `gorm:"column:cal_threshold_value_bi"`
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
	SessionId string `gorm:"primaryKey;column:session_id"`
	BlockNum  int64  `gorm:"block_num"`
	Updated   bool   `gorm:"-"`
}
