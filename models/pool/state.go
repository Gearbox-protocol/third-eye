package pool

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (mdl *Pool) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStatePresent = true
	state, ok := obj.(*core.PoolState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	mdl.State = state
}

func (mdl *Pool) GetUnderlyingState() interface{} {
	return mdl.State
}
