package dc_wrapper

import (
	"fmt"
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/ethereum/go-ethereum/common"
)

type DCTesting struct {
	calls  map[int64]*test.DCCalls
	client core.ClientI
}

func (t DCTesting) getPoolData(blockNum int64, addr string) (dcv2.PoolData, error) {
	obj := t.calls[blockNum].Pools[addr]
	return dcv2.PoolData{
		Addr:                   common.HexToAddress(obj.Addr),
		IsWETH:                 obj.IsWETH,
		Underlying:             common.HexToAddress(obj.UnderlyingToken),
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
func (t *DCTesting) getCMData(blockNum int64, key string) (dcv2.CreditManagerData, error) {
	if t.calls == nil || t.calls[blockNum] == nil {
		return dcv2.CreditManagerData{}, nil
	}
	obj := t.calls[blockNum].CMs[key]
	return dcv2.CreditManagerData{
		Addr:               common.HexToAddress(obj.Addr),
		Underlying:         common.HexToAddress(obj.UnderlyingToken),
		IsWETH:             obj.IsWETH,
		CanBorrow:          obj.CanBorrow,
		BorrowRate:         (*big.Int)(obj.BorrowRate),
		MinAmount:          (*big.Int)(obj.MinAmount),
		MaxAmount:          (*big.Int)(obj.MaxAmount),
		MaxLeverageFactor:  (*big.Int)(obj.MaxLeverageFactor),
		AvailableLiquidity: (*big.Int)(obj.AvailableLiquidity),
	}, nil
}
func (t *DCTesting) getAccountData(blockNum int64, key string) (dcv2.CreditAccountData, error) {
	obj := t.calls[blockNum].Accounts[key]
	//
	var maskInBits string
	if obj.Version != 2 {
		mask := getMask(t.client, blockNum, core.NULL_ADDR, common.HexToAddress(obj.Addr))
		maskInBits = fmt.Sprintf("%b", mask)
	}
	maskLen := len(maskInBits)
	//
	var balances []dcv2.TokenBalance
	for ind, entry := range obj.Balances {
		balance := dcv2.TokenBalance{
			Token:     common.HexToAddress(entry.Token),
			Balance:   (*big.Int)(entry.Balance),
			IsAllowed: entry.IsAllowed,
		}
		if obj.Version == 2 {
			balance.IsEnabled = entry.IsEnabled
		} else {
			if maskLen-ind-1 >= 0 {
				balance.IsEnabled = maskInBits[maskLen-ind-1] == '1'
			}
		}
		balances = append(balances, balance)
	}
	return dcv2.CreditAccountData{
		Addr:                       common.HexToAddress(obj.Addr),
		Borrower:                   common.HexToAddress(obj.Borrower),
		InUse:                      obj.InUse,
		CreditManager:              common.HexToAddress(obj.CreditManager),
		Underlying:                 common.HexToAddress(obj.UnderlyingToken),
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
