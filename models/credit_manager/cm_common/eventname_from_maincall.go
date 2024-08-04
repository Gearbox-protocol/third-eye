package cm_common

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func (mdl CommonCMAdapter) setLiquidateStatus(sessionId string, isExpired bool) {
	status := schemas.Liquidated
	if mdl.State.Paused {
		status = schemas.LiquidatePaused
	} else if isExpired {
		status = schemas.LiquidateExpired
	}
	mdl.UpdateClosedSessionStatus(sessionId, status)
}

func (mdl CommonCMAdapter) getEventNameFromCallv2(mainCallName string, sessionId string) string {
	var mainEventFromCall string
	switch mainCallName {
	case ds.FacadeMulticallCall:
		mdl.SetSessionIsUpdated(sessionId)
		mainEventFromCall = "MultiCallStarted(address)"
	case ds.FacadeOpenMulticallCall:
		mdl.SetSessionIsUpdated(sessionId)
		mainEventFromCall = "OpenCreditAccount(address,address,uint256,uint16)"
	case ds.FacadeLiquidateCall, ds.FacadeLiquidateExpiredCall:
		mdl.setLiquidateStatus(sessionId, mainCallName == ds.FacadeLiquidateExpiredCall) // SET_LIQ_STATUS_AFTER_CALL
		mainEventFromCall = "LiquidateCreditAccount(address,address,address,uint256)"
	case ds.FacadeCloseAccountCall:
		mainEventFromCall = "CloseCreditAccount(address,address)"
	}
	return mainEventFromCall
}

func (mdl CommonCMAdapter) getEventNameFromCallv3(mainCallName string, sessionId string) string {
	var mainEventFromCall string
	switch mainCallName {
	case ds.FacadeMulticallCall, ds.FacadeBotMulticallCall:
		mdl.SetSessionIsUpdated(sessionId)
		mainEventFromCall = "StartMultiCall(address,address)"
	case ds.FacadeOpenMulticallCall:
		mdl.SetSessionIsUpdated(sessionId)
		mainEventFromCall = "OpenCreditAccount(address,address,address,uint256)"
	case ds.FacadeLiquidateCall:
		mainEventFromCall = "LiquidateCreditAccount(address,address,address,uint256)"
	case ds.FacadeCloseAccountCall:
		mainEventFromCall = "CloseCreditAccount(address,address)"
	}
	return mainEventFromCall
}

func (mdl CommonCMAdapter) getEventNameFromCall(version core.VersionType, mainCallName string, sessionId string) string {
	switch version {
	case core.NewVersion(2):
		return mdl.getEventNameFromCallv2(mainCallName, sessionId)
	case core.NewVersion(300):
		return mdl.getEventNameFromCallv3(mainCallName, sessionId)
	default:
		log.Fatalf("version(%d) is not called for getEventNameFromCall", version)
		return ""
	}
}
