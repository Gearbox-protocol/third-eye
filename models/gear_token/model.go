package gear_token

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/eRC20"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
)

type GearToken struct {
	*core.SyncAdapter
	contractETH *eRC20.ERC20
	State       map[string]*core.GearBalance
}

func NewGearToken(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *GearToken {
	pool := NewGearTokenFromAdapter(
		core.NewSyncAdapter(addr, core.GearToken, discoveredAt, client, repo),
	)
	pool.SetUnderlyingState(map[string]*core.GearBalance{})
	return pool
}

func NewGearTokenFromAdapter(adapter *core.SyncAdapter) *GearToken {
	cmContract, err := eRC20.NewERC20(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &GearToken{
		SyncAdapter: adapter,
		contractETH: cmContract,
	}
	return obj
}
