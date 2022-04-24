package data_compressor

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/core/types"
)

type DataCompressor struct {
	*ds.SyncAdapter
}

func NewDataCompressor(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *DataCompressor {
	obj := &DataCompressor{
		SyncAdapter: ds.NewSyncAdapter(addr, "AddressProvider", discoveredAt, client, repo),
	}
	return obj
}

func (mdl *DataCompressor) OnLog(txLog types.Log) {
}
