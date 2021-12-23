package pool

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (p *Pool) calculatePoolStat(blockNum int64) {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}

	state, err := p.Repo.GetDCWrapper().GetPoolData(opts, common.HexToAddress(p.Address))
	if err != nil {
		log.Fatal("[PoolService] Cant get data from data compressor", err)
		return
	}
	token := p.Repo.GetToken(p.State.UnderlyingToken)
	p.State.IsWETH = state.IsWETH
	// log.Infof("Pool:%s ciRAY: %s and linearCI: %s\n", p.Address, state.CumulativeIndexRAY.String(), state.LinearCumulativeIndex.String())
	p.Repo.AddPoolStat(&core.PoolStat{
		BlockNum:        blockNum,
		Address:         p.Address,
		UniqueUsers:     p.Repo.GetPoolUniqueUserLen(p.Address),
		TotalBorrowedBI: (*core.BigInt)(state.TotalBorrowed),
		TotalBorrowed:   utils.GetFloat64Decimal(state.TotalBorrowed, token.Decimals),

		ExpectedLiquidityBI: (*core.BigInt)(state.ExpectedLiquidity),
		ExpectedLiquidity:   utils.GetFloat64Decimal(state.ExpectedLiquidity, token.Decimals),
		ExpectedLiquidityLimitBI: (*core.BigInt)(state.ExpectedLiquidityLimit),

		AvailableLiquidityBI: (*core.BigInt)(state.AvailableLiquidity),
		AvailableLiquidity:   utils.GetFloat64Decimal(state.AvailableLiquidity, token.Decimals),

		DepositAPYBI: (*core.BigInt)(state.DepositAPYRAY),
		DepositAPY:   utils.GetFloat64Decimal(state.DepositAPYRAY, 27),

		BorrowAPYBI: (*core.BigInt)(state.BorrowAPYRAY),
		BorrowAPY:   utils.GetFloat64Decimal(state.BorrowAPYRAY, 27),

		DieselRateBI:       (*core.BigInt)(state.DieselRateRAY),
		DieselRate:         utils.GetFloat64Decimal(state.DieselRateRAY, 27),
		WithdrawFee:        int(state.WithdrawFee.Int64()),
		CumulativeIndexRAY: (*core.BigInt)(state.LinearCumulativeIndex),
		ID:                 0,
	})
}
