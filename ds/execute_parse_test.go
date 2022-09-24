package ds

import (
	"encoding/hex"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

func getCallsAndEvents(callSigs []string, eventNames []string) (mainAction MainactionWithMulticall, events []*schemas.AccountOperation) {
	for _, name := range eventNames {
		events = append(events, &schemas.AccountOperation{Action: name})
	}
	var multicalls []multicall.Multicall2Call
	for _, sig := range callSigs {
		bytes, err := hex.DecodeString(sig)
		log.CheckFatal(err)
		multicalls = append(multicalls, multicall.Multicall2Call{
			CallData: bytes,
		})
	}
	mainAction = MainactionWithMulticall{MultiCalls: multicalls}
	return
}
func TestCmpLenSimple(t *testing.T) {
	log.SetTestLogging(t)
	mainAction, events := getCallsAndEvents(
		[]string{"59781034"},
		[]string{"AddCollateral(address,address,uint256)"},
	)
	if !mainAction.SameLenAsEvents(events) {
		log.Fatal()
	}
}
func TestCmpLenWith2Events(t *testing.T) {
	log.SetTestLogging(t)
	mainAction, events := getCallsAndEvents(
		[]string{"59781034", "d0e30db0"},
		[]string{"AddCollateral(address,address,uint256)", "ExecuteOrder"},
	)
	if !mainAction.SameLenAsEvents(events) {
		log.Fatal()
	}
}
func TestCmpLenWith2PlusRevertEvents(t *testing.T) {
	log.SetTestLogging(t)
	mainAction, events := getCallsAndEvents(
		[]string{"59781034", "81314b59", "c7fbf4de", "bdbeaa31"},
		[]string{"AddCollateral(address,address,uint256)", "ExecuteOrder", "ExecuteOrder"},
	)
	if !mainAction.SameLenAsEvents(events) {
		log.Fatal()
	}
}

func TestCmpLenTrivalRevert(t *testing.T) {
	log.SetTestLogging(t)
	mainAction, events := getCallsAndEvents(
		[]string{"81314b59"},
		nil,
	)
	if !mainAction.SameLenAsEvents(events) {
		log.Fatal()
	}
}
