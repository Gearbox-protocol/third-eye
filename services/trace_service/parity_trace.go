package trace_service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

// https://docs.alchemy.com/reference/trace-api
type ParityFetcher struct {
	rpcs   []string
	client http.Client
}

func NewParityFetcher(rpc string) *ParityFetcher {
	return &ParityFetcher{
		rpcs:   strings.Split(rpc, ","),
		client: http.Client{},
	}
}

type traceResp struct {
	Error struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result []RPCTrace
}
type traceRespQN struct {
	Error struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result QNRPCTrace
}

// https://docs.alchemy.com/reference/trace-transaction
func (app ParityFetcher) getDataOnRPC(rpc, txHash string) ([]RPCTrace, error) {
	format := `{"id":1,"jsonrpc":"2.0","params":["%s"],"method":"trace_transaction"}`
	params := fmt.Sprintf(format, txHash)
	//
	buf := &bytes.Buffer{}
	buf.WriteString(params)
	req, _ := http.NewRequest(http.MethodPost, rpc, buf)
	req.Header.Add("Content-Type", "application/json")
	//
	resp, err := app.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("while making request %s", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("while reading body from response %s", err)
	}
	traceObj := traceResp{}
	err = json.Unmarshal(data, &traceObj)
	if err != nil {
		return nil, fmt.Errorf("while unmarshaling %s", err)
	}
	if traceObj.Error.Code != 0 {
		return nil, fmt.Errorf(traceObj.Error.Message)
	}
	if len(traceObj.Result) == 0 {
		log.Info(txHash, string(data))
	}
	return traceObj.Result, nil
}
func (app ParityFetcher) getDataOnQuicknode(rpc, txHash string) (*QNRPCTrace, error) {
	format := `{"method":"debug_traceTransaction","params":["%s", {"tracer": "callTracer"}], "id":1,"jsonrpc":"2.0"}`
	params := fmt.Sprintf(format, txHash)
	//
	buf := &bytes.Buffer{}
	buf.WriteString(params)
	req, _ := http.NewRequest(http.MethodPost, rpc, buf)
	req.Header.Add("Content-Type", "application/json")
	//
	resp, err := app.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("while making request %s", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("while reading body from response %s", err)
	}
	traceObj := traceRespQN{}
	err = json.Unmarshal(data, &traceObj)
	if err != nil {
		return nil, fmt.Errorf("while unmarshaling %s", err)
	}
	if traceObj.Error.Code != 0 {
		return nil, fmt.Errorf(traceObj.Error.Message)
	}
	return &traceObj.Result, nil
}

func (app ParityFetcher) getData(txhash string, providedrpc ...string) ([]RPCTrace, error) {
	var errs utils.Errors
	rpc := app.rpcs
	if len(providedrpc) != 0 {
		rpc = providedrpc
	}
	for _, rpc := range rpc {
		data, err := app.getDataOnRPC(rpc, txhash)
		if err == nil {
			return data, nil
		}
		errs = append(errs, err)
	}
	return nil, errs
}

func (app ParityFetcher) getTxTrace(txHash string, rpc ...string) (*TenderlyTrace, error) {
	rpcTrace, err := app.getData(txHash, rpc...)
	if err != nil {
		return nil, err
	}
	return convertToTenderlyTrace(rpcTrace, txHash), nil
}
func (app ParityFetcher) getTxTraceQuickNode(txHash string, rpc string) (*TenderlyTrace, error) {
	rpcTrace, err := app.getDataOnQuicknode(rpc, txHash)
	if err != nil {
		return nil, err
	}
	client, err := ethclient.Dial(rpc)
	log.CheckFatal(err)
	tx, err := client.TransactionReceipt(context.TODO(), common.HexToHash(txHash))
	log.CheckFatal(err)
	return &TenderlyTrace{
		CallTrace:   rpcTrace.Convert(),
		BlockNumber: tx.BlockNumber.Int64(),
		TxHash:      txHash,
	}, nil
}

type RPCTrace struct {
	Action struct {
		From     string `json:"from"`
		To       string `json:"to"`
		CallType string `json:"callType"`
		Value    string `json:"value"`
		Input    string `json:"input"`
	} `json:"action"`
	BlockNumber  int64 `json:"blockNumber"`
	Subtraces    int   `json:"subtraces"`
	TraceAddress []int `json:"traceAddress"`
}

func convertToTenderlyTrace(old []RPCTrace, txHash string) *TenderlyTrace {
	call, _ := toTenderlyCall(old[0], txHash)
	for _, rpcEntry := range old[1:] {
		nextCall, path := toTenderlyCall(rpcEntry, txHash)
		cur := call
		for _, step := range path[:len(path)-1] {
			cur = cur.Calls[step]
		}
		cur.Calls[path[len(path)-1]] = nextCall
	}
	return &TenderlyTrace{
		CallTrace:   call,
		BlockNumber: old[0].BlockNumber,
		TxHash:      txHash,
	}
}

func toTenderlyCall(old RPCTrace, txHash string) (*Call, []int) {
	callerOp := strings.ToUpper(old.Action.CallType)
	var valueStr string
	if callerOp != "STATICCALL" {
		if len(old.Action.Value) >= 2 {
			value, ok := new(big.Int).SetString(old.Action.Value[2:], 16)
			if !ok {
				log.Errorf("For txhash (%s) can't parse ethValue %s", txHash, old.Action.Value)
				valueStr = "0"
			} else {
				valueStr = value.String()
			}
		} else {
			log.Errorf("For txhash (%s) can't parse old.Action.value %s", txHash, utils.ToJson(old))
			valueStr = "0"
		}
	}
	var calls []*Call
	if old.Subtraces > 0 {
		calls = make([]*Call, old.Subtraces)
	}
	return &Call{
		From:     old.Action.From,
		To:       old.Action.To,
		CallerOp: strings.ToUpper(old.Action.CallType),
		Input:    old.Action.Input,
		Value:    valueStr,
		Calls:    calls,
	}, old.TraceAddress
}

type QNRPCTrace struct {
	From string `json:"from"`
	To   string `json:"to"`
	// CallType string       `json:"callType"`
	Value  string       `json:"value"`
	Type   string       `json:"type"`
	Input  string       `json:"input"`
	Output string       `json:"output"`
	Calls  []QNRPCTrace `json:"calls"`
	// Action struct {
	// } `json:"action"`
	// BlockNumber  int64 `json:"blockNumber"`
	// Subtraces    int   `json:"subtraces"`
	// TraceAddress []int `json:"traceAddress"`
}

func (x QNRPCTrace) Convert() *Call {
	calls := make([]*Call, len(x.Calls))
	for i, c := range x.Calls {
		calls[i] = c.Convert()
	}
	return &Call{
		From:     x.From,
		To:       x.To,
		CallerOp: strings.ToUpper(x.Type),
		Input:    x.Input,
		Value:    x.Value,
		Calls:    calls,
	}
}
