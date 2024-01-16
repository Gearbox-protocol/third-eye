package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *LMRewardsv3) updateFarmedPerToken(farmAddr string, currentTs uint64) {
	// sol:updateFarmedPerToken
	farm := mdl.farms[farmAddr]
	farm.Checkpoint = currentTs
	farm.Fpt = (*core.BigInt)(farm.calcFarmedPerToken(currentTs))
}

func (mdl *LMRewardsv3) performTransfer(farmAddr, from, to string, amount *big.Int) {
	fromZero := from == core.NULL_ADDR.Hex()
	toZero := to == core.NULL_ADDR.Hex()
	//
	if fromZero || toZero {
		diff := amount
		if toZero {
			diff = new(big.Int).Neg(diff)
		}
		mdl.farms[farmAddr].TotalSupply = (*core.BigInt)(new(big.Int).Add(
			mdl.farms[farmAddr].TotalSupply.Convert(),
			diff,
		))
	}
	//
	diesel := mdl.Repo.GetToken(mdl.farms[farmAddr].DieselToken)
	if to == "0x2E67A94b39c1946D100D85Ba724c116a652515B9" {
		log.Info("", diesel.Decimals)
	}

	if mdl.users[common.HexToAddress(farmAddr)] == nil {
		mdl.users[common.HexToAddress(farmAddr)] = map[string]*UserLMDetails{}
	}
	farmAndItsUsers := mdl.users[common.HexToAddress(farmAddr)]
	if !fromZero {
		if farmAndItsUsers[from] == nil {
			farmAndItsUsers[from] = &UserLMDetails{
				Correction: (*core.BigInt)(new(big.Int)),
				BalancesBI: (*core.BigInt)(new(big.Int)),
				Account:    from,
				Farm:       farmAddr,
				DieselSym:  diesel.Symbol,
			}
		}
		farmAndItsUsers[from].SubBalances(amount, diesel.Decimals)
	}
	if !toZero {
		if farmAndItsUsers[to] == nil {
			farmAndItsUsers[to] = &UserLMDetails{
				Correction: (*core.BigInt)(new(big.Int)),
				BalancesBI: (*core.BigInt)(new(big.Int)),
				Account:    to,
				Farm:       farmAddr,
				DieselSym:  diesel.Symbol,
			}
		}
		farmAndItsUsers[to].AddBalances(amount, diesel.Decimals)
	}
}
func (mdl *LMRewardsv3) updateBalances(farmAddr, from, to string, amount *big.Int, currentTs uint64) {
	//
	mdl.performTransfer(farmAddr, from, to, amount)
	//
	fromZero := from == core.NULL_ADDR.Hex()
	toZero := to == core.NULL_ADDR.Hex()

	farm := mdl.farms[farmAddr]
	farmAndItsUsers := mdl.users[common.HexToAddress(farmAddr)]
	if amount.Sign() > 0 && from != to {
		if fromZero || toZero {
			mdl.updateFarmedPerToken(farmAddr, currentTs)
		}

		//
		diff := new(big.Int).Mul(amount, farm.calcFarmedPerToken(currentTs))
		if !fromZero {
			farmAndItsUsers[from].SubCorrection(diff)
		}
		if !toZero {
			farmAndItsUsers[to].AddCorrection(diff)
		}
	}
}
