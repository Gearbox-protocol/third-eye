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
	Decimals         int8
	Symbol           string
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
	CalculateDebtAndClear()
}

type LiquidableAccount struct {
	SessionId string `gorm:"primaryKey;column:session_id"`
	BlockNum  int64  `gorm:"block_num"`
	Updated   bool   `gorm:"-"`
}
