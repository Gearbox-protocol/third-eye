package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// https://etherscan.io/address/0x9ef444a6d7f4a5adcd68fd5329aa5240c90e14d2#code farmingPool
func (mdl LMRewardsv3) OnLog(txLog types.Log) {
	addr := txLog.Address
	blockNum := int64(txLog.BlockNumber)
	currentTs := mdl.Repo.SetAndGetBlock(blockNum).Timestamp
	if mdl.farms[addr.Hex()] != nil {
		farmAddr := addr.Hex()
		switch txLog.Topics[0] {
		case core.Topic("Transfer(address,address,uint256)"):
			from := common.BytesToAddress(txLog.Topics[1][:]).Hex()
			to := common.BytesToAddress(txLog.Topics[2][:]).Hex()
			amount := new(big.Int).SetBytes(txLog.Data)
			mdl.updateBalances(farmAddr, from, to, amount, currentTs, blockNum)
			mdl.check(blockNum, currentTs, farmAddr, from, to)
		case core.Topic("RewardUpdated(uint256,uint256)"):
			// sol:updateFarmedPerToken
			mdl.updateFarmedPerToken(farmAddr, currentTs, blockNum)
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
	} else { // for tracking diesel balance
		poolAddr := addr
		if core.Topic("Transfer(address,address,uint256)") == txLog.Topics[0] {
			from := common.BytesToAddress(txLog.Topics[1][:]).Hex()
			to := common.BytesToAddress(txLog.Topics[2][:]).Hex()
			amount := new(big.Int).SetBytes(txLog.Data)
			farmAddr := mdl.pools[poolAddr]
			mdl.updateDieselBalances(farmAddr, from, to, amount)
		}
	}
}

// LMRewardsv2 has fake address so no need for adding .Address value to addrs
func (mdl *LMRewardsv3) GetAllAddrsForLogs() (addrs []common.Address) {
	mdl.getFarmsAndPoolsv3()
	//
	for addr, farm := range mdl.farms {
		addrs = append(addrs, common.HexToAddress(addr))
		addrs = append(addrs, common.HexToAddress(farm.Pool))
	}
	return addrs
}
