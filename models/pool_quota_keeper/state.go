package pool_quota_keeper

import (
	"log"

	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
)

// - for quotaDetails load the latest entry for that poolKeeper
func (mdl *PoolQuotaKeeper) SetUnderlyingState(obj interface{}) {
	// not set as there is no underlying state to return
	// mdl.UnderlyingStatePresent = true
	switch state := obj.(type) {
	case *schemas_v3.QuotaDetails:
		mdl.quotas[state.Token] = state
	default:
		log.Fatal("Type assertion for credit manager state failed")
	}
}
