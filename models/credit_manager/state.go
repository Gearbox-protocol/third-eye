package credit_manager

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CreditManager) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStatePresent = true
	switch obj.(type) {
	case (*schemas.CreditManagerState):
		state := obj.(*schemas.CreditManagerState)
		mdl.State = state
	case (map[string]string):
		sessions := obj.(map[string]string)
		mdl.State.Sessions = sessions
	case *schemas.PnlOnRepay:
		pnl := obj.(*schemas.PnlOnRepay)
		mdl.pnlOnCM.Set(pnl)
	default:
		log.Fatal("Type assertion for credit manager state failed")
	}
}

func (mdl *CreditManager) GetUnderlyingState() interface{} {
	return mdl.State
}

func (mdl *CreditManager) AddCreditOwnerSession(owner, sessionId string) {
	mdl.State.Sessions[owner] = sessionId
}

func (mdl *CreditManager) RemoveCreditOwnerSession(owner string) {
	delete(mdl.State.Sessions, owner)
}

func (mdl *CreditManager) GetCreditOwnerSession(owner string, dontFail ...bool) string {
	sessionId := mdl.State.Sessions[owner]
	if len(dontFail) == 0 && sessionId == "" {
		panic(
			fmt.Sprintf("session id not found for %s in %+v\n", owner, mdl.State.Sessions),
		)
	}
	return sessionId
}

func (mdl *CreditManager) GetUnderlyingToken() string {
	return mdl.State.UnderlyingToken
}

func (mdl *CreditManager) calculateCMStat(blockNum int64) {

	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	state, err := mdl.Repo.GetDCWrapper().GetCreditManagerData(
		opts,
		common.HexToAddress(mdl.GetAddress()),
		common.HexToAddress(mdl.GetAddress()),
	)
	if err != nil {
		log.Fatal("[CreditManagerModel] Cant get data from data compressor", err)
	}
	mdl.State.IsWETH = state.IsWETH
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
	// mdl.State.MinAmount = (*core.BigInt)(state.MinAmount)
	// mdl.State.MaxAmount = (*core.BigInt)(state.MaxAmount)

	mdl.State.BorrowRateBI = (*core.BigInt)(state.BorrowRate)
	mdl.State.BorrowRate = utils.GetFloat64Decimal(state.BorrowRate, 25)

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
			BorrowRateBI:            core.NewBigInt(mdl.State.BorrowRateBI),
			BorrowRate:              mdl.State.BorrowRate,
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
