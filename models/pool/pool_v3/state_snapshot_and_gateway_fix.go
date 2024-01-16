package pool_v3

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/ethereum/go-ethereum/common"
)

// asset is already in underlying
func (mdl *Poolv3) fixPoolLedgerAddrForGateway() {
	// for remove liquidity
	for _, removeLiqEvent := range mdl.gatewayHandler.GetRemoveLiqEventsAndClear() {
		// add poolLedger to repository
		mdl.Repo.AddPoolLedger(removeLiqEvent)
	}
}

func (mdl *Poolv3) OnBlockChange(inputBlock int64) (call multicall.Multicall2Call, processFn func(multicall.Multicall2Result)) {
	// if no addLiquidity/removeLiquidity events are emitted then lastEventBlock is zero.Thus,  fixPoolLedgerAddrForGateway will not be called and pool snapshot is not created
	if mdl.lastEventBlock == 0 ||
		// datacompressor works for pool address only after the address is registered with contractregister
		// i.e. discoveredAt
		mdl.lastEventBlock < mdl.DiscoveredAt {
		return multicall.Multicall2Call{}, nil
	}
	if inputBlock != mdl.lastEventBlock {
		log.Fatal("[PoolServiceModel]: OnBlockChange called with wrong block number")
	}
	// set to zero, we only create poolstat snapshot when there is a event with changed pool cumulative interest rate
	mdl.lastEventBlock = 0
	return mdl.getCallAndProcessFn(inputBlock)
}

func (mdl *Poolv3) getCallAndProcessFn(inputB int64) (multicall.Multicall2Call, func(multicall.Multicall2Result)) {
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetPoolData(mdl.GetVersion(), inputB, common.HexToAddress(mdl.Address))
	if err != nil {
		if err == dc_wrapper.NO_DC_FOUND_ERR {
			return multicall.Multicall2Call{}, nil
		}
		log.Fatal("[PoolService] Cant create call for data compressor", err)
	}
	return call, func(result multicall.Multicall2Result) {
		if !result.Success {
			log.Fatal("[PoolService] Cant get data from RPC", inputB, mdl.Address)
		}
		poolState, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatal("[PoolService] Cant process result for data compressor", err, inputB, mdl.Address, mdl.GetVersion())
		}
		//
		mdl.createSnapshot(inputB, poolState)
		// it uses dieselRate from create snapshot so this call should be after createSnapshots
		mdl.fixPoolLedgerAddrForGateway()
	}
}

func (mdl *Poolv3) onBlockChangeInternally(inputB int64) {
	call, processFn := mdl.getCallAndProcessFn(inputB)
	if processFn != nil {
		result := core.MakeMultiCall(mdl.Client, inputB, false, []multicall.Multicall2Call{call})
		processFn(result[0])
	}
}

func (p *Poolv3) createSnapshot(blockNum int64, state dc.PoolCallData) {
	token := p.Repo.GetToken(p.State.UnderlyingToken)
	p.State.IsWETH = dc.IsWETH(p.Client, state.Underlying)
	p.State.BaseBorrowAPYBI = (*core.BigInt)(state.BaseInterestRate)
	p.State.DepositAPYBI = (*core.BigInt)(state.SupplyRate)
	p.Repo.AddPoolStat(&schemas.PoolStat{
		BlockNum:        blockNum,
		Address:         p.Address,
		UniqueUsers:     p.Repo.GetPoolUniqueUserLen(p.Address),
		TotalBorrowedBI: (*core.BigInt)(state.TotalBorrowed),
		TotalBorrowed:   utils.GetFloat64Decimal(state.TotalBorrowed, token.Decimals),

		ExpectedLiquidityBI: (*core.BigInt)(state.ExpectedLiquidity),
		ExpectedLiquidity:   utils.GetFloat64Decimal(state.ExpectedLiquidity, token.Decimals),
		// ExpectedLiquidityLimitBI: (*core.BigInt)(state.ExpectedLiquidityLimit),

		AvailableLiquidityBI: (*core.BigInt)(state.AvailableLiquidity),
		AvailableLiquidity:   utils.GetFloat64Decimal(state.AvailableLiquidity, token.Decimals),

		DepositAPYBI: (*core.BigInt)(state.SupplyRate),
		// for 4% is depositAPY is 4 that is why apy is divided by decimal 25 not 27
		DepositAPY: utils.GetFloat64Decimal(state.SupplyRate, 25),

		BaseBorrowAPYBI: (*core.BigInt)(state.BaseInterestRate),
		BaseBorrowAPY:   utils.GetFloat64Decimal(state.BaseInterestRate, 25),

		// dieselrate is how much each diesel rate is worth in terms of underlying token
		// that's why it is divide by 27 not 25. it is not a percentage.
		DieselRateBI:       (*core.BigInt)(state.DieselRateRAY),
		DieselRate:         utils.GetFloat64Decimal(state.DieselRateRAY, 27),
		WithdrawFee:        int(state.WithdrawFee.Convert().Int64()),
		CumulativeIndexRAY: (*core.BigInt)(state.CumulativeIndexRAY),
	})

}
