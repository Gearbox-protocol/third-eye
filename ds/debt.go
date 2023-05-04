package ds

import (
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
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
	TokenLiqThreshold *core.BigInt     `json:"tokenLiqThreshold"`
	Symbol            string           `json:"symbol"`
	Version           core.VersionType `json:"version"`
}
type DebtProfile struct {
	DCData                         *dcv2.CreditAccountData
	*schemas.Debt                  `json:"debt"`
	*schemas.CreditSessionSnapshot `json:"css"`
	RPCBalances                    []dcv2.TokenBalance     `json:"rpcBalances"`
	Tokens                         map[string]TokenDetails `json:"tokens"`
	UnderlyingDecimals             int8                    `json:"underlyingDecimals"`
	*CumIndexAndUToken             `json:"poolDetails"`
}

type CumIndexAndUToken struct {
	CumulativeIndex *big.Int
	Token           string
	Decimals        int8
	Symbol          string
	PoolAddr        string
}

func (profile *DebtProfile) String() string {
	return utils.ToJson(profile)
}
