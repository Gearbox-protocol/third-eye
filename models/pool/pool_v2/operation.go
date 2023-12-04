package pool_v2

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *Poolv2) fixPoolLedgerAddrForGateway() {
	// for remove liquidity
	for _, removeLiqEvent := range mdl.gatewayHandler.GetRemoveLiqEventsAndClear() {
		// calculate removed liquidity amount in underlying token
		numerator := new(big.Int).Mul(removeLiqEvent.AmountBI.Convert(), mdl.dieselRate)
		underlyingRemovedAmount := new(big.Int).Quo(numerator, utils.GetExpInt(27))
		// set removed  amount fields in poolLedger
		removeLiqEvent.AmountBI = (*core.BigInt)(underlyingRemovedAmount)
		removeLiqEvent.Amount = utils.GetFloat64Decimal(underlyingRemovedAmount, mdl.Repo.GetToken(mdl.State.UnderlyingToken).Decimals)
		// add poolLedger to repository
		mdl.Repo.AddPoolLedger(removeLiqEvent)
	}
}

func (mdl *Poolv2) OnBlockChange(inputBlock int64) (call multicall.Multicall2Call, processFn func(multicall.Multicall2Result)) {
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

func (mdl *Poolv2) getCallAndProcessFn(inputB int64) (multicall.Multicall2Call, func(multicall.Multicall2Result)) {
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetPoolData(mdl.GetVersion(), inputB, common.HexToAddress(mdl.Address))
	if err != nil {
		log.Fatal("[PoolService] Cant create call for data compressor", err)
	}
	return call, func(result multicall.Multicall2Result) {
		if !result.Success {
			log.Fatal("[PoolService] Cant process result for data compressor", err)
		}
		poolState, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatal("[PoolService] Cant process result for data compressor", err)
		}
		//
		mdl.createSnapshot(inputB, poolState)
		// it uses dieselRate from create snapshot so this call should be after createSnapshots
		mdl.fixPoolLedgerAddrForGateway()
	}
}

func (mdl *Poolv2) onBlockChangeInternally(inputB int64) {
	call, processFn := mdl.getCallAndProcessFn(inputB)
	result := core.MakeMultiCall(mdl.Client, inputB, false, []multicall.Multicall2Call{call})
	processFn(result[0])
}

func (p *Poolv2) createSnapshot(blockNum int64, state dc.PoolCallData) {
	token := p.Repo.GetToken(p.State.UnderlyingToken)
	p.State.IsWETH = dc.IsWETH(p.Client, state.Underlying)
	// TODO: change borrow apy
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
		// TODO change to base borrow rate
		BaseBorrowAPY: utils.GetFloat64Decimal(state.BaseInterestRate, 25),

		// dieselrate is how much each diesel rate is worth in terms of underlying token
		// that's why it is divide by 27 not 25. it is not a percentage.
		DieselRateBI:       (*core.BigInt)(state.DieselRateRAY),
		DieselRate:         utils.GetFloat64Decimal(state.DieselRateRAY, 27),
		WithdrawFee:        int(state.WithdrawFee.Convert().Int64()),
		CumulativeIndexRAY: (*core.BigInt)(state.CumulativeIndexRAY),
	})
	p.dieselRate = state.DieselRateRAY.Convert()
}
