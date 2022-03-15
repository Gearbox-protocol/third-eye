package services

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Gearbox-protocol/third-eye/artifacts/curveV1Adapter"
	"github.com/Gearbox-protocol/third-eye/artifacts/iSwapRouter"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapV2Adapter"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapV3Adapter"
	"github.com/Gearbox-protocol/third-eye/artifacts/yearnAdapter"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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

func NewExecuteParser(config *config.Config) core.ExecuteParserI {
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

func (ep *ExecuteParser) getTenderlyData(txHash string) (*TxTrace, error) {
	link := fmt.Sprintf("https://api.tenderly.co/api/v1/public-contract/%d/trace/%s", ep.ChainId, txHash)
	req, _ := http.NewRequest(http.MethodGet, link, nil)
	resp, err := ep.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// log.Infof("%s",body)
	trace := &TxTrace{}
	err = json.Unmarshal(body, trace)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return trace, nil
}

// executeOrder

func (ep *ExecuteParser) GetTxTrace(txHash string) *TxTrace {
	trace, err := ep.getTenderlyData(txHash)
	if err != nil {
		log.Fatal(err)
	}
	if trace.CallTrace == nil {
		log.Info("Call trace nil retrying in 30 sec")
		time.Sleep(30 * time.Second)
		trace, err = ep.getTenderlyData(txHash)
		if err != nil {
			log.Fatal(err)
		}
		return trace
	}
	return trace
}

func (ep *ExecuteParser) GetExecuteCalls(txHash, creditManagerAddr string, paramsList []core.ExecuteParams) []*core.KnownCall {
	trace := ep.GetTxTrace(txHash)
	filter := ExecuteFilter{paramsList: paramsList, creditManager: common.HexToAddress(creditManagerAddr)}
	calls := filter.getExecuteCalls(trace.CallTrace)
	executeTransfers := filter.getExecuteTransfers(trace, ep.IgnoreCMEventIds)

	// check if parsed execute Order currently
	if len(calls) == len(executeTransfers) {
		for i, call := range calls {
			call.Transfers = executeTransfers[i]
		}
	} else {
		log.Fatalf("Calls %d execute details %d tx:%s creditManager:%s",
			len(calls), len(executeTransfers), txHash, creditManagerAddr)
	}
	return calls
}


func (ep *ExecuteParser) GetMainEventLogs(txHash, creditManager string) ([]*core.FuncWithMultiCall) {
	trace := ep.GetTxTrace(txHash)
	return ep.getOpenEvents(trace.CallTrace, common.HexToAddress(creditManager))
}
func (ep *ExecuteParser) GetTransfers(txHash string, accounts []string) (core.Transfers) {
	trace := ep.GetTxTrace(txHash)
	return ep.getTransfersToUser(trace, accounts)
}

func (ep *ExecuteParser) getOpenEvents(call *Call, creditManager common.Address) []*core.FuncWithMultiCall {
	openFuncs := []*core.FuncWithMultiCall{}
	if utils.Contains([]string{"CALL", "DELEGATECALL", "JUMP"}, call.CallerOp) {
		if creditManager ==  && len(call.Input) >= 10 {
			switch call.Input[:10] {
			case "":
				openFuncs = append(openFuncs, &core.FuncWithMultiCall{})
			}
		} else {
			for _, c := range call.Calls {
				c.Depth = call.Depth + 1
				openFuncs = append(openFuncs, ep.getOpenEvents(c, creditManager)...)
			}
		}
	}
	return openFuncs
}

func (ep *ExecuteParser) getTransfersToUser(trace *TxTrace, accounts []string) core.Transfers {
	transfers := core.Transfers{}
	for _, raw := range trace.Logs {
		eventLog := raw.Raw
		to := common.HexToAddress(eventLog.Topics[2])
		var isTransferToUser bool
		for _, account := range accounts {
			isTransferToUser = isTransferToUser || common.HexToAddress(account) == to
		}
		if eventLog.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" &&
			isTransferToUser {
				amt, b := new(big.Int).SetString(eventLog.Data[2:], 16)
				if !b {
					log.Fatal("failed at serializing transfer data in int")
				}
				token := common.HexToAddress(eventLog.Address).Hex()
				oldBalance := new(big.Int)
				if transfers[token] != nil {
					oldBalance = transfers[token]
				}
				transfers[token] = new(big.Int).Add(amt,oldBalance)
			}
	}
	return transfers
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
