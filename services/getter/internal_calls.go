package getter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type InternalCallFetcher struct {
	rpc    string
	client http.Client
}

func NewInternalCallFetcher(rpc string) *InternalCallFetcher {
	return &InternalCallFetcher{
		rpc:    rpc,
		client: http.Client{},
	}
}

func (app InternalCallFetcher) getData(txHash string) ([]RPCTrace, error) {
	format := `{"id":1,"jsonrpc":"2.0","params":["%s"],"method":"trace_transaction"}`
	params := fmt.Sprintf(format, txHash)
	//
	buf := &bytes.Buffer{}
	buf.WriteString(params)
	req, _ := http.NewRequest(http.MethodPost, app.rpc, buf)
	resp, err := app.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("While making request %s", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("While reading body from response %s", err)
	}
	fmt.Println(utils.ToJson(data))
	trace := []RPCTrace{}
	err = json.Unmarshal(data, &trace)
	if err != nil {
		return nil, fmt.Errorf("While unmarshaling %s", err)
	}
	return trace, nil
}

func (app InternalCallFetcher) GetData(txHash string) *TenderlyTrace {
	rpcTrace, err := app.getData(txHash)
	log.CheckFatal(err)
	return getTenderlyCall(rpcTrace, txHash)
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

func getTenderlyCall(old []RPCTrace, txHash string) *TenderlyTrace {
	call, _ := toCall(old[0])
	for _, rpcEntry := range old[1:] {
		nextCall, path := toCall(rpcEntry)
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

func toCall(old RPCTrace) (*Call, []int) {
	callerOp := strings.ToUpper(old.Action.CallType)
	var valueStr string
	if callerOp != "STATICCALL" {
		value, err := strconv.ParseInt(old.Action.Value[2:], 16, 64)
		log.CheckFatal(err)
		valueStr = fmt.Sprintf("%d", value)
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
