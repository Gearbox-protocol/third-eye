package pool

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	p.Repo.AddPoolStat(&schemas.PoolStat{
		BlockNum:        blockNum,
		Address:         p.Address,
		UniqueUsers:     p.Repo.GetPoolUniqueUserLen(p.Address),
		TotalBorrowedBI: (*core.BigInt)(state.TotalBorrowed),
		TotalBorrowed:   utils.GetFloat64Decimal(state.TotalBorrowed, token.Decimals),

		ExpectedLiquidityBI:      (*core.BigInt)(state.ExpectedLiquidity),
		ExpectedLiquidity:        utils.GetFloat64Decimal(state.ExpectedLiquidity, token.Decimals),
		ExpectedLiquidityLimitBI: (*core.BigInt)(state.ExpectedLiquidityLimit),

		AvailableLiquidityBI: (*core.BigInt)(state.AvailableLiquidity),
		AvailableLiquidity:   utils.GetFloat64Decimal(state.AvailableLiquidity, token.Decimals),

		DepositAPYBI: (*core.BigInt)(state.DepositAPYRAY),
		// for 4% is depositAPY is 4 that is why apy is divided by decimal 25 not 27
		DepositAPY: utils.GetFloat64Decimal(state.DepositAPYRAY, 25),

		BorrowAPYBI: (*core.BigInt)(state.BorrowAPYRAY),
		BorrowAPY:   utils.GetFloat64Decimal(state.BorrowAPYRAY, 25),

		// dieselrate is how much each diesel rate is worth in terms of underlying token
		// that's why it is divide by 27 not 25. it is not a percentage.
		DieselRateBI:       (*core.BigInt)(state.DieselRateRAY),
		DieselRate:         utils.GetFloat64Decimal(state.DieselRateRAY, 27),
		WithdrawFee:        int(state.WithdrawFee.Int64()),
		CumulativeIndexRAY: (*core.BigInt)(state.LinearCumulativeIndex),
	})
	p.dieselRate = state.DieselRateRAY
}
