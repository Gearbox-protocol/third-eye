package credit_manager

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v2"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v3"
)

func NewCMFromAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	switch adapter.GetVersion() {
	case core.NewVersion(1):
		return cm_v1.NewCMv1FromAdapter(adapter)
	case core.NewVersion(2):
		return cm_v2.NewCMv2FromAdapter(adapter)
	case core.NewVersion(300):
		return cm_v3.NewCMv3FromAdapter(adapter)
	}
	panic("")
}

type BlockChangeI interface {
	OnBlockChange(int64) ([]multicall.Multicall2Call, []func(multicall.Multicall2Result))
	UpdateSessionWithDirectTokenTransferBefore(int64)
	IsAddrChanged() bool
}

func GetCMForWrapper(adapter ds.SyncAdapterI) (cm BlockChangeI) {
	switch v := adapter.(type) {
	case *cm_v1.CMv1:
		cm = v
	case *cm_v2.CMv2:
		cm = v
	case *cm_v3.CMv3:
		cm = v
	}
	return
}
