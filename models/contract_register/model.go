package contract_register
import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"

)
type ContractRegister struct {
	*core.SyncAdapter
	*core.State
}


func NewContractRegister(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *ContractRegister {
	obj := &ContractRegister{
		SyncAdapter: &core.SyncAdapter{
			Type: "ContractRegister",
			Address: addr,
			Client: client,
		},
		State: &core.State{Repo: repo},
	}
	firstDetection := obj.DiscoverFirstLog()
	obj.SyncAdapter.DiscoveredAt = discoveredAt
	obj.SyncAdapter.FirstLogAt = firstDetection
	obj.SyncAdapter.LastSync = firstDetection
	return obj
}
