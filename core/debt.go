package core

import (
	"encoding/json"
	"github.com/Gearbox-protocol/third-eye/log"
	"math/big"
)

type Debt struct {
	Id                              int64  `gorm:"primaryKey;column:id;autoincrement:true"`
	BlockNumber                     int64  `gorm:"column:block_num"`
	SessionId                       string `gorm:"column:session_id"`
	HealthFactor                    int64  `gorm:"column:health_factor"`
	TotalValueBI                    string `gorm:"column:total_value"`
	BorrowedAmountPlusInterestBI    string `gorm:"column:borrowed_amt_with_interest"`
	CalHealthFactor                 int64  `gorm:"column:cal_health_factor"`
	CalTotalValue                   string `gorm:"column:cal_total_value"`
	CalBorrowedAmountPlusInterestBI string `gorm:"column:cal_borrowed_amt_with_interest"`
	CalThresholdValueBI             string `gorm:"column:cal_threshold_value"`
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
	RPCBalances            JsonBalance             `json:"rpcBalances"`
	Tokens                 map[string]TokenDetails `json:"tokens"`
	UnderlyingDecimals     int8                    `json:"underlyingDecimals"`
	*CumIndexAndUToken     `json:"poolDetails"`
}

type CumIndexAndUToken struct {
	CumulativeIndex *big.Int
	Token           string
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
