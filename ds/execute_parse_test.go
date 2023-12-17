package ds

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

func getCallsAndEvents(facade string, callSigs []string, eventNames []string) (mainAction FacadeCallNameWithMulticall, events []*schemas.AccountOperation) {
	for _, name := range eventNames {
		events = append(events, &schemas.AccountOperation{Action: name})
	}
	var multicalls []multicall.Multicall2Call
	for _, input := range callSigs {
		_s := strings.Split(input, "@")
		bytes, err := hex.DecodeString(_s[1])
		log.CheckFatal(err)
		multicalls = append(multicalls, multicall.Multicall2Call{
			CallData: bytes,
			Target:   common.HexToAddress(_s[0]),
		})
	}
	mainAction = FacadeCallNameWithMulticall{
		facade:     facade,
		multiCalls: multicalls,
	}
	return
}
func TestCmpLenSimpleV2(t *testing.T) {
	log.SetTestLogging(t)
	mainAction, events := getCallsAndEvents(
		"0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB",
		[]string{"0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB@59781034"},
		[]string{"AddCollateral(address,address,uint256)"},
	)
	if !mainAction.SameMulticallLenAsEvents(core.NewVersion(2), events) {
		log.Fatal()
	}
}
func TestCmpLenWith2EventsV2(t *testing.T) {
	log.SetTestLogging(t)
	mainAction, events := getCallsAndEvents(
		"0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB",
		[]string{"0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB@59781034", "0x020bafa614d63087C4B3C8244F1e2c8A3859Ce4E@d0e30db0"},
		[]string{"AddCollateral(address,address,uint256)", "ExecuteOrder"},
	)
	if !mainAction.SameMulticallLenAsEvents(core.NewVersion(2), events) {
		log.Fatal()
	}
}
func TestCmpLenWith2PlusRevertEventsV2(t *testing.T) {
	log.SetTestLogging(t)
	mainAction, events := getCallsAndEvents(
		"0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB",
		[]string{
			"0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB@59781034", "0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB@81314b59", // collateral and revertIfLessthan
			"0x020bafa614d63087C4B3C8244F1e2c8A3859Ce4E@c7fbf4de", "0x020bafa614d63087C4B3C8244F1e2c8A3859Ce4E@bdbeaa31",
		},
		[]string{"AddCollateral(address,address,uint256)", "ExecuteOrder", "ExecuteOrder"},
	)
	if !mainAction.SameMulticallLenAsEvents(core.NewVersion(2), events) {
		log.Fatal()
	}
}

func TestCmpLenTrivalRevertV2(t *testing.T) {
	log.SetTestLogging(t)
	//
	mainAction, events := getCallsAndEvents(
		"0xb82C5D0A6750aD1E5c2f74CFD7e2E4788f2b0aBB",
		[]string{"0x020bafa614d63087C4B3C8244F1e2c8A3859Ce4E@81314b59"},
		nil,
	)
	if !mainAction.SameMulticallLenAsEvents(core.NewVersion(2), events) {
		log.Fatal()
	}
}
