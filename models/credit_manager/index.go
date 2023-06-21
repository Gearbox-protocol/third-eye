package credit_manager

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v2"
)

func NewCMFromAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	switch adapter.GetVersion() {
	case 1:
		return cm_v1.NewCMv1FromAdapter(adapter)
	case 2:
		return cm_v2.NewCMv2FromAdapter(adapter)
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
	}
	return
}
