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

func getCallsAndEvents(facade string, callSigs []string, eventNames []string, withdrawalTokens ...string) (mainAction FacadeCallNameWithMulticall, events []*schemas.AccountOperation) {
	withdrawalTokensInd := 0
	for _, name := range eventNames {
		event := &schemas.AccountOperation{Action: name}
		if name == "WithdrawCollateral(address,address,uint256,address)" {
			event.Args = &core.Json{"token": common.HexToAddress(withdrawalTokens[withdrawalTokensInd]).Hex()}
			withdrawalTokensInd += 1
		}
		events = append(events, event)
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
	if !mainAction.SameMulticallLenAsEvents(nil, core.NewVersion(2), events) {
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
	if !mainAction.SameMulticallLenAsEvents(nil, core.NewVersion(2), events) {
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
	if !mainAction.SameMulticallLenAsEvents(nil, core.NewVersion(2), events) {
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
	if !mainAction.SameMulticallLenAsEvents(nil, core.NewVersion(2), events) {
		log.Fatal()
	}
}
func TestSmallMulticallwitheventsV3(t *testing.T) {
	log.SetTestLogging(t)
	//
	mainAction, events := getCallsAndEvents(
		"0x30efB02202352290D1900bc76De0Ab00A2A6DACb",
		[]string{"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@28b83c4800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@6d75b9ee000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000000000000000000000000000000006fb4ca69b",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@712c10ad00000000000000000000000087fa6c0296c986d1c901d72571282d57916b964affffffffffffffffffffffffffffffffffffffff8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@2a7ba1f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			"0x30d8ab4cd05792a2Eca88D592148Cd94E1186552@3d18b912",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@1f1088a000000000000000000000000087fa6c0296c986d1c901d72571282d57916b964affffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000723265936ec47139d543c63f081de65e587cf5c1",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@1f1088a0000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd52ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000723265936ec47139d543c63f081de65e587cf5c1",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@1f1088a00000000000000000000000004e3fbd56cd56c3e72c1403e103b45db9da5b9d2bffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000723265936ec47139d543c63f081de65e587cf5c1",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@1f1088a0000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000723265936ec47139d543c63f081de65e587cf5c1",
			"0x30efB02202352290D1900bc76De0Ab00A2A6DACb@1f1088a00000000000000000000000008292bb45bf1ee4d140127049757c2e0ff06317edffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000723265936ec47139d543c63f081de65e587cf5c1"},
		[]string{
			"AddCollateral(address,address,uint256)",
			"UpdateQuota",
			"DecreaseDebt(address,uint256)",
			"ExecuteOrder",
			"ExecuteOrder",
			"WithdrawCollateral(address,address,uint256,address)",
			"WithdrawCollateral(address,address,uint256,address)",
			"WithdrawCollateral(address,address,uint256,address)",
			"WithdrawCollateral(address,address,uint256,address)",
			"WithdrawCollateral(address,address,uint256,address)",
		},
		// withdrawaltokens
		"0x87fa6c0296c986d1c901d72571282d57916b964a",
		"0xd533a949740bb3306d119cc777fa900ba034cd52",
		"0x4e3fbd56cd56c3e72c1403e103b45db9da5b9d2b",
		"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		"0x8292bb45bf1ee4d140127049757c2e0ff06317ed",
	)
	if !mainAction.SameMulticallLenAsEvents(nil, core.NewVersion(300), events) {
		log.Fatal()
	}
}
