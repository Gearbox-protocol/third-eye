package pool

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v2"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
)

func NewPoolFromAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	switch adapter.GetVersion() {
	case core.NewVersion(1), core.NewVersion(2):
		return pool_v2.NewPoolFromAdapter(adapter)
	case core.NewVersion(300):
		return pool_v3.NewPoolFromAdapter(adapter)
	}
	return nil
}
