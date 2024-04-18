package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *LMRewardsv3) updateDieselBalances(farmAddr, from, to string, amount *big.Int) {
	fromZero := from == core.NULL_ADDR.Hex()
	toZero := to == core.NULL_ADDR.Hex()

	if amount.Sign() > 0 && from != to {
		//
		if mdl.users[common.HexToAddress(farmAddr)] == nil {
			mdl.users[common.HexToAddress(farmAddr)] = map[string]*UserLMDetails{}
		}
		//
		diesel := mdl.Repo.GetToken(mdl.farms[farmAddr].DieselToken)
		farmAndItsUsers := mdl.users[common.HexToAddress(farmAddr)]
		if !fromZero {
			if farmAndItsUsers[from] == nil {
				farmAndItsUsers[from] = &UserLMDetails{
					Correction:    (*core.BigInt)(new(big.Int)),
					BalancesBI:    (*core.BigInt)(new(big.Int)),
					Account:       from,
					Farm:          farmAddr,
					DieselSym:     diesel.Symbol,
					DieselBalance: (*core.BigInt)(new(big.Int)),
				}
			}
			farmAndItsUsers[from].updated = true
			farmAndItsUsers[from].DieselBalance = (*core.BigInt)(new(big.Int).Sub(
				farmAndItsUsers[from].DieselBalance.Convert(),
				amount,
			))
		}
		if !toZero {
			if farmAndItsUsers[to] == nil {
				farmAndItsUsers[to] = &UserLMDetails{
					Correction:    (*core.BigInt)(new(big.Int)),
					BalancesBI:    (*core.BigInt)(new(big.Int)),
					Account:       to,
					Farm:          farmAddr,
					DieselSym:     diesel.Symbol,
					DieselBalance: (*core.BigInt)(new(big.Int)),
				}
			}
			farmAndItsUsers[from].updated = true
			farmAndItsUsers[to].DieselBalance = (*core.BigInt)(new(big.Int).Add(
				farmAndItsUsers[to].DieselBalance.Convert(),
				amount,
			))
		}
	}
}
