package pool_lmrewards

import (
	"log"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//
func (mdl *PoolLMRewards) GetAddrsForLogs() (addrs []common.Address) {
	addrs = append(addrs, common.HexToAddress(mdl.Address))
	return addrs
}

//

//
func (mdl *PoolLMRewards) OnLog(txLog types.Log) {
	currentBlock := int64(txLog.BlockNumber)
	if currentBlock != mdl.lastBlockNum {
		mdl.calculateRewards(mdl.lastBlockNum+1, currentBlock)
		mdl.lastBlockNum = currentBlock
	}
	//
	switch txLog.Topics[0] {
	case core.Topic("Transfer(address,address,uint256)"):
		token := mdl.Repo.GetToken(txLog.Address.Hex())
		tokenSym := token.Symbol
		from := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		to := common.HexToAddress(txLog.Topics[2].Hex()).Hex()
		amount, ok := new(big.Int).SetString(common.BytesToHash(txLog.Data).Hex()[2:], 16)
		if !ok {
			log.Fatal("Failed parsing value")
		}
		if to == core.NULL_ADDR.Hex() {
			mdl.totalSupplies[tokenSym] = new(big.Int).Sub(
				utils.NotNilBigInt(mdl.totalSupplies[tokenSym]),
				amount,
			)
		} else {
			mdl.addBalance(tokenSym, to, amount)
		}

		if from == core.NULL_ADDR.Hex() {
			mdl.totalSupplies[tokenSym] = new(big.Int).Add(
				utils.NotNilBigInt(mdl.totalSupplies[tokenSym]),
				amount,
			)
		} else {
			mdl.addBalance(tokenSym, from, amount.Neg(nil))
		}
	}
}

func (mdl PoolLMRewards) addBalance(tokenSym, user string, amount *big.Int) {
	if mdl.dieselBalances[tokenSym] == nil {
		mdl.dieselBalances[tokenSym] = map[string]*big.Int{}
	}
	mdl.dieselBalances[tokenSym][user] = new(big.Int).Add(
		utils.NotNilBigInt(mdl.dieselBalances[tokenSym][user]),
		amount,
	)
}

func (mdl PoolLMRewards) calculateRewards(from, to int64) {
	snapshots := core.GetRewardPerToken(mdl.chainId, from, to)

	snapStart := from
	if len(snapshots) > 0 && from < snapshots[0].Block {
		snapStart = snapshots[0].Block
	}
	for snapInd, snapshot := range snapshots {
		for dieselSym, userAndbalance := range mdl.dieselBalances {
			rewardPerBlock := utils.GetInt64(snapshot.RewardPerBlock[dieselSym], mdl.decimals[dieselSym])
			for user, balance := range userAndbalance {
				norm := new(big.Int).Mul(balance, rewardPerBlock)
				userRewardPerBlock := new(big.Int).Quo(norm, mdl.totalSupplies[dieselSym])

				snapEnd := to
				if snapInd != len(snapshots)-1 {
					snapEnd = snapshots[snapInd+1].Block - 1
				}
				reward := new(big.Int).Mul(userRewardPerBlock, big.NewInt(snapEnd-snapStart+1))
				mdl.rewards[user] = new(big.Int).Add(
					utils.NotNilBigInt(mdl.rewards[user]),
					reward,
				)
				// update start
				snapStart = snapEnd + 1
			}
		}
	}
}

//
func (mdl *PoolLMRewards) GetOtherAddrsForLogs() (addrs []common.Address) {
	for addr := range mdl.Repo.GetDieselTokens() {
		addrs = append(addrs, common.HexToAddress(addr))
		token := mdl.Repo.GetToken(addr)
		mdl.decimals[token.Symbol] = token.Decimals
	}
	return addrs
}
