package dc_wrapper

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/ethereum/go-ethereum/common"
)

type DCTesting struct {
	calls map[int64]*test.DCCalls
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
