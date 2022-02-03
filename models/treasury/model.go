package treasury

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"

	//
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Treasury struct {
	*core.SyncAdapter
	node *core.Node
}

func NewTreasury(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *Treasury {
	syncAdapter := &core.SyncAdapter{
		Contract: &core.Contract{
			Address:      addr,
			DiscoveredAt: discoveredAt,
			FirstLogAt:   discoveredAt,
			ContractName: core.Treasury,
			Client:       client,
		},
		LastSync: discoveredAt - 1,
		Repo:     repo,
	}
	return NewTreasuryFromAdapter(
		syncAdapter,
	)
}

func NewTreasuryFromAdapter(adapter *core.SyncAdapter) *Treasury {
	obj := &Treasury{
		SyncAdapter: adapter,
	}
	obj.OnlyQuery = true
	obj.node = &core.Node{
		Client: adapter.Client,
	}
	return obj
}

func (mdl *Treasury) OnLog(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("Transfer(address,address,uint256)"):
		from := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		to := common.HexToAddress(txLog.Topics[2].Hex()).Hex()
		value, ok := new(big.Int).SetString(common.BytesToHash(txLog.Data).Hex()[2:], 16)
		if !ok {
			log.Fatal("Failed parsing value")
		}
		switch mdl.Address {
		case from:
			mdl.Repo.AddTreasuryTransfer(int64(txLog.BlockNumber), txLog.Index,
				txLog.Address.Hex(), new(big.Int).Neg(value))
		case to:
			mdl.Repo.AddTreasuryTransfer(int64(txLog.BlockNumber), txLog.Index,
				txLog.Address.Hex(), value)
		}
	}
}

func (mdl *Treasury) Query(queryTill int64) {
	queryFrom := mdl.GetLastSync() + 1
	tokenAddrs := mdl.Repo.GetTokens()
	hexAddrs := []common.Address{}

	treasuryAddrTopic := []common.Hash{
		common.HexToHash(mdl.Address),
	}
	for _, tokenAddr := range tokenAddrs {
		hexAddrs = append(hexAddrs, common.HexToAddress(tokenAddr))
	}
	logs, err := mdl.node.GetLogsForTransfer(queryFrom, queryTill, hexAddrs, treasuryAddrTopic)
	log.CheckFatal(err)
	for _, log := range logs {
		mdl.OnLog(log)
	}
}

func (mdl *Treasury) AfterSyncHook(syncTill int64) {
	// for treasury/dao
	mdl.Repo.CalCurrentTreasuryValue(syncTill)
	mdl.SyncAdapter.AfterSyncHook(syncTill)
}
