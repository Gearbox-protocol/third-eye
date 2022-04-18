package ds

import (
	"encoding/json"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"math/big"
)

type DebtEngineI interface {
	Clear()
	ProcessBackLogs()
	CalculateDebtAndClear(to int64)
	CalCurrentDebts(to int64)
	CalculateDebt()
	GetDebts() core.Json
}

type TokenDetails struct {
	Price             *big.Int
	Decimals          int8
	TokenLiqThreshold *core.BigInt `json:"tokenLiqThreshold"`
	Symbol            string       `json:"symbol"`
	Version           int16        `json:"version"`
}
type DebtProfile struct {
	*schemas.Debt                  `json:"debt"`
	*schemas.CreditSessionSnapshot `json:"css"`
	RPCBalances                    *core.JsonBalance       `json:"rpcBalances"`
	Tokens                         map[string]TokenDetails `json:"tokens"`
	UnderlyingDecimals             int8                    `json:"underlyingDecimals"`
	*CumIndexAndUToken             `json:"poolDetails"`
	DCFields                       struct {
		HealthFactor                 *core.BigInt
		TotalValueBI                 *core.BigInt
		BorrowedAmountPlusInterestBI *core.BigInt
	}
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

func (profile *DebtProfile) Json() []byte {
	profile.DCFields.BorrowedAmountPlusInterestBI = profile.Debt.BorrowedAmountPlusInterestBI
	profile.DCFields.TotalValueBI = profile.Debt.TotalValueBI
	profile.DCFields.HealthFactor = profile.Debt.HealthFactor
	str, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}
	return str
}
