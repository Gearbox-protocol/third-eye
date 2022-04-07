package data_compressor

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

type DataCompressor struct {
	*core.SyncAdapter
}

func NewDataCompressor(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *DataCompressor {
	obj := &DataCompressor{
		SyncAdapter: core.NewSyncAdapter(addr, "AddressProvider", discoveredAt, client, repo),
	}
	return obj
}

func (mdl *DataCompressor) OnLog(txLog types.Log) {
}
