package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *LMRewardsv3) updateFarmedPerToken(farmAddr string, currentTs uint64, blockNum int64) {
	// sol:updateFarmedPerToken
	farm := mdl.farms[farmAddr]
	//
	farm.Fpt = (*core.BigInt)(farm.calcFarmedPerToken(currentTs))
	farm.Checkpoint = currentTs
}

func (mdl *LMRewardsv3) performTransfer(farmAddr, from, to string, amount *big.Int, blockNum int64) {
	fromZero := from == core.NULL_ADDR.Hex()
	toZero := to == core.NULL_ADDR.Hex()
	//
	if from == to {
		return
	}
	if fromZero || toZero {
		diff := amount
		if toZero {
			diff = new(big.Int).Neg(diff)
		}
		mdl.farms[farmAddr].TotalSupply = (*core.BigInt)(new(big.Int).Add(
			mdl.farms[farmAddr].TotalSupply.Convert(),
			diff,
		))
		// totalSupplyData, err := core.CallFuncWithExtraBytes(mdl.Client, "18160ddd", common.HexToAddress(farmAddr), blockNum, nil)
		// log.CheckFatal(err)
		// totalSupply := new(big.Int).SetBytes(totalSupplyData)
		// if totalSupply.Cmp(mdl.farms[farmAddr].TotalSupply.Convert()) != 0 {
		// 	log.Fatal("Total supply mismatch", totalSupply, mdl.farms[farmAddr].TotalSupply.Convert())
		// }
	}
	//
	diesel := mdl.Repo.GetToken(mdl.farms[farmAddr].DieselToken)
	farmAndItsUsers := mdl.farmUserRewards[common.HexToAddress(farmAddr)]
	if !fromZero {
		farmAndItsUsers[from].SubBalances(amount, diesel.Decimals)
	}
	if !toZero {
		farmAndItsUsers[to].AddBalances(amount, diesel.Decimals)
	}
}
func (mdl *LMRewardsv3) updateBalances(farmAddr, from, to string, amount *big.Int, currentTs uint64, blockNum int64) {
	//
	fromZero := from == core.NULL_ADDR.Hex()
	toZero := to == core.NULL_ADDR.Hex()

	if amount.Sign() > 0 && from != to {
		if fromZero || toZero {
			mdl.updateFarmedPerToken(farmAddr, currentTs, blockNum)
		}

		//
		if mdl.farmUserRewards[common.HexToAddress(farmAddr)] == nil {
			mdl.farmUserRewards[common.HexToAddress(farmAddr)] = map[string]*UserLMDetails{}
		}
		//
		farm := mdl.farms[farmAddr]
		diff := new(big.Int).Mul(amount, farm.calcFarmedPerToken(currentTs))
		//
		farmAndItsUsers := mdl.farmUserRewards[common.HexToAddress(farmAddr)]
		if !fromZero {
			if farmAndItsUsers[from] == nil {
				farmAndItsUsers[from] = &UserLMDetails{
					Correction:      (*core.BigInt)(new(big.Int)),
					FarmedBalanceBI: (*core.BigInt)(new(big.Int)),
					Account:         from,
					Farm:            farmAddr,
				}
			}
			farmAndItsUsers[from].SubCorrection(diff)
		}
		if !toZero {
			if farmAndItsUsers[to] == nil {
				farmAndItsUsers[to] = &UserLMDetails{
					Correction:      (*core.BigInt)(new(big.Int)),
					FarmedBalanceBI: (*core.BigInt)(new(big.Int)),
					Account:         to,
					Farm:            farmAddr,
				}
			}
			farmAndItsUsers[to].AddCorrection(diff)
		}
		mdl.performTransfer(farmAddr, from, to, amount, blockNum)
	}
	//
}
