package services

import (
	"encoding/hex"
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/services/trace_service"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

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
				event, err := getCreditFacadeMainEvent(call.Input, creditFacadev2Parser)
				if err != nil {
					return nil, err
				}
				mainEvents = append(mainEvents, event)
			case "ebe4107c", // multicall(address,calls)
				"cfe46585", // closeCreditAccount(creditAccount,to,skipTokenMask,convertToETH,calls)
				// "5d91a0e0", // liquidateCreditAccount
				"37a8db9e": // openCreditAccount(debt,onBehalfOf,calls,referralCode)
				event, err := getCreditFacadeMainEvent(call.Input, creditFacadev3Parser)
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

var creditFacadev2Parser, creditFacadev3Parser *abi.ABI

func init() {
	creditFacadev2Parser = core.GetAbi("CreditFacade")
	creditFacadev3Parser = core.GetAbi("CreditFacadev3")
}
func getCreditFacadeMainEvent(input string, parser *abi.ABI) (*ds.FacadeCallNameWithMulticall, error) {
	hexData, err := hex.DecodeString(input[2:])
	if err != nil {
		return nil, err
	}
	method, err := parser.MethodById(hexData[:4])
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
