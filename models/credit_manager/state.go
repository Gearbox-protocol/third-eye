package credit_manager

import (
	"fmt"

	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (mdl *CreditManager) SetState(obj interface{}) {
	state, ok := obj.(*core.CreditManager)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	mdl.State = state
}

func (mdl *CreditManager) AddCreditOwnerSession(owner, sessionId string) {
	mdl.State.Sessions.Set(owner, sessionId)
}

func (mdl *CreditManager) RemoveCreditOwnerSession(owner string) {
	mdl.State.Sessions.Remove(owner)
}

func (mdl *CreditManager) GetCreditOwnerSession(owner string) string {
	sessionId := mdl.State.Sessions.Get(owner)
	if sessionId == "" {
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
	state, err := mdl.Repo.GetDataCompressor(blockNum).GetCreditManagerData(
		opts,
		common.HexToAddress(mdl.GetAddress()),
		common.HexToAddress(mdl.GetAddress()),
	)
	if err != nil {
		log.Fatal("[CreditManagerModel] Cant get data from data compressor", err)
	}

	mdl.State.MinAmount = (*core.BigInt)(state.MinAmount)
	mdl.State.MaxAmount = (*core.BigInt)(state.MaxAmount)

	mdl.State.BorrowRateBI = (*core.BigInt)(state.BorrowRate)
	mdl.State.BorrowRate = utils.GetFloat64Decimal(state.BorrowRate, 25)

	mdl.State.AvailableLiquidityBI = (*core.BigInt)(state.AvailableLiquidity)
	mdl.State.AvailableLiquidity = utils.GetFloat64Decimal(state.AvailableLiquidity, mdl.GetUnderlyingDecimal())

	stats := &core.CreditManagerStat{
		Address:  mdl.Address,
		BlockNum: blockNum,
		CreditManagerData: &core.CreditManagerData{
			// fetched from data compressor
			MinAmount:            core.NewBigInt(mdl.State.MinAmount),
			MaxAmount:            core.NewBigInt(mdl.State.MaxAmount),
			BorrowRateBI:         core.NewBigInt(mdl.State.BorrowRateBI),
			BorrowRate:           mdl.State.BorrowRate,
			AvailableLiquidityBI: core.NewBigInt(mdl.State.AvailableLiquidityBI),
			AvailableLiquidity:   mdl.State.AvailableLiquidity,
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
