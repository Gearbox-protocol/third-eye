package services

import (
	"encoding/hex"
	"reflect"
	"strings"

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
	// no need to get the logs as mainCalls doesn't deal with logs
	trace := ep.GetTxTrace(txHash, false)
	// log.Info(utils.ToJson(trace))
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
			// v2
			case "caa5c23f", // multicall
				"5f73fbec", // closeCreditAccount
				"82871ace", // liquidateExpiredCreditAccount
				"5d91a0e0", // liquidateCreditAccount (v2)
				"7071b7c5": // openCreditAccountMulticall
				event, err := getCreditFacadeMainEvent(call.To, call.Input, creditFacadev2Parser)
				if err != nil {
					return nil, err
				}
				mainEvents = append(mainEvents, event)
				// v3
			case "ebe4107c", // multicall(address,calls) // v310 , v3
				"7e2ca9db", // botMulticall(address,calls)
				"e3f46b26", // liquidateCreditAccount (v3) // v310, v3
				"36b2ced3", // closeCreditAccount(creditAccount,to,skipTokenMask,convertToETH,calls)  // v310, v3
				// "5d91a0e0", // liquidateCreditAccount
				//				"85589e10", //Partial
				"92beab1d": // openCreditAccount(onBehalfOf,calls,referralCode) // v310, v3
				// log.Info("v3 main event", call.Input[2:10])
				event, err := getCreditFacadeMainEvent(call.To, call.Input, creditFacadev3Parser)
				if err != nil {
					return nil, err
				}
				// log.Info(utils.ToJson(event.GetMulticalls()))
				mainEvents = append(mainEvents, event)
			default:
				log.Warn(call.Input[2:10], "1234")
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
func getABI(data string) *abi.ABI {
	abi, err := abi.JSON(strings.NewReader(data))
	log.CheckFatal(err)
	return &abi
}

var pp = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repaidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSeizedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"struct PriceUpdate[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"name\":\"partiallyLiquidateCreditAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seizedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

func getCreditFacadeMainEvent(contract string, input string, parser *abi.ABI) (*ds.FacadeCallNameWithMulticall, error) {
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
		contract, // v2 and v3 it means facade
		ds.FacadeAccountMethodSigToCallName(method.Name),
		multicalls,
	), nil
}
