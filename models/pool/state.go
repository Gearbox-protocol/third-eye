package pool

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

func (mdl *Pool) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStateToSave = true
	state, ok := obj.(*schemas.PoolState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	mdl.State = state
	mdl.Repo.AddDieselToken(mdl.State.DieselToken, mdl.State.UnderlyingToken, mdl.Address)
}

func (mdl *Pool) GetUnderlyingState() interface{} {
	return mdl.State
}
