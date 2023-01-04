package services

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/convexAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/curveAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/curveV1Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/iSwapRouter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/lidov1Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/lidov1Gateway"
	"github.com/Gearbox-protocol/sdk-go/artifacts/testAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/universalAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/wstETHv1Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnv2Adapter"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/services/trace_service"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// has code only for exeuction call

type ExecuteFilter struct {
	paramsList    []ds.ExecuteParams
	paramsIndex   int
	creditManager common.Address
}

func (ef *ExecuteFilter) getExecuteCalls(call *trace_service.Call) []*ds.KnownCall {
	var calls []*ds.KnownCall
	if ef.paramsIndex >= len(ef.paramsList) {
		return calls
	}
	ep := ef.paramsList[ef.paramsIndex]
	if utils.Contains([]string{"CALL", "DELEGATECALL", "JUMP"}, call.CallerOp) {
		// Execute call on credit manager
		if ef.creditManager == common.HexToAddress(call.To) && len(call.Input) >= 10 && call.Input[:10] == "0x6ce4074a" {
			dappcall := dappCall(call, ep.Protocol)
			// this check is there as there are 2 executeOrder call in
			// https://kovan.etherscan.io/tx/0x9aeb9ccfb3e100c3c9e6ed5a140784e910a962be36e15f244938645b21c48a96
			// only first call to the dapp as the gearbox don't recursively call adapter/creditManager executeOrder
			calls = append(calls, dappcall)
			ef.paramsIndex += 1
		} else {
			for _, c := range call.Calls {
				calls = append(calls, ef.getExecuteCalls(c)...)
			}
		}
	}
	return calls
}

// this is called after ExecuteOrder event is seen on credit manager for both v1 and v2
func dappCall(call *trace_service.Call, dappAddr common.Address) *ds.KnownCall {
	if utils.Contains([]string{"CALL", "DELEGATECALL", "JUMP"}, call.CallerOp) && dappAddr == common.HexToAddress(call.To) {
		name, arguments := ParseCallData(call.Input)
		if arguments == nil {
			log.Fatalf("%s %#v %#v\n", name, arguments, call)
		}
		return &ds.KnownCall{
			Name: name,
			Args: arguments,
		}
	}
	for _, c := range call.Calls {
		knownCall := dappCall(c, dappAddr)
		if knownCall != nil {
			return knownCall
		}
	}
	return nil
}

// tenderly has logs for events(which we mainly use for Transfer on token) and balance_diff for native eth exchange
// handling native eth exchange is not needed for execution transfer or swaps
func (ef *ExecuteFilter) getExecuteTransfers(txLogs []trace_service.Log, cmEvents map[common.Hash]bool) []core.Transfers {
	balances := make(core.Transfers)
	var execEventBalances []core.Transfers
	parsingTransfer := false
	paramsIndex := -1
	for _, raw := range txLogs {
		eventLog := raw.Raw
		eventSig := eventLog.Topics[0]
		eventLogAddress := eventLog.Address.Hex()
		// if any other creditmanager event is emitted add to the execute
		if cmEvents[eventSig] && parsingTransfer {
			execEventBalances = append(execEventBalances, balances)
			balances = make(core.Transfers)
			parsingTransfer = false
		}
		// ExecuteOrder
		if eventSig == core.Topic("ExecuteOrder(address,address)") {
			paramsIndex += 1
			balances = make(core.Transfers)
			parsingTransfer = true
		}
		// Transfer
		if eventSig == core.Topic("Transfer(address,address,uint256)") &&
			len(eventLog.Topics) == 3 && parsingTransfer {
			src := common.BytesToAddress(eventLog.Topics[1][:])
			dest := common.BytesToAddress(eventLog.Topics[2][:])
			amt, b := new(big.Int).SetString(eventLog.Data[2:], 16)
			if !b {
				log.Fatal("failed at serializing transfer data in int")
			}
			creditAccount := ef.paramsList[paramsIndex].CreditAccount
			if balances[eventLogAddress] == nil {
				balances[eventLogAddress] = big.NewInt(0)
			}
			if src == creditAccount {
				balances[eventLogAddress] = new(big.Int).Sub(balances[eventLogAddress], amt)
			} else if dest == creditAccount {
				balances[eventLogAddress] = new(big.Int).Add(balances[eventLogAddress], amt)
			}
		}
	}
	if parsingTransfer {
		execEventBalances = append(execEventBalances, balances)
	}
	return execEventBalances
}

var abiJSONs = []string{curveV1Adapter.CurveV1AdapterABI, yearnAdapter.YearnAdapterABI,
	uniswapv2Adapter.Uniswapv2AdapterABI, uniswapv3Adapter.Uniswapv3AdapterABI,
	iSwapRouter.ISwapRouterABI, testAdapter.TestAdapterABI,
	// creditfacade for credit manager onlogs
	// v2
	lidov1Adapter.Lidov1AdapterABI, lidov1Gateway.Lidov1GatewayABI, wstETHv1Adapter.WstETHv1AdapterABI,
	convexAdapter.ConvexAdapterABI, curveAdapter.CurveAdapterABI,
	yearnv2Adapter.Yearnv2AdapterABI, universalAdapter.UniversalAdapterABI,
}

var abiParsers []abi.ABI

func init() {
	for _, abiJSON := range abiJSONs {
		abiParser, err := abi.JSON(strings.NewReader(abiJSON))
		if err != nil {
			log.Fatal(err)
		}
		abiParsers = append(abiParsers, abiParser)
	}
}

//https://ethereum.stackexchange.com/questions/29809/how-to-decode-input-data-with-abi-using-golang/100247
func ParseCallData(input string) (string, *core.Json) {
	hexData, err := hex.DecodeString(input[2:])
	if err != nil {
		log.Fatal(err)
	}
	for _, parser := range abiParsers {
		// check if the methods for parser matches the input sig
		method, err := parser.MethodById(hexData[:4])
		if err != nil {
			continue
		}
		// unpack in the map
		data := map[string]interface{}{}
		err = method.Inputs.UnpackIntoMap(data, hexData[4:])
		if err != nil {
			log.Fatal(err)
		}
		// add order
		var argNames []interface{}
		for _, input := range method.Inputs {
			argNames = append(argNames, input.Name)
		}
		data["_order"] = argNames
		jsonData := core.Json(data)
		return method.Sig, &jsonData
	}
	log.Fatal("No method for input: ", input)
	return "", nil
}
