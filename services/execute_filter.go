package services

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ExecuteFilter struct {
	paramsList    []ds.ExecuteParams
	paramsIndex   int
	creditManager common.Address
}

func (ef *ExecuteFilter) getExecuteCalls(call *Call) []*ds.KnownCall {
	var calls []*ds.KnownCall
	if ef.paramsIndex >= len(ef.paramsList) {
		return calls
	}
	ep := ef.paramsList[ef.paramsIndex]
	if utils.Contains([]string{"CALL", "DELEGATECALL", "JUMP"}, call.CallerOp) {
		// Execute call on credit manager
		if ef.creditManager == common.HexToAddress(call.To) && len(call.Input) >= 10 && call.Input[:10] == "0x6ce4074a" {

			dappcall := call.dappCall(ep.Protocol)
			// this check is there as there are 2 executeOrder call in
			// https://kovan.etherscan.io/tx/0x9aeb9ccfb3e100c3c9e6ed5a140784e910a962be36e15f244938645b21c48a96
			// only first call to the dapp as the gearbox don't recursively call adapter/creditManager executeOrder
			dappcall.Depth = call.Depth
			calls = append(calls, dappcall)
			ef.paramsIndex += 1
		} else {
			for _, c := range call.Calls {
				c.Depth = call.Depth + 1
				calls = append(calls, ef.getExecuteCalls(c)...)
			}
		}
	}
	return calls
}

func (call *Call) dappCall(dappAddr common.Address) *ds.KnownCall {
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
		knownCall := c.dappCall(dappAddr)
		if knownCall != nil {
			return knownCall
		}
	}
	return nil
}

func (ef *ExecuteFilter) getExecuteTransfers(trace *TxTrace, cmEvents []string) []core.Transfers {
	balances := make(core.Transfers)
	var execEventBalances []core.Transfers
	parsingTransfer := false
	paramsIndex := -1
	for _, raw := range trace.Logs {
		eventLog := raw.Raw
		eventSig := eventLog.Topics[0]
		eventLogAddress := common.HexToAddress(eventLog.Address).Hex()
		// if any creditmanager event add to the execute
		if utils.Contains(cmEvents, eventSig) && parsingTransfer {
			execEventBalances = append(execEventBalances, balances)
			balances = make(core.Transfers)
			parsingTransfer = false
		}
		// ExecuteOrder
		if eventSig == "0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03" {
			paramsIndex += 1
			balances = make(core.Transfers)
			parsingTransfer = true
		}
		// Transfer
		if eventSig == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" &&
			len(eventLog.Topics) == 3 && parsingTransfer {
			src := common.HexToAddress(eventLog.Topics[1])
			dest := common.HexToAddress(eventLog.Topics[2])
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
