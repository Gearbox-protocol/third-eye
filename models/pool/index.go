package pool

import (
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v2"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
)

func NewPoolFromAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	switch adapter.GetVersion() {
	case 1, 2:
		return pool_v2.NewPoolFromAdapter(adapter)
	case 3:
		return pool_v3.NewPoolFromAdapter(adapter)
	}
	return nil
}
