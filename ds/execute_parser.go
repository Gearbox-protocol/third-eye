package ds

import (
	"encoding/hex"
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type ExecuteParams struct {
	SessionId     string
	Protocol      common.Address
	CreditAccount common.Address
	Borrower      common.Address
	Index         uint
	BlockNumber   int64
}
type GBv2Multicall []struct {
	Target   common.Address `json:"target"`
	CallData []uint8        `json:"callData"`
}
type FacadeCallNameWithMulticall struct {
	Name       string                     `json:"name"`
	multiCalls []multicall.Multicall2Call `json:"-"`
	TestLen    int                        `json:"len"`
}

func (a FacadeCallNameWithMulticall) GetMulticalls() []multicall.Multicall2Call {
	return a.multiCalls
}

func NewFacadeCallNameWithMulticall(name string, multicalls []multicall.Multicall2Call) *FacadeCallNameWithMulticall {
	return &FacadeCallNameWithMulticall{
		Name:       name,
		multiCalls: multicalls,
	}
}

func (obj FacadeCallNameWithMulticall) String() string {
	str := ""
	for _, entry := range obj.multiCalls {
		funcSig := hex.EncodeToString(entry.CallData[:4])
		str += fmt.Sprintf("%s@%s ", entry.Target, funcSig)
	}
	return str
}

func (f FacadeCallNameWithMulticall) LenOfMulticalls() int {
	if f.TestLen != 0 {
		return f.TestLen
	}
	return len(f.multiCalls)
}

// handles revertIflessthan case where event is not emitted.
// also handles cases where number of execute order events emitted is less than execute calls
func (f *FacadeCallNameWithMulticall) SameMulticallLenAsEvents(version core.VersionType, events []*schemas.AccountOperation) bool {
	if f.TestLen != 0 {
		return f.TestLen == len(events)
	}
	if version.Eq(2) {
		return f.v2(events)
	} else if version.Eq(300) {
		return f.v3(events)
	}

	return false
}

func (f *FacadeCallNameWithMulticall) v3(events []*schemas.AccountOperation) bool {
	eventInd := 0
	callInd := 0
	callLen := len(f.multiCalls)
	eventLen := len(events)
	for callInd < callLen && (eventLen == 0 || eventInd < eventLen) {
		multiCall := f.multiCalls[callInd]
		sig := hex.EncodeToString(multiCall.CallData[:4])
		switch sig {
		case "59781034": // add collateral
			if events[eventInd].Action != "AddCollateral(address,address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "2b7c7b11": // increase debt
			if events[eventInd].Action != "IncreaseBorrowedAmount(address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "2a7ba1f7": // decrease debt
			if events[eventInd].Action != "DecreaseBorrowedAmount(address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "c690908a": // enable token
			if events[eventInd].Action != "TokenEnabled(address,address)" {
				return false
			}
			// eventInd++
			callInd++
		case "23e27a64": // disable token
			if events[eventInd].Action != "TokenDisabled(address,address)" {
				return false
			}
			// eventInd++
			callInd++
		case "81314b59": // revert if less than // ignore for event
			callInd++
		case "6631d161": // payBot
			callInd++
		case "0768bbfe": // setFullCheckParams
			callInd++
		case "565a820d": // revokeAdapterAllowances
			callInd++
		case "384f69fa": // scheduleWithdrawal
			callInd++
		case "712c10ad": // updateQuota
			callInd++
		case "a6181cb0": // onDemandPriceUpdate
			callInd++
		default: //execute order
			// it might happen that some of the execution call are not executed so len of provided multicalls will be more than executed calls.
			executeEvent := 0
			for eventInd < len(events) && events[eventInd].Action == "ExecuteOrder" {
				executeEvent++
				eventInd++
			}
			executeCall := 0
			for callInd < callLen && !utils.Contains([]string{"59781034", "2b7c7b11", "2a7ba1f7", "c690908a", "23e27a64",
				"81314b59", "6631d161", "0768bbfe", "565a820d", "384f69fa", "712c10ad", "a6181cb0"},
				hex.EncodeToString(f.multiCalls[callInd].CallData[:4])) {
				executeCall++
				callInd++
			}
			if executeEvent > executeCall { // if execute events more than calls
				return false
			}
		}
	}
	return callInd == callLen && eventInd == eventLen
}

// handles revertIflessthan case where event is not emitted. L1
// handles failed tokenDisabled call. L2
// also handles cases where number of execute order events emitted is less than execute calls // L3
// event len can be zero in case of all failed calls, so no events.
func (f *FacadeCallNameWithMulticall) v2(events []*schemas.AccountOperation) bool {
	eventInd := 0
	callInd := 0
	callLen := len(f.multiCalls)
	eventLen := len(events)
	for callInd < callLen && (eventLen == 0 || eventInd < eventLen) {
		multiCall := f.multiCalls[callInd]
		sig := hex.EncodeToString(multiCall.CallData[:4])
		switch sig {
		case "59781034": // add collateral,
			if events[eventInd].Action != "AddCollateral(address,address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "6d75b9ee": // add collateral extended 2.2
			if events[eventInd].Action != "AddCollateral(address,address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "2b7c7b11": // increase debt
			if events[eventInd].Action != "IncreaseBorrowedAmount(address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "2a7ba1f7": // decrease debt
			if events[eventInd].Action != "DecreaseBorrowedAmount(address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "c690908a": // enable token
			if events[eventInd].Action != "TokenEnabled(address,address)" {
				return false
			}
			eventInd++
			callInd++
		case "23e27a64": // disable token
			if events[eventInd].Action == "TokenDisabled(address,address)" { // L2
				eventInd++
			}
			callInd++ // disabled call can fail.
		case "81314b59": // revert if less than // ignore for event // revertIfReceivedLessThan // L1
			callInd++
		default: //execute order // L3
			// it might happen that some of the execution call are not executed so len of provided multicalls will be more than executed calls.
			// takes longest array of execute order events and calls, and compares their sizes. len events< len calls
			executeEvent := 0
			for eventInd < len(events) && events[eventInd].Action == "ExecuteOrder" {
				executeEvent++
				eventInd++
			}
			executeCall := 0
			for callInd < callLen && !utils.Contains([]string{
				"59781034", // add collateral
				"6d75b9ee", // add collateral 2.2
				"2b7c7b11", // increase debt
				"2a7ba1f7", // decrease debt
				"c690908a", // enable token
				"23e27a64", // disable token
				"81314b59", // revertIfReceivedLessThan
			},
				hex.EncodeToString(f.multiCalls[callInd].CallData[:4])) {
				executeCall++
				callInd++
			}
			if executeEvent > executeCall { // if execute events more than calls
				return false
			}
		}
	}
	return callInd == callLen && eventInd == eventLen

}

type BorrowerAndTo struct {
	Borrower common.Address
	To       common.Address
}
type ExecuteParserI interface {
	GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ExecuteParams) []*KnownCall
	// ignores revertIfLessThan
	GetMainCalls(txHash, creditFacade string) []*FacadeCallNameWithMulticall
	GetTransfersAtClosev2(txHash string, account, underlyingToken string, users BorrowerAndTo) core.Transfers
}

type KnownCall struct {
	// Input string
	Name      string         `json:"name"`
	Args      *core.Json     `json:"args"`
	Transfers core.Transfers `json:"transfers"`
}
