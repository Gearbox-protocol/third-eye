package services

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/Gearbox-protocol/gearscan/artifacts/curveV1Adapter"
	"github.com/Gearbox-protocol/gearscan/artifacts/iSwapRouter"
	"github.com/Gearbox-protocol/gearscan/artifacts/uniswapV2Adapter"
	"github.com/Gearbox-protocol/gearscan/artifacts/uniswapV3Adapter"
	"github.com/Gearbox-protocol/gearscan/artifacts/yearnAdapter"
	"github.com/Gearbox-protocol/gearscan/config"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/utils"

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
	Client              http.Client
	IgnoreCMEventIds    []string
	ExecuteOrderFuncSig string
	ChainId             uint
}

func NewExecuteParser(config *config.Config) *ExecuteParser {
	return &ExecuteParser{
		Client:              http.Client{},
		IgnoreCMEventIds:    utils.GetCreditManagerEventIds(),
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
	CallTrace *Call  `json:"call_trace"`
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

func (ep *ExecuteParser) GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ExecuteParams) []*KnownCall {
	trace, err := ep.GetTxTrace(txHash)
	if err != nil {
		log.Fatal(err)
	}
	filter := ExecuteFilter{paramsList: paramsList, creditManager: common.HexToAddress(creditManagerAddr)}
	calls := filter.getExecuteCalls(trace.CallTrace)
	executeTransfers := filter.getExecuteTransfers(trace, ep.IgnoreCMEventIds)

	// check if parsed execute Order currently
	if len(calls) == len(executeTransfers) && len(calls) == len(paramsList) {
		for i, call := range calls {
			call.Balances = executeTransfers[i]
			call.LogId = int(paramsList[i].Index)
		}
	} else {
		log.Fatal("Calls ", len(calls), ", execute details ", len(executeTransfers), ", execute params", len(paramsList))
	}
	return calls
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
