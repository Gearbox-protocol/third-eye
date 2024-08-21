package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// https://etherscan.io/address/0x9ef444a6d7f4a5adcd68fd5329aa5240c90e14d2#code farmingPool
func (mdl LMRewardsv3) OnLog(txLog types.Log) {
	addr := txLog.Address.Hex()
	blockNum := int64(txLog.BlockNumber)
	currentTs := mdl.Repo.SetAndGetBlock(blockNum).Timestamp
	//
	if mdl.farms[addr] != nil {
		farmAddr := addr
		if mdl.farms[addr] != nil && mdl.farms[addr].SyncedTill >= blockNum { // if farm
			return
		}
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
		poolAddr := txLog.Address
		if core.Topic("Transfer(address,address,uint256)") == txLog.Topics[0] {
			from := common.BytesToAddress(txLog.Topics[1][:]).Hex()
			to := common.BytesToAddress(txLog.Topics[2][:]).Hex()
			amount := new(big.Int).SetBytes(txLog.Data)
			if lastSync := mdl.poolsToSyncedTill[poolAddr]; lastSync >= blockNum {
				return
			}
			// if to == "0x3E965117A51186e41c2BB58b729A1e518A715e5F" || from == "0x3E965117A51186e41c2BB58b729A1e518A715e5F" { 
			// 	log.Info(from == "0x3E965117A51186e41c2BB58b729A1e518A715e5F" , amount, txLog.TxHash, mdl.poolsToSyncedTill[poolAddr])
			// }
			mdl.updateDieselBalances(txLog, poolAddr, from, to, amount)
		}
	}
}

// over right to set new farms
func (mdl *LMRewardsv3) GetLastSync() int64 {
	mdl.getFarmsAndPoolsv3()
	return mdl.SyncAdapter.GetLastSync()
}

// LMRewardsv2 has fake address so no need for adding .Address value to addrs
func (mdl *LMRewardsv3) GetAllAddrsForLogs() (addrs []common.Address) {
	//
	for addr, farm := range mdl.farms {
		addrs = append(addrs, common.HexToAddress(addr))
		addrs = append(addrs, common.HexToAddress(farm.Pool))
	}
	return addrs
}
