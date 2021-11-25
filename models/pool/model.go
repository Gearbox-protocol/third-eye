package pool

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/artifacts/poolService"
)

type Pool struct {
	*core.SyncAdapter
	*core.State
	ContractETH *poolService.PoolService
}

func NewPool(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *Pool {
	return NewPoolFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "Pool", discoveredAt, client),
	)
}

func NewPoolFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *Pool {
	cmContract, err:=poolService.NewPoolService(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &Pool{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
		ContractETH: cmContract,
	}
	return obj
}

func (mdl *Pool) OnLog(txLog types.Log) {
}
