package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"fmt"
)

func (mdl *CreditManager) SetState(obj interface{}) {
	state, ok := obj.(*core.CreditManager)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	mdl.State = state
}


func (mdl *CreditManager) AddCreditOwnerSession(owner, sessionId string) {
	mdl.State.Sessions.Set(owner, sessionId)
}

func (mdl *CreditManager) RemoveCreditOwnerSession(owner string) {
	mdl.State.Sessions.Remove(owner)
}

func (mdl *CreditManager) GetCreditOwnerSession(owner string) string {
	sessionId := mdl.State.Sessions.Get(owner)
	if sessionId == "" {
		panic(
			fmt.Sprintf("session id not found for %s in %+v\n", owner, mdl.State.Sessions),
		)
	}
	return sessionId
}

func (mdl *CreditManager) GetUnderlyingToken() string {
	return mdl.State.UnderlyingToken
}
