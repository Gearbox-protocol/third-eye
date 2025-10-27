package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ss struct {
	Amount *big.Int
	IsFrom bool
	TxHash common.Hash
	Total  *big.Int
}

func GetLogs() func() ss {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)
	node := pkg.Node{Client: client}
	addr := "0x3E965117A51186e41c2BB58b729A1e518A715e5F" // nikitakle address on mainnet
	log.Info(core.GetChainId(client))
	txlogs, err := node.GetLogsForTransfer(0, node.GetLatestBlockNumber(), []common.Address{common.HexToAddress("0xda0002859B2d05F66a753d8241fCDE8623f26F4f")}, []common.Hash{common.HexToHash(addr)})
	log.CheckFatal(err)
	total := new(big.Int)
	ans := []ss{}
	for _, txLog := range txlogs {
		from := common.BytesToAddress(txLog.Topics[1][:])
		to := common.BytesToAddress(txLog.Topics[2][:])
		amount := new(big.Int).SetBytes(txLog.Data)
		log.Info(to.Hex() == addr, amount)
		if from.Hex() == addr {
			amount = new(big.Int).Neg(amount)
		}
		total = new(big.Int).Add(total, amount)
		ans = append(ans, ss{Amount: new(big.Int).SetBytes(txLog.Data), IsFrom: from.Hex() == addr, TxHash: txLog.TxHash, Total: total})

	}
	log.Info(total)
	ind := 0
	return func() ss {
		if ind < len(ans) {
			ind++
			return ans[ind-1]
		}
		return ss{}
	}
}

func (a ss) Same(b ss) bool {
	return a.Amount.Cmp(b.Amount) == 0 && a.IsFrom == b.IsFrom && a.TxHash == b.TxHash && a.Total.Cmp(b.Total) == 0
}

var fn func() ss

func init() {
	// fn =GetLogs()
}

// https://etherscan.io/address/0x9ef444a6d7f4a5adcd68fd5329aa5240c90e14d2#code farmingPool
func (mdl LMRewardsv3) OnLog(txLog types.Log) {
	addr := txLog.Address.Hex()
	blockNum := int64(txLog.BlockNumber)
	//
	if mdl.farms[addr] != nil {
		farmAddr := addr
		if mdl.farms[addr] != nil && mdl.farms[addr].FarmSyncedTill >= blockNum { // if farm
			return
		}
		currentTs := mdl.Repo.SetAndGetBlock(blockNum).Timestamp
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
			tokenDetails := mdl.Repo.GetToken(poolAddr.Hex())
			mdl.Repo.AddDieselTransfer(&schemas.DieselTransfer{
				LogId:       int64(txLog.Index),
				BlockNum:    int64(txLog.BlockNumber),
				TokenSymbol: tokenDetails.Symbol,
				From:        from,
				To:          to,
				Amount:      utils.GetFloat64Decimal(amount, tokenDetails.Decimals),
			})
			mdl.updateDieselBalances(txLog, poolAddr, from, to, amount)
			// if to == "0x3E965117A51186e41c2BB58b729A1e518A715e5F" || from == "0x3E965117A51186e41c2BB58b729A1e518A715e5F" {
			// 	current :=ss{
			// 		Amount: amount,
			// 		IsFrom: from == "0x3E965117A51186e41c2BB58b729A1e518A715e5F",
			// 		TxHash: txLog.TxHash,
			// 		Total:  mdl.dieselBalances[poolAddr]["0x3E965117A51186e41c2BB58b729A1e518A715e5F"].BalanceBI.Convert(),

			// 		}
			// log.Info(utils.ToJson(current))
			// if poolAddr == common.HexToAddress("0xda0002859B2d05F66a753d8241fCDE8623f26F4f") {
			// 	expected := fn()
			// 	if !expected.Same(current) {
			// 		log.Fatal("expected", utils.ToJson(expected), "current", utils.ToJson(current))
			// 	}
			// }
			// }
		}
	}
}

// over right to set new farms
func (mdl *LMRewardsv3) GetLastSync() int64 {
	lastSync := mdl.SyncAdapter.GetLastSync()
	mdl.getFarmsAndPoolsv3(lastSync)
	return mdl.SyncAdapter.LastSync
}

// LMRewardsv2 has fake address so no need for adding .Address value to addrs
func (mdl *LMRewardsv3) GetAllAddrsForLogs() (addrs []common.Address) {
	//
	addrs = append(addrs, common.HexToAddress(mdl.GetAddress())) // if no pools then no addresses were returned previous as a result the lastsync wasn't getting updated for lmrewardsv3
	for farmAddr := range mdl.farms {
		addrs = append(addrs, common.HexToAddress(farmAddr))
	}
	for pool := range mdl.poolsToSyncedTill {
		addrs = append(addrs, pool)
	}
	return addrs
}
