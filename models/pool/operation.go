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

	state, err := p.Repo.GetDataCompressor(blockNum).GetPoolData(opts, common.HexToAddress(p.Address))
	if err != nil {
		log.Fatal("[PoolService] Cant get data from data compressor", err)
		return
	}
	token:=p.Repo.GetToken(p.Repo.GetPool(p.Address).UnderlyingToken)
	p.Repo.AddPoolStat(&core.PoolStat{
		BlockNum: blockNum,
		Address: p.Address,
		UniqueUsers: p.Repo.GetPoolUniqueUserLen(p.Address),
		TotalBorrowedBI: (*core.BigInt)(state.TotalBorrowed),
		TotalBorrowed: utils.GetFloat64Decimal(state.TotalBorrowed, token.Decimals),

		ExpectedLiquidityBI: (*core.BigInt)(state.ExpectedLiquidity),
		ExpectedLiquidity: utils.GetFloat64Decimal(state.ExpectedLiquidity, token.Decimals),

		AvailableLiquidityBI: (*core.BigInt)(state.AvailableLiquidity),
		AvailableLiquidity: utils.GetFloat64Decimal(state.AvailableLiquidity, token.Decimals),

		DepositAPYBI: (*core.BigInt)(state.DepositAPYRAY),
		DepositAPY: utils.GetFloat64Decimal(state.DepositAPYRAY, 25),

		BorrowAPYBI: (*core.BigInt)(state.BorrowAPYRAY),
		BorrowAPY: utils.GetFloat64Decimal(state.BorrowAPYRAY, 25),

		DieselRateBI: (*core.BigInt)(state.DieselRateRAY),
		DieselRate: utils.GetFloat64Decimal(state.DieselRateRAY, 25),
		WithdrawFee: int(state.WithdrawFee.Int64()),
		ID: 0,
	})
}