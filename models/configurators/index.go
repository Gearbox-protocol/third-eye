package configurators

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/configurators/configurator_v2"
	"github.com/Gearbox-protocol/third-eye/models/configurators/configurator_v3"
	"github.com/Gearbox-protocol/third-eye/models/configurators/credit_filter"
)

func NewConfiguratorFromAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	switch adapter.GetVersion() {
	case core.NewVersion(1):
		return credit_filter.NewCreditFilterFromAdapter(adapter)
	case core.NewVersion(2):
		return configurator_v2.NewConfiguratorv2FromAdapter(adapter)
	case core.NewVersion(300):
		return configurator_v3.NewConfiguratorv3FromAdapter(adapter)
	}
	panic("")
}
