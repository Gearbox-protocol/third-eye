package cm_common

import (
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (mdl *CMCommon) CalculateCMStat(blockNum int64, state dcv2.CreditManagerData) {
	//
	mdl.State.IsWETH = state.IsWETH
	//
	bororwAmountForBlock := mdl.getBorrowAmountForBlockAndClear()
	mdl.State.TotalBorrowedBI = core.AddCoreAndInt(mdl.State.TotalBorrowedBI, bororwAmountForBlock)
	mdl.State.TotalBorrowed = utils.GetFloat64Decimal(mdl.State.TotalBorrowedBI.Convert(), mdl.GetUnderlyingDecimal())
	//
	// pnl on repay
	pnl := mdl.pnlOnCM.Get(blockNum)
	if pnl != nil {
		mdl.State.TotalBorrowedBI = core.SubCoreAndInt(mdl.State.TotalBorrowedBI, pnl.BorrowedAmount)
		mdl.State.TotalBorrowed = utils.GetFloat64Decimal(mdl.State.TotalBorrowedBI.Convert(), mdl.GetUnderlyingDecimal())
		mdl.State.TotalLossesBI = core.AddCoreAndInt(mdl.State.TotalLossesBI, pnl.Loss)
		mdl.State.TotalLosses = utils.GetFloat64Decimal(mdl.State.TotalLossesBI.Convert(), mdl.GetUnderlyingDecimal())
		mdl.State.TotalProfitBI = core.AddCoreAndInt(mdl.State.TotalProfitBI, pnl.Profit)
		mdl.State.TotalProfit = utils.GetFloat64Decimal(mdl.State.TotalProfitBI.Convert(), mdl.GetUnderlyingDecimal())
	}
	mdl.State.MinAmount = (*core.BigInt)(state.MinAmount)
	mdl.State.MaxAmount = (*core.BigInt)(state.MaxAmount)

	mdl.State.AvailableLiquidityBI = (*core.BigInt)(state.AvailableLiquidity)
	mdl.State.AvailableLiquidity = utils.GetFloat64Decimal(state.AvailableLiquidity, mdl.GetUnderlyingDecimal())

	stats := &schemas.CreditManagerStat{
		Address:  mdl.Address,
		BlockNum: blockNum,
		CreditManagerData: &schemas.CreditManagerData{
			// fetched from data compressor
			OpenedAccountsCount:     mdl.State.OpenedAccountsCount,
			TotalOpenedAccounts:     mdl.State.TotalOpenedAccounts,
			TotalClosedAccounts:     mdl.State.TotalClosedAccounts,
			TotalRepaidAccounts:     mdl.State.TotalRepaidAccounts,
			TotalLiquidatedAccounts: mdl.State.TotalLiquidatedAccounts,
			AvailableLiquidityBI:    core.NewBigInt(mdl.State.AvailableLiquidityBI),
			AvailableLiquidity:      mdl.State.AvailableLiquidity,
			// calculated in this application
			TotalBorrowed:   mdl.State.TotalBorrowed,
			TotalBorrowedBI: core.NewBigInt(mdl.State.TotalBorrowedBI),
			TotalRepaid:     mdl.State.TotalRepaid,
			TotalRepaidBI:   core.NewBigInt(mdl.State.TotalRepaidBI),
			TotalProfit:     mdl.State.TotalProfit,
			TotalProfitBI:   core.NewBigInt(mdl.State.TotalProfitBI),
			TotalLosses:     mdl.State.TotalLosses,
			TotalLossesBI:   core.NewBigInt(mdl.State.TotalLossesBI),
		},
	}
	mdl.Repo.AddCreditManagerStats(stats)
}

func (mdl *CMCommon) CMStatsOnOpenAccount(borrowAmount *big.Int) {
	// manager state
	mdl.State.TotalOpenedAccounts++
	mdl.State.OpenedAccountsCount++
	mdl.AddBorrowAmountForBlock(borrowAmount)
}

// borroweAmount can't be negative
func (mdl *CMCommon) AddBorrowAmountForBlock(borrowAmount *big.Int) {
	if borrowAmount.Sign() < 0 {
		log.Fatal("Borrowed Amount can't be negative. As repaid amount is tracked on pool")
	}
	if mdl.borrowedAmountForBlock == nil {
		mdl.borrowedAmountForBlock = new(big.Int)
	}
	mdl.borrowedAmountForBlock = new(big.Int).Add(mdl.borrowedAmountForBlock, borrowAmount)
}

func (mdl *CMCommon) getBorrowAmountForBlockAndClear() *big.Int {
	if mdl.borrowedAmountForBlock == nil {
		return new(big.Int)
	}
	lastValue := mdl.borrowedAmountForBlock
	mdl.borrowedAmountForBlock = new(big.Int)
	return lastValue
}
