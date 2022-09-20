package services

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/curveV1Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/iSwapRouter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/testAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Adapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnAdapter"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"

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
	Client              http.Client
	IgnoreCMEventIds    map[string]bool
	ExecuteOrderFuncSig string
	ChainId             uint
}

func getCMEventIds() map[string]bool {
	ids := map[string]bool{}
	if abiObj, err := abi.JSON(strings.NewReader(creditFacade.CreditFacadeABI)); err == nil {
		for _, event := range abiObj.Events {
			ids[event.ID.Hex()] = true
		}
	}
	if abiObj, err := abi.JSON(strings.NewReader(creditManager.CreditManagerABI)); err == nil {
		for _, event := range abiObj.Events {
			ids[event.ID.Hex()] = true
		}
	}
	return ids
}
func NewExecuteParser(config *config.Config) ds.ExecuteParserI {
	return &ExecuteParser{
		Client:              http.Client{},
		IgnoreCMEventIds:    getCMEventIds(),
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

func (ep *ExecuteParser) GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ds.ExecuteParams) []*ds.KnownCall {
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

var abiJSONs = []string{curveV1Adapter.CurveV1AdapterABI, yearnAdapter.YearnAdapterABI,
	uniswapv2Adapter.Uniswapv2AdapterABI, uniswapv3Adapter.Uniswapv3AdapterABI,
	iSwapRouter.ISwapRouterABI, testAdapter.TestAdapterABI,
	// creditfacade for credit manager onlogs
}

var abiParsers []abi.ABI
var creditFacadeParser abi.ABI

func init() {
	for _, abiJSON := range abiJSONs {
		abiParser, err := abi.JSON(strings.NewReader(abiJSON))
		if err != nil {
			log.Fatal(err)
		}
		abiParsers = append(abiParsers, abiParser)
	}

	abiParser, err := abi.JSON(strings.NewReader(creditFacade.CreditFacadeABI))
	if err != nil {
		log.Fatal(err)
	}
	creditFacadeParser = abiParser
}

//////////////////////////
// parser functions for v2
//////////////////////////
// GetMainEventLogs
func (ep *ExecuteParser) GetMainEventLogs(txHash, creditFacade string) []*ds.MainactionWithMulticall {
	trace := ep.GetTxTrace(txHash)
	data, err := ep.getMainEvents(trace.CallTrace, common.HexToAddress(creditFacade))
	if err != nil {
		log.Fatal(err.Error(), "for txHash", txHash)
	}
	return data
}

func (ep *ExecuteParser) getMainEvents(call *Call, creditFacade common.Address) ([]*ds.MainactionWithMulticall, error) {
	mainEvents := []*ds.MainactionWithMulticall{}
	if utils.Contains([]string{"CALL", "DELEGATECALL", "JUMP"}, call.CallerOp) {
		if creditFacade == common.HexToAddress(call.To) && len(call.Input) >= 10 {
			switch call.Input[2:10] {
			case "caa5c23f", // multicall
				"5f73fbec", // closeCreditAccount
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
				c.Depth = call.Depth + 1
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

func getCreditFacadeMainEvent(input string) (*ds.MainactionWithMulticall, error) {
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
		if hex.EncodeToString(call.CallData[:4]) == "81314b59" { // ignore revertifreceivedlessthan
			continue
		}
		multicalls = append(multicalls, multicall.Multicall2Call{
			Target:   call.Target,
			CallData: call.CallData,
		})
	}
	return &ds.MainactionWithMulticall{
		Name:       method.Name,
		MultiCalls: multicalls,
	}, nil
}

/// GetTransfers
func (ep *ExecuteParser) GetTransfers(txHash, borrower, account, underlyingToken string, accounts []string) core.Transfers {
	trace := ep.GetTxTrace(txHash)
	return ep.getTransfersToUser(trace, borrower, account, underlyingToken, accounts)
}

func (ep *ExecuteParser) getTransfersToUser(trace *TxTrace, borrower, account, underlyingToken string, accounts []string) core.Transfers {
	transfers := core.Transfers{}
	for _, raw := range trace.Logs {
		eventLog := raw.Raw
		if eventLog.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" { // transfer event
			to := common.HexToAddress(eventLog.Topics[2])
			from := common.HexToAddress(eventLog.Topics[1]).Hex()
			token := common.HexToAddress(eventLog.Address).Hex()
			var sign *big.Int
			if from == borrower && to.Hex() == account && token == underlyingToken {
				sign = big.NewInt(-1)
			} else {
				var isTransferToUser bool
				for _, account := range accounts {
					isTransferToUser = isTransferToUser || common.HexToAddress(account) == to
				}
				if !isTransferToUser {
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
