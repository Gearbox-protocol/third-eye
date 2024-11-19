package gear_token

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/eRC20"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type GearToken struct {
	*ds.SyncAdapter
	contractETH *eRC20.ERC20
	State       map[string]*schemas.GearBalance
}

func NewGearToken(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *GearToken {
	gearTokenAdapter := NewGearTokenFromAdapter(
		ds.NewSyncAdapter(addr, ds.GearToken, discoveredAt, client, repo),
	)
	gearTokenAdapter.SetUnderlyingState([]*schemas.GearBalance{})
	return gearTokenAdapter
}

func NewGearTokenFromAdapter(adapter *ds.SyncAdapter) *GearToken {
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
