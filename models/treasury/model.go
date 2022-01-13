package treasury

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"

	//
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
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
		value, ok := new(big.Int).SetString(txLog.Topics[1].Hex()[2:], 16)
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

func (mdl *Treasury) Query(queryTill int64, wg *sync.WaitGroup) {
	defer wg.Done()
	queryFrom := mdl.GetLastSync() + 1
	if queryFrom > queryTill {
		return
	}
	tokenAddrs := mdl.Repo.GetTokens()
	hexAddrs := []common.Address{}

	topics := [][]common.Hash{
		{
			core.Topic("Transfer(address,address,uint256)"),
		},
		{
			common.HexToHash(mdl.Address),
		},
		{
			common.HexToHash(mdl.Address),
		},
	}
	for _, tokenAddr := range tokenAddrs {
		hexAddrs = append(hexAddrs, common.HexToAddress(tokenAddr))
	}

	logs, err := mdl.node.GetLogs(queryFrom, queryTill, hexAddrs, topics)
	log.CheckFatal(err)
	for _, log := range logs {
		mdl.OnLog(log)
	}
	// after sync
	mdl.AfterSyncHook(queryTill)
}
