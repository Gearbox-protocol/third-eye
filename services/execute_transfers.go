package services

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/services/trace_service"
	"github.com/ethereum/go-ethereum/common"
)

// tenderly has logs for events(which we mainly use for Transfer on token) and balance_diff for native eth exchange
// handling native eth exchange is not needed for execution transfer or swaps
func (ef *ExecuteFilter) getExecuteTransfersv2(txLogs []trace_service.Log, cmEvents map[common.Hash]bool) []core.Transfers {
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
		if utils.Contains([]common.Hash{
			core.Topic("ExecuteOrder(address,address)"),
		}, eventSig) {
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

func (ef *ExecuteFilter) getExecuteTransfersv3(txLogs []trace_service.Log, cmEvents map[common.Hash]bool) (execEventBalances []core.Transfers) {
	balances := make(core.Transfers)
	paramsIndex := 0
	for _, raw := range txLogs {
		eventLog := raw.Raw
		eventSig := eventLog.Topics[0]
		eventLogAddress := eventLog.Address.Hex()
		//
		// ExecuteOrder
		if utils.Contains([]common.Hash{
			core.Topic("Execute(address,address)"),
		}, eventSig) {
			execEventBalances = append(execEventBalances, balances)
			paramsIndex += 1 // for getting the cm of next executeOrder
			balances = make(core.Transfers)
		}
		// if any other creditmanager event is emitted add to the execute
		if cmEvents[eventSig] {
			balances = make(core.Transfers)
		}

		// Transfer
		if eventSig == core.Topic("Transfer(address,address,uint256)") &&
			len(eventLog.Topics) == 3 {
			src := common.BytesToAddress(eventLog.Topics[1][:])
			dest := common.BytesToAddress(eventLog.Topics[2][:])
			amt, b := new(big.Int).SetString(eventLog.Data[2:], 16)
			if !b {
				log.Fatal("failed at serializing transfer data in int")
			}
			if paramsIndex >= len(ef.paramsList) {
				return
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
	return
}
