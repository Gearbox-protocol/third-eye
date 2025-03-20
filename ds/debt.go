package ds

import (
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type DebtEngineI interface {
	Clear()
	ProcessBackLogs()
	CalculateDebtAndClear(to int64, lastSync schemas.LastSync)
	CalCurrentDebts(to int64)
	CalculateDebt()
	GetDebts() core.Json
}

type TokenDetails struct {
	Price             *big.Int
	Decimals          int8
	TokenLiqThreshold *big.Int         `json:"tokenLiqThreshold"`
	Symbol            string           `json:"symbol"`
	Version           core.VersionType `json:"version"`
}
type DebtProfile struct {
	DCData                         *dc.CreditAccountCallData
	*schemas.Debt                  `json:"debt"`
	*schemas.CreditSessionSnapshot `json:"css"`
	RPCBalances                    []dcv2.TokenBalance     `json:"rpcBalances"`
	Tokens                         map[string]TokenDetails `json:"tokens"`
	UnderlyingDecimals             int8                    `json:"underlyingDecimals"`
	*CumIndexAndUToken             `json:"poolDetails"`
	CalcString                     string `json:"calcString"`
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

func CallTraceAllowed(client core.ClientI) bool {
	return utils.Contains([]int64{1, 42161, 10, 7878, 1337}, core.GetChainId(client))
}
