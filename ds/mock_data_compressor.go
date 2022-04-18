package ds

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type DCTesting struct {
	calls map[int64]*DCCalls
}

type DCCalls struct {
	Pools    map[string]TestPoolCallData
	CMs      map[string]TestCMCallData
	Accounts map[string]TestAccountCallData
}

func NewDCCalls() *DCCalls {
	return &DCCalls{
		Pools:    make(map[string]TestPoolCallData),
		CMs:      make(map[string]TestCMCallData),
		Accounts: make(map[string]TestAccountCallData),
	}
}

type TestTokenBalance struct {
	Token     string       `json:"token"`
	Balance   *core.BigInt `json:"balance"`
	IsAllowed bool         `json:"isAllowed"`
}
type TestAccountCallData struct {
	Addr                       string             `json:"address"`
	Borrower                   string             `json:"borrower"`
	InUse                      bool               `json:"inUse"`
	CreditManager              string             `json:"creditManager"`
	UnderlyingToken            string             `json:"underlyingToken"`
	BorrowedAmountPlusInterest *core.BigInt       `json:"borrowAmountPlusInterest"`
	TotalValue                 *core.BigInt       `json:"totalValue"`
	HealthFactor               *core.BigInt       `json:"healthFactor"`
	BorrowRate                 *core.BigInt       `json:"borrowRate"`
	Balances                   []TestTokenBalance `json:"balances"`
	RepayAmount                *core.BigInt       `json:"repayAmount"`
	LiquidationAmount          *core.BigInt       `json:"liquidationAmount"`
	CanBeClosed                bool               `json:"canBeClosed"`
	BorrowedAmount             *core.BigInt       `json:"borrowedAmount"`
	CumulativeIndexAtOpen      *core.BigInt       `json:"cumulativeIndexAtOpen"`
	Since                      *core.BigInt       `json:"since"`
}

type TestPoolCallData struct {
	Addr                   string       `json:"address"`
	IsWETH                 bool         `json:"isWETH"`
	UnderlyingToken        string       `json:"underlyingToken"`
	DieselToken            string       `json:"dieselToken"`
	LinearCumulativeIndex  *core.BigInt `json:"linearCumulativeIndex"`
	AvailableLiquidity     *core.BigInt `json:"availableLiquidity"`
	ExpectedLiquidity      *core.BigInt `json:"expectedLiquidity"`
	ExpectedLiquidityLimit *core.BigInt `json:"expectedLiquidityLimit"`
	TotalBorrowed          *core.BigInt `json:"totalBorrowed"`
	DepositAPYRAY          *core.BigInt `json:"depositAPY"`
	BorrowAPYRAY           *core.BigInt `json:"borrowAPY"`
	DieselRateRAY          *core.BigInt `json:"dieselRate"`
	WithdrawFee            *core.BigInt `json:"withdrawFee"`
	CumulativeIndexRAY     *core.BigInt `json:"cumulativeIndex"`
}

type TestCMCallData struct {
	Addr               string       `json:"address"`
	HasAccount         bool         `json:"hasAddress"`
	UnderlyingToken    string       `json:"underlyingToken"`
	IsWETH             bool         `json:"isWETH"`
	CanBorrow          bool         `json:"canBorrow"`
	BorrowRate         *core.BigInt `json:"borrowRate"`
	MinAmount          *core.BigInt `json:"minAmount"`
	MaxAmount          *core.BigInt `json:"maxAmount"`
	MaxLeverageFactor  *core.BigInt `json:"maxLeverageFactor"`
	AvailableLiquidity *core.BigInt `json:"availableLiquidity"`
}

func (t *DCTesting) getPoolData(blockNum int64, key string) (mainnet.DataTypesPoolData, error) {
	obj := t.calls[blockNum].Pools[key]
	return mainnet.DataTypesPoolData{
		Addr:                   common.HexToAddress(obj.Addr),
		IsWETH:                 obj.IsWETH,
		UnderlyingToken:        common.HexToAddress(obj.UnderlyingToken),
		DieselToken:            common.HexToAddress(obj.DieselToken),
		LinearCumulativeIndex:  (*big.Int)(obj.LinearCumulativeIndex),
		AvailableLiquidity:     (*big.Int)(obj.AvailableLiquidity),
		ExpectedLiquidity:      (*big.Int)(obj.ExpectedLiquidity),
		ExpectedLiquidityLimit: (*big.Int)(obj.ExpectedLiquidityLimit),
		TotalBorrowed:          (*big.Int)(obj.TotalBorrowed),
		DepositAPYRAY:          (*big.Int)(obj.DepositAPYRAY),
		BorrowAPYRAY:           (*big.Int)(obj.BorrowAPYRAY),
		DieselRateRAY:          (*big.Int)(obj.DieselRateRAY),
		WithdrawFee:            (*big.Int)(obj.WithdrawFee),
		CumulativeIndexRAY:     (*big.Int)(obj.CumulativeIndexRAY),
	}, nil
}
func (t *DCTesting) getCMData(blockNum int64, key string) (mainnet.DataTypesCreditManagerData, error) {
	obj := t.calls[blockNum].CMs[key]
	return mainnet.DataTypesCreditManagerData{
		Addr:               common.HexToAddress(obj.Addr),
		HasAccount:         obj.HasAccount,
		UnderlyingToken:    common.HexToAddress(obj.UnderlyingToken),
		IsWETH:             obj.IsWETH,
		CanBorrow:          obj.CanBorrow,
		BorrowRate:         (*big.Int)(obj.BorrowRate),
		MinAmount:          (*big.Int)(obj.MinAmount),
		MaxAmount:          (*big.Int)(obj.MaxAmount),
		MaxLeverageFactor:  (*big.Int)(obj.MaxLeverageFactor),
		AvailableLiquidity: (*big.Int)(obj.AvailableLiquidity),
	}, nil
}
func (t *DCTesting) getAccountData(blockNum int64, key string) (mainnet.DataTypesCreditAccountDataExtended, error) {
	obj := t.calls[blockNum].Accounts[key]
	var balances []mainnet.DataTypesTokenBalance
	for _, balance := range obj.Balances {
		balances = append(balances, mainnet.DataTypesTokenBalance{
			Token:     common.HexToAddress(balance.Token),
			Balance:   (*big.Int)(balance.Balance),
			IsAllowed: balance.IsAllowed,
		})
	}
	return mainnet.DataTypesCreditAccountDataExtended{
		Addr:                       common.HexToAddress(obj.Addr),
		Borrower:                   common.HexToAddress(obj.Borrower),
		InUse:                      obj.InUse,
		CreditManager:              common.HexToAddress(obj.CreditManager),
		UnderlyingToken:            common.HexToAddress(obj.UnderlyingToken),
		BorrowedAmountPlusInterest: (*big.Int)(obj.BorrowedAmountPlusInterest),
		TotalValue:                 (*big.Int)(obj.TotalValue),
		HealthFactor:               (*big.Int)(obj.HealthFactor),
		BorrowRate:                 (*big.Int)(obj.BorrowRate),
		Balances:                   balances,
		RepayAmount:                (*big.Int)(obj.RepayAmount),
		LiquidationAmount:          (*big.Int)(obj.LiquidationAmount),
		CanBeClosed:                obj.CanBeClosed,
		BorrowedAmount:             (*big.Int)(obj.BorrowedAmount),
		CumulativeIndexAtOpen:      (*big.Int)(obj.CumulativeIndexAtOpen),
		Since:                      (*big.Int)(obj.Since),
	}, nil
}
