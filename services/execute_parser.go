package services

import (
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/services/trace_service"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

type ExecuteParser struct {
	IgnoreCMEventIds map[common.Hash]bool
	trace_service.InternalFetcher
}

func getCMEventIds() map[common.Hash]bool {
	ids := map[common.Hash]bool{}
	if abiObj, err := abi.JSON(strings.NewReader(creditFacade.CreditFacadeABI)); err == nil {
		for _, event := range abiObj.Events {
			ids[event.ID] = true
		}
	}
	if abiObj, err := abi.JSON(strings.NewReader(creditManager.CreditManagerABI)); err == nil {
		for _, event := range abiObj.Events {
			ids[event.ID] = true
		}
	}
	return ids
}
func NewExecuteParser(cfg *config.Config, client core.ClientI) ds.ExecuteParserI {
	return &ExecuteParser{
		IgnoreCMEventIds: getCMEventIds(),
		InternalFetcher:  trace_service.NewInternalFetcher(cfg, client),
	}
}

func (ep *ExecuteParser) GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ds.ExecuteParams) []*ds.KnownCall {
	if len(paramsList) == 0 {
		return nil
	}
	trace := ep.GetTxTrace(txHash, true)
	filter := ExecuteFilter{paramsList: paramsList, creditManager: common.HexToAddress(creditManagerAddr)}
	calls := filter.getExecuteCalls(trace.CallTrace)

	executeTransfers := filter.getExecuteTransfers(trace.Logs, ep.IgnoreCMEventIds)

	// check if parsed execute Order currently
	if len(calls) == len(executeTransfers) {
		for i, call := range calls {
			call.Transfers = executeTransfers[i]
		}
	} else {
		fileName := fmt.Sprintf("trace-%s-%s.json", txHash, time.Now())
		os.WriteFile(fileName, []byte(utils.ToJson(trace)), os.ModePerm)
		log.Fatalf("Calls %d execute details %d tx:%s creditManager:%s",
			len(calls), len(executeTransfers), txHash, creditManagerAddr)
	}
	return calls
}

//////////////////////////
// parser functions for v2
//////////////////////////
// GetMainCalls
func (ep *ExecuteParser) GetMainCalls(txHash, creditFacade string) []*ds.FacadeCallNameWithMulticall {
	trace := ep.GetTxTrace(txHash, false)
	data, err := ep.getMainEvents(trace.CallTrace, common.HexToAddress(creditFacade))
	if err != nil {
		log.Fatal(err.Error(), "for txHash", txHash)
	}
	return data
}

func (ep *ExecuteParser) getMainEvents(call *trace_service.Call, creditFacade common.Address) ([]*ds.FacadeCallNameWithMulticall, error) {
	mainEvents := []*ds.FacadeCallNameWithMulticall{}
	if utils.Contains([]string{"CALL", "DELEGATECALL", "JUMP"}, call.CallerOp) {
		if creditFacade == common.HexToAddress(call.To) && len(call.Input) >= 10 {
			switch call.Input[2:10] {
			case "caa5c23f", // multicall
				"5f73fbec", // closeCreditAccount
				"82871ace", // liquidateExpiredCreditAccount
				"5d91a0e0", // liquidateCreditAccount
				"7071b7c5": // openCreditAccountMulticall
				event, err := getCreditFacadeMainEvent(call.Input)
				if err != nil {
					return nil, err
				}
				mainEvents = append(mainEvents, event)
			}
		} else {
			for _, c := range call.Calls {
				data, err := ep.getMainEvents(c, creditFacade)
				if err != nil {
					return nil, err
				}
				mainEvents = append(mainEvents, data...)
			}
		}
	}
	return mainEvents, nil
}

var creditFacadeParser *abi.ABI

func init() {
	creditFacadeParser = core.GetAbi("CreditFacade")
}
func getCreditFacadeMainEvent(input string) (*ds.FacadeCallNameWithMulticall, error) {
	hexData, err := hex.DecodeString(input[2:])
	if err != nil {
		return nil, err
	}
	method, err := creditFacadeParser.MethodById(hexData[:4])
	if err != nil {
		return nil, err
	}
	// unpack in the map
	data := map[string]interface{}{}
	err = method.Inputs.UnpackIntoMap(data, hexData[4:])
	if err != nil {
		log.Fatal(err)
	}
	calls, ok := data["calls"].([]struct {
		Target   common.Address `json:"target"`
		CallData []uint8        `json:"callData"`
	})
	if !ok {
		log.Fatal("calls type is different the creditFacade multicall: ", reflect.TypeOf(data["calls"]))
	}
	multicalls := []multicall.Multicall2Call{}
	for _, call := range calls {
		multicalls = append(multicalls, multicall.Multicall2Call{
			Target:   call.Target,
			CallData: call.CallData,
		})
	}
	return ds.NewFacadeCallNameWithMulticall(
		ds.FacadeAccountMethodSigToCallName(method.Name),
		multicalls,
	), nil
}

// GetTransfers
// currently only valid for closeCreditAccount v2
func (ep *ExecuteParser) GetTransfers(txHash, account, underlyingToken string, users ds.BorrowerAndTo) core.Transfers {
	trace := ep.GetTxTrace(txHash, true)
	return getCloseAccountv2Transfers(trace, account, underlyingToken, users)
}

// currently only valid for closeCreditAccount v2
func getCloseAccountv2Transfers(trace *trace_service.TenderlyTrace, account, underlyingToken string, users ds.BorrowerAndTo) core.Transfers {
	transfers := getTransfersToUser(trace.Logs, account, underlyingToken, users)
	// convertWETH is set, only valid for closecreditaccountv2
	convertWETHInd := 2 + 8 + 64 + 64 + 64
	// for close call if convertEThInd is true
	if trace.CallTrace.Input[:10] == "0x5f73fbec" && trace.CallTrace.Input[convertWETHInd-1] == '1' {
		ethAmount := ethTransferDueToConvertWETH(trace.CallTrace, users)
		if ethAmount == nil {
			// log.Msgf("Can't get unwrapped WETH amount at closeCreditAccount(%s) sent to user. Tx: %s.", account, users.Borrower, trace.TxHash)
			ethAmount = new(big.Int)
		}
		if transfers[underlyingToken] == nil {
			transfers[underlyingToken] = new(big.Int)
		}
		transfers[underlyingToken] = new(big.Int).Add(transfers[underlyingToken], ethAmount)
	}
	return transfers
}

// eth transfer due to convertWETH
func ethTransferDueToConvertWETH(call *trace_service.Call, users ds.BorrowerAndTo) (ethAmount *big.Int) {
	if len(call.Input) == 10+64*2 && call.Input[:10] == "0x5869dba8" && common.HexToAddress(call.Input[10:74]) == users.To {
		ethAmount, _ := new(big.Int).SetString(call.Input[74:], 16)
		return ethAmount
	}
	for _, innerCall := range call.Calls {
		if ethAmount := ethTransferDueToConvertWETH(innerCall, users); ethAmount != nil {
			return ethAmount
		}
	}
	return nil
}

// is valid for closeCreditAccount v2
// tenderly has logs for events(we mainly use for Transfer on token) and calls( for unwrapETH on wethgateway)
// wrapWETH is also present in closecreditaccount, but it sends the wrapped eth back to user and then the user has approval on weth for creditmanager so in second step the weth is transferred
// handling native eth refund is only needed when convertETH is true
// native eth transfer from account is handled in parent function, not in this function
func getTransfersToUser(txLogs []trace_service.Log, account, underlyingToken string, users ds.BorrowerAndTo) core.Transfers {
	transfers := core.Transfers{}
	for _, raw := range txLogs {
		eventLog := raw.Raw
		if eventLog.Topics[0] == core.Topic("Transfer(address,address,uint256)") { // transfer event
			to := common.BytesToAddress(eventLog.Topics[2][:])
			from := common.BytesToAddress(eventLog.Topics[1][:])
			token := eventLog.Address.Hex()
			var sign *big.Int
			if from == users.Borrower && to.Hex() == account && token == underlyingToken {
				sign = big.NewInt(-1)
			} else {
				if !(to == users.Borrower || to == users.To) {
					continue
				}
				sign = big.NewInt(1)
			}
			amt, b := new(big.Int).SetString(eventLog.Data[2:], 16)
			if !b {
				log.Fatal("failed at serializing transfer data in int")
			}
			amt = new(big.Int).Mul(sign, amt)
			oldBalance := new(big.Int)
			if transfers[token] != nil {
				oldBalance = transfers[token]
			}
			transfers[token] = new(big.Int).Add(amt, oldBalance)
		}
	}
	return transfers
}
