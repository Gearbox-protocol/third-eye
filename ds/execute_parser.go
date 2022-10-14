package ds

import (
	"encoding/hex"
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/utils"
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
type MainactionWithMulticall struct {
	Name       string                     `json:"name"`
	MultiCalls []multicall.Multicall2Call `json:"-"`
	TestLen    int                        `json:"len"`
}

func (obj MainactionWithMulticall) String() string {
	str := ""
	for _, entry := range obj.MultiCalls {
		funcSig := hex.EncodeToString(entry.CallData[:4])
		str += fmt.Sprintf("%s@%s ", entry.Target, funcSig)
	}
	return str
}

func (f MainactionWithMulticall) Len() int {
	return len(f.MultiCalls)
}

func (f *MainactionWithMulticall) SameLenAsEvents(events []*schemas.AccountOperation) bool {
	// for _, event := range events {
	// 	log.Info(event.Action)
	// }
	// log.Info("###")
	// for _, call := range f.MultiCalls {
	// 	log.Info(hex.EncodeToString(call.CallData[:4]))
	// }
	if f.TestLen != 0 {
		return f.TestLen == len(events)
	}
	eventInd := 0
	callInd := 0
	callLen := len(f.MultiCalls)
	eventLen := len(events)
	for callInd < callLen && (eventLen == 0 || eventInd < eventLen) {
		multiCall := f.MultiCalls[callInd]
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
			eventInd++
			callInd++
		case "23e27a64": // disable token
			if events[eventInd].Action != "TokenDisabled(address,address)" {
				return false
			}
			eventInd++
			callInd++
		case "81314b59": // revert if less than // ignore for event
			callInd++
		default: //execute order
			// it might happen that some of the execution call are not executed so len of provided multicalls will be more than executed calls.
			executeEvent := 0
			for eventInd < len(events) && events[eventInd].Action == "ExecuteOrder" {
				executeEvent++
				eventInd++
			}
			executeCall := 0
			for callInd < callLen && !utils.Contains([]string{"59781034", "2b7c7b11", "2a7ba1f7", "c690908a", "23e27a64", "81314b59"},
				hex.EncodeToString(f.MultiCalls[callInd].CallData[:4])) {
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
	GetMainEventLogs(txHash, creditFacade string) []*MainactionWithMulticall
	GetTransfers(txHash string, account, underlyingToken string, users BorrowerAndTo) core.Transfers
}

type KnownCall struct {
	// Input string
	Depth     uint8          `json:"depth"`
	Name      string         `json:"name"`
	Args      *core.Json     `json:"args"`
	Transfers core.Transfers `json:"transfers"`
}
