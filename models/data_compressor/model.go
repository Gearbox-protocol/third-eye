package data_compressor
import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)
type DataCompressor struct {
	*core.SyncAdapter
	*core.State
}


func NewDataCompressor(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *DataCompressor {
	obj := &DataCompressor{
		SyncAdapter: &core.SyncAdapter{
			Type: "DataCompressor",
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

func (mdl *DataCompressor) OnLog(txLog types.Log){
}