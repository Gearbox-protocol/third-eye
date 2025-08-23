package ds

import (
	"encoding/hex"
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
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
	facade     string                     `json:"-"`
	multiCalls []multicall.Multicall2Call `json:"-"`
	TestLen    int                        `json:"len"`
}

func (a FacadeCallNameWithMulticall) GetMulticalls() []multicall.Multicall2Call {
	return a.multiCalls
}

func NewFacadeCallNameWithMulticall(facade, name string, multicalls []multicall.Multicall2Call) *FacadeCallNameWithMulticall {
	return &FacadeCallNameWithMulticall{
		Name:       name,
		multiCalls: multicalls,
		facade:     common.HexToAddress(facade).Hex(),
	}
}

func (obj FacadeCallNameWithMulticall) String() string {
	str := ""
	for _, entry := range obj.multiCalls {
		funcSig := hex.EncodeToString(entry.CallData)
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
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err, f, utils.ToJson(events))
		}
	}()

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
	if len(f.facade) != 42 {
		log.Fatal("facade address is missing from executeParser DS")
	}
	eventInd := 0
	callInd := 0
	callLen := len(f.multiCalls)
	eventLen := len(events)
	for callInd < callLen || eventInd < eventLen {
		multiCall := f.multiCalls[callInd]
		sig := hex.EncodeToString(multiCall.CallData[:4])
		// log.Info(callInd, eventInd)
		switch sig {
		case "59781034", // add collateral
			"6d75b9ee", // add collateral extended 2.2 // v310 too
			"438f8fcc": // AddCollateralWithPermit // v310 too
			if events[eventInd].Action != "AddCollateral(address,address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "2b7c7b11": // increase debt
			if events[eventInd].Action != "IncreaseDebt(address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "2a7ba1f7": // decrease debt
			if events[eventInd].Action != "DecreaseDebt(address,uint256)" {
				return false
			}
			eventInd++
			callInd++
		case "c690908a": // enable token
			// if events[eventInd].Action != "TokenEnabled(address,address)" {
			// 	return false
			// }
			// eventInd++
			callInd++
		case "23e27a64": // disable token
			// if events[eventInd].Action != "TokenDisabled(address,address)" {
			// 	return false
			// }
			// eventInd++
			callInd++
		case "81314b59": // revert if less than  // can't find in v3
			callInd++
		case "82ff942c",
			"2f2ca49b": // v300 // second is storeExpectedBalances((address,int256)[])
			// "6161dc85": // exactDiffInput((bytes,uint256,uint256,uint256))
			callInd++
		case "0768bbfe": // setFullCheckParams
			callInd++
		case "565a820d": // revokeAdapterAllowances
			callInd++
		case "712c10ad": // updateQuota
			if eventInd < eventLen && events[eventInd].Action == "UpdateQuota" {
				eventInd++
			}
			callInd++
		case "6c68e109": // onDemandPriceUpdate
			callInd++
		case "28b83c48": // onDemandPriceUpdates
			callInd++
		case "f42aeb00": // compareBalances // v310
			callInd++
		case "1f1088a0": // withdrawcollateral
			if eventInd < eventLen {
				if events[eventInd].Action != "WithdrawCollateral(address,address,uint256,address)" {
					return false
				}
				eventToken := (*events[eventInd].Args)["token"]
				if eventToken != nil &&
					common.HexToAddress(eventToken.(string)) == common.BytesToAddress(multiCall.CallData[4:4+32]) {
					eventInd++
				}
			}
			callInd++
		default: //execute
			// it might happen that some of the execution call are not executed so len of provided multicalls will be more than executed calls.
			executeEvent := 0
			for eventInd < len(events) && events[eventInd].Action == "ExecuteOrder" {
				executeEvent++
				eventInd++
			}
			executeCall := 0
			for callInd < callLen && f.multiCalls[callInd].Target.Hex() != f.facade { // until multicall call that is not for facade is seen
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
		case "59781034", // add collateral,
			"6d75b9ee": // add collateral extended 2.2
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
	GetExecuteCalls(version core.VersionType, txHash, creditManagerAddr string, paramsList []ExecuteParams) []*KnownCall
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
