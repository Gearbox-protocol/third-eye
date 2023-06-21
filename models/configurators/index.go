package configurators

import (
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/configurators/configurator_v2"
	"github.com/Gearbox-protocol/third-eye/models/configurators/credit_filter"
)

func NewConfiguratorFromAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	switch adapter.GetVersion() {
	case 1:
		return credit_filter.NewCreditFilterFromAdapter(adapter)
	case 2:
		return configurator_v2.NewConfiguratorv2FromAdapter(adapter)
	}
	panic("")
}
