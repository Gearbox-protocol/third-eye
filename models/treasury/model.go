package treasury

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"

	//
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Treasury struct {
	*ds.SyncAdapter
	node    *core.Node
	HexAddr common.Address
}

func NewTreasury(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *Treasury {
	syncAdapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address:      addr,
				DiscoveredAt: discoveredAt,
				FirstLogAt:   discoveredAt,
				ContractName: ds.Treasury,
				Client:       client,
			},
			LastSync: discoveredAt - 1,
		},
		Repo: repo,
	}
	return NewTreasuryFromAdapter(
		syncAdapter,
	)
}

func NewTreasuryFromAdapter(adapter *ds.SyncAdapter) *Treasury {
	obj := &Treasury{
		SyncAdapter: adapter,
	}
	obj.OnlyQuery = true
	obj.node = &core.Node{
		Client: adapter.Client,
	}
	obj.HexAddr = common.HexToAddress(obj.Address)
	return obj
}

func (mdl *Treasury) onLog(txLog types.Log, pools map[common.Address]bool) {
	zeroAddr := common.Address{}

	switch txLog.Topics[0] {
	case core.Topic("Transfer(address,address,uint256)"):
		from := common.BytesToAddress(txLog.Topics[1][32-20:])
		to := common.BytesToAddress(txLog.Topics[2][32-20:])
		if !((from == zeroAddr && to == mdl.HexAddr) || // repay profit mint
			(from == mdl.HexAddr && to == zeroAddr) || //repay loss  burn
			(pools[from] && to == mdl.HexAddr)) { // remove liquidity
			return
		}
		value, ok := new(big.Int).SetString(common.BytesToHash(txLog.Data).Hex()[2:], 16)
		if !ok {
			log.Fatal("Failed parsing value")
		}
		switch mdl.HexAddr {
		case from:
			mdl.Repo.AddTreasuryTransfer(int64(txLog.BlockNumber), txLog.Index,
				txLog.Address.Hex(), new(big.Int).Neg(value))
		case to:
			mdl.Repo.AddTreasuryTransfer(int64(txLog.BlockNumber), txLog.Index,
				txLog.Address.Hex(), value)
		}
	}
}

func (mdl *Treasury) getAddrs() (tokens []common.Address, pools map[common.Address]bool) {
	dieselTokensToPoolAndUToken := mdl.Repo.GetDieselTokens()
	tokens = make([]common.Address, 0, len(dieselTokensToPoolAndUToken)*2)
	pools = map[common.Address]bool{}
	for dieselToken, details := range dieselTokensToPoolAndUToken {
		tokens = append(tokens, common.HexToAddress(dieselToken), common.HexToAddress(details.UToken))
		pools[common.HexToAddress(details.Pool)] = true
	}
	return
}

func (mdl *Treasury) Query(queryTill int64) {
	queryFrom := mdl.GetLastSync() + 1
	tokenAddrs, pools := mdl.getAddrs()
	// from topics
	fromTopics := []common.Hash{
		common.HexToHash(mdl.Address),
		{}, // zero addr for mint on repay profit
	}
	for pool := range pools {
		fromTopics = append(fromTopics, common.BytesToHash(pool[:]))
	}
	// to topics
	toTopics := []common.Hash{
		{}, //  zero addr for burn on repay loss or when the treasury withdraw the underlying asset from pool
		common.HexToHash(mdl.Address),
	}
	logs, err := mdl.node.GetLogs(queryFrom, queryTill, tokenAddrs, [][]common.Hash{
		{core.Topic("Transfer(address,address,uint256)")},
		fromTopics,
		toTopics,
	})
	log.CheckFatal(err)
	for _, log := range logs {
		mdl.onLog(log, pools)
	}
}

func (mdl *Treasury) AfterSyncHook(syncTill int64) {
	// for treasury/dao
	mdl.Repo.CalCurrentTreasuryValue(syncTill)
	mdl.SyncAdapter.AfterSyncHook(syncTill)
}
func (mdl *Treasury) OnLog(txLog types.Log) {}
