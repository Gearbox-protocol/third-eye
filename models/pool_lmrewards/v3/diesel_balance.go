package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *LMRewardsv3) updateDieselBalances(txLog types.Log, pool common.Address, from, to string, amount *big.Int) {
	fromZero := from == core.NULL_ADDR.Hex()
	toZero := to == core.NULL_ADDR.Hex()

	if amount.Sign() > 0 && from != to {
		//
		if mdl.dieselBalances[pool] == nil {
			mdl.dieselBalances[pool] = map[string]*ds.DieselBalance{}
		}
		//
		userAndBalance := mdl.dieselBalances[pool]
		decimals := mdl.Repo.GetToken(pool.Hex()).Decimals
		if !fromZero {
			if userAndBalance[from] == nil {
				userAndBalance[from] = &ds.DieselBalance{
					Pool:      pool.Hex(),
					User:      from,
					BalanceBI: (*core.BigInt)(new(big.Int)),
				}
			}
			userAndBalance[from].Updated = true
			if userAndBalance[from].BalanceBI == nil {
				log.Fatal("Pool, from, to, amount txHash", pool, from, to, amount, txLog.TxHash, txLog.Index)
			}
			userAndBalance[from].BalanceBI = (*core.BigInt)(new(big.Int).Sub(
				userAndBalance[from].BalanceBI.Convert(),
				amount,
			))
			userAndBalance[from].Balance = utils.GetFloat64Decimal(userAndBalance[from].BalanceBI, decimals)
		}
		if !toZero {
			if userAndBalance[to] == nil {
				userAndBalance[to] = &ds.DieselBalance{
					User:      to,
					BalanceBI: (*core.BigInt)(new(big.Int)),
					Pool:      pool.Hex(),
				}
			}
			userAndBalance[to].Updated = true
			userAndBalance[to].BalanceBI = core.AddCoreAndInt(userAndBalance[to].BalanceBI, amount)
			userAndBalance[to].Balance = utils.GetFloat64Decimal(userAndBalance[to].BalanceBI, decimals)
		}
	}
}
