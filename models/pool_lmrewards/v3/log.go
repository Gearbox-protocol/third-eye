package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// https://etherscan.io/address/0x9ef444a6d7f4a5adcd68fd5329aa5240c90e14d2#code farmingPool
func (mdl LMRewardsv3) OnLog(txLog types.Log) {
	farmAddr := txLog.Address.Hex()
	currentTs := mdl.Repo.SetAndGetBlock(int64(txLog.BlockNumber)).Timestamp
	switch txLog.Topics[0] {
	case core.Topic("Transfer(address,address,uint256)"):
		from := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		to := common.BytesToAddress(txLog.Topics[2][:]).Hex()
		amount := new(big.Int).SetBytes(txLog.Data)
		mdl.updateBalances(farmAddr, from, to, amount, currentTs)
	case core.Topic("RewardUpdated(uint256,uint256)"):
		// sol:updateFarmedPerToken
		mdl.updateFarmedPerToken(farmAddr, currentTs)
		// farmInfo.startFarming
		farm := mdl.farms[farmAddr]
		reward := new(big.Int).SetBytes(txLog.Data[:32])
		period := new(big.Int).SetBytes(txLog.Data[32:])
		if period.Int64() == 0 { // startFarming can't have 0 period as it reverts in sol
			farm.stopFarming(reward, currentTs)
		} else {
			farm.startFarming(reward, period.Uint64(), currentTs)
		}
	}
}

func (mdl *LMRewardsv3) getFarmsv3() {
	if len(mdl.farms) != 0 { // already set
		return
	}
	pools, found := mdl.Repo.GetDCWrapper().GetPoolListv3()
	if found && len(mdl.farms) == 0 {
		addrToSym := core.GetAddrToSymbolByChainId(core.GetChainId(mdl.Client))
		poolAndFarms := []*Farmv3{}
		for _, pool := range pools {
			for _, zapper := range pool.Zappers {
				if _, ok := addrToSym[zapper.TokenOut]; !ok && zapper.TokenIn == pool.Underlying {
					poolAndFarms = append(poolAndFarms, &Farmv3{
						Farm:        zapper.TokenOut.Hex(),
						Pool:        pool.Addr.Hex(),
						DieselToken: pool.DieselToken.Hex(),
						// initial
						Fpt:         (*core.BigInt)(new(big.Int)),
						TotalSupply: (*core.BigInt)(new(big.Int)),
						Reward:      (*core.BigInt)(new(big.Int)),
					})
				}
			}
		}
		mdl.SetUnderlyingState(poolAndFarms)
	}
}

// LMRewardsv2 has fake address so no need for adding .Address value to addrs
func (mdl *LMRewardsv3) GetAllAddrsForLogs() (addrs []common.Address) {
	mdl.getFarmsv3()
	//
	for addr := range mdl.farms {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	addrs = append(addrs, common.HexToAddress(mdl.Address))
	return addrs
}
