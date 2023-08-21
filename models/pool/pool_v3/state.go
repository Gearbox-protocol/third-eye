package pool_v3

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

func (mdl *Poolv3) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStatePresent = true
	state, ok := obj.(*schemas.PoolState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	mdl.State = state
	mdl.Repo.AddDieselToken(mdl.State.DieselToken, mdl.State.UnderlyingToken, mdl.Address)
}

func (mdl *Poolv3) GetUnderlyingState() interface{} {
	return mdl.State
}
