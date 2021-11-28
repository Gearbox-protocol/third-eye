package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/Gearbox-protocol/gearscan/artifacts/creditManager"
	"github.com/Gearbox-protocol/gearscan/artifacts/curveV1Adapter"
	"github.com/Gearbox-protocol/gearscan/artifacts/iSwapRouter"
	"github.com/Gearbox-protocol/gearscan/artifacts/uniswapV2Adapter"
	"github.com/Gearbox-protocol/gearscan/artifacts/uniswapV3Adapter"
	"github.com/Gearbox-protocol/gearscan/artifacts/yearnAdapter"
	"github.com/Gearbox-protocol/gearscan/config"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func getCreditManagerEventIds() []string {
	var ids []string
	if a, err := abi.JSON(strings.NewReader(creditManager.CreditManagerABI)); err == nil {
		for _, event := range a.Events {
			if event.RawName != "ExecuteOrder" {
				// fmt.Println(event.RawName, event.ID.Hex())
				ids = append(ids, event.ID.Hex())
			}
		}
	}
	return ids
}

func NewExecuteParser(config *config.Config) *ExecuteParser {
	return &ExecuteParser{
		Client:              http.Client{},
		IgnoreCMEventIds:    getCreditManagerEventIds(),
		ExecuteOrderFuncSig: "0x6ce4074a",
		ChainId:             config.ChainId,
	}
}

type Call struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	CallerOp string  `json:"caller_op"`
	Input    string  `json:"input"`
	Value    string  `json:"value"`
	Calls    []*Call `json:"calls"`
	Depth    uint8
}

type Log struct {
	Name string `json:"name"`
	Raw  struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
		Data    string   `json:"data"`
	} `json:"raw"`
}

type TxTrace struct {
	CallTrace Call   `json:"call_trace"`
	TxHash    string `json:"transaction_id"`
	Logs      []Log  `json:"logs"`
}

func (ep *ExecuteParser) GetTxTrace(txHash string) (*TxTrace, error) {
	link := fmt.Sprintf("https://api.tenderly.co/api/v1/public-contract/%d/trace/%s", ep.ChainId, txHash)
	req, _ := http.NewRequest(http.MethodGet, link, nil)
	resp, err := ep.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	trace := &TxTrace{}
	err = json.NewDecoder(resp.Body).Decode(trace)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return trace, nil
}

// executeOrder

type Balances map[string]*big.Int

func (bal *Balances) String() string {
	var str string
	first := true
	for addr, amt := range (map[string]*big.Int)(*bal) {
		if !first {
			str += ","
		}
		str += fmt.Sprintf("%s=>%s", addr, amt.String())
		first = false
	}
	return str
}

type KnownCall struct {
	// Input string
	From     common.Address
	To       common.Address
	Depth    uint8
	Name     string
	Args     string
	Balances Balances
	LogId    int
}

func (call *Call) dappCall(dappAddr common.Address) *KnownCall {
	if (call.CallerOp == "CALL" || call.CallerOp == "DELEGATECALL") && dappAddr == common.HexToAddress(call.To) {
		name, arguments := ParseCallData(call.Input)
		return &KnownCall{
			From: common.HexToAddress(call.From),
			To:   common.HexToAddress(call.To),
			Name: name,
			Args: arguments,
		}
	}
	for _, c := range call.Calls {
		knownCall := c.dappCall(dappAddr)
		if knownCall != nil {
			return knownCall
		}
	}
	return nil
}

func (call *Call) getManagerCalls(managerAddr, dappAddr common.Address) []*KnownCall {
	var calls []*KnownCall
	if call.CallerOp == "CALL" || call.CallerOp == "DELEGATECALL" {
		if managerAddr == common.HexToAddress(call.To) && call.Input[:10] == "0x6ce4074a" {

			dappcall := call.dappCall(dappAddr)
			dappcall.Depth = call.Depth
			// only first call to the dapp as the gearbox don't recursively call adapter/creditManager executeOrder
			calls = append(calls, dappcall)
		} else {
			for _, c := range call.Calls {
				c.Depth = call.Depth + 1
				calls = append(calls, c.getManagerCalls(managerAddr, dappAddr)...)
			}
		}
	}
	return calls
}

func (tt *TxTrace) GetKnownInternalCalls(managerAddr, dappAddr common.Address) []*KnownCall {
	tt.CallTrace.Depth = 1
	knowncalls := tt.CallTrace.getManagerCalls(managerAddr, dappAddr)

	return knowncalls
}

type ExecuteParser struct {
	Client              http.Client
	IgnoreCMEventIds    []string
	ExecuteOrderFuncSig string
	ChainId             uint
}

type ExecuteDetails struct {
	Balances Balances
	LogId    int
}

func (ep *ExecuteParser) GetExecuteCalls(txLog *types.Log, creditAccount, dappAddr common.Address) []*KnownCall {
	trace, err := ep.GetTxTrace(txLog.TxHash.Hex())
	if err != nil {
		log.Fatal(err)
	}
	calls := trace.GetKnownInternalCalls(txLog.Address, dappAddr)

	firstExecuteRelativeIndex := -1
	balances := make(Balances)
	var executeDetails []ExecuteDetails
	lastExecuteIndex := -1
	for i, raw := range trace.Logs {
		eventLog := raw.Raw
		eventSig := eventLog.Topics[0]
		eventLogAddress := common.HexToAddress(eventLog.Address).Hex()
		// fmt.Printf("%s %+v %+v %+v\n", raw.Name, eventSig,executeDetails, balances)
		if Contains(ep.IgnoreCMEventIds, eventSig) && lastExecuteIndex != -1 {
			executeDetails = append(executeDetails, ExecuteDetails{Balances: balances, LogId: lastExecuteIndex})
			balances = make(Balances)
			lastExecuteIndex = -1
			// ExecuteOrder
		} else if eventSig == "0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03" {
			if firstExecuteRelativeIndex == -1 {
				firstExecuteRelativeIndex = i
			}
			lastExecuteIndex = int(txLog.Index) + i - firstExecuteRelativeIndex
			balances = make(Balances)
			// Transfer
		} else if eventSig == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && len(eventLog.Topics) == 3 {
			src := common.HexToAddress(eventLog.Topics[1])
			dest := common.HexToAddress(eventLog.Topics[2])
			amt, b := new(big.Int).SetString(eventLog.Data[2:], 16)
			if !b {
				log.Fatal("failed at serializing transfer data in int")
			}
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
	if lastExecuteIndex != -1 {
		executeDetails = append(executeDetails, ExecuteDetails{Balances: balances, LogId: lastExecuteIndex})
	}
	// fmt.Printf("%+v\n",calls[0])
	if len(calls) == len(executeDetails) {
		for i, call := range calls {
			ed := executeDetails[i]
			call.Balances = ed.Balances
			call.LogId = ed.LogId
		}
	} else {
		log.Fatal("Calls ", len(calls), ", execute details ", len(executeDetails))
	}
	return calls
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var abiJSONs = []string{curveV1Adapter.CurveV1AdapterABI, yearnAdapter.YearnAdapterABI,
	uniswapV2Adapter.UniswapV2AdapterABI, uniswapV3Adapter.UniswapV3AdapterABI,
	iSwapRouter.ISwapRouterABI}

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
func ParseCallData(input string) (string, string) {
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
		// json marshal
		var args []byte
		args, err = json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		return method.Sig, string(args)
	}
	log.Fatal("No method for input: ", input)
	return "", ""
}
