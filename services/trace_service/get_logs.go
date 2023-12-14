package trace_service

import (
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type TxLogger struct {
	node     pkg.Node
	store    map[int]map[string][]Log
	nums     []int
	storeLen int
}

func NewTxLogger(client core.ClientI, storeLen int64) TxLogger {
	return TxLogger{
		node:     pkg.Node{Client: client},
		store:    map[int]map[string][]Log{},
		storeLen: int(storeLen),
		nums:     make([]int, 0, storeLen/2),
	}
}

// works only for logs of a txHash which has executeOrder or execute or closeCreditAccount in them
func (m *TxLogger) GetLogs(blockNum int, txHash string) []Log {
	if m.store[blockNum] == nil || m.store[blockNum][txHash] == nil {
		m.nums = insertInSlice(m.nums, int(blockNum))
		m.store[blockNum] = m.fetchLogs(int64(blockNum))
	}
	//
	ansTxLogs := m.store[blockNum][txHash]
	// delete(m.store[blockNum], txHash)
	m.nums = deleteInSlice(m.nums, blockNum-m.storeLen, m.store)
	if len(ansTxLogs) == 0 {
		log.Fatal("TxLogger returned 0 logs for ", txHash)
	}
	return ansTxLogs
}

type operator struct {
	curTxHash      string
	storeCurTxHash bool
}

// it fetches all logs for a block,
// filter all tx that have logs which are atleast one of Execute, ExecuteOrder or CloseCreditAccount
// then it returns a map of txHash to logs
func (m TxLogger) fetchLogs(blockNum int64) map[string][]Log {
	//
	txLogs, err := m.node.GetLogs(blockNum, blockNum, nil, nil)
	if err != nil {
		log.Fatalf("Err(%s) while getting logs from etherscan for ", err, blockNum)
	}
	logStore := map[string][]Log{}
	op := operator{}
	for _, txLog := range txLogs {
		newTxHash := txLog.TxHash.Hex()
		op.next(logStore, newTxHash)
		//
		formattedLog := Log{
			Name: "",
			Raw: RawLog{
				Address: txLog.Address,
				Topics:  txLog.Topics,
				Data:    fmt.Sprintf("0x%s", hex.EncodeToString(txLog.Data)),
			},
		}

		logStore[newTxHash] = append(logStore[newTxHash], formattedLog)
		//
		valid := len(txLog.Topics) > 0 && utils.Contains([]common.Hash{
			core.Topic("ExecuteOrder(address,address)"),       // executeOrder
			core.Topic("CloseCreditAccount(address,address)"), // close v2
			//v3
			core.Topic("Execute(address,address)"),            // execute on v3
			core.Topic("CloseCreditAccount(address,address)"), // close on v3
		}, txLog.Topics[0])
		op.storeCurTxHash = op.storeCurTxHash || valid
	}
	op.next(logStore, "")
	//
	return logStore
}

func (op *operator) next(logStore map[string][]Log, newTxHash string) {
	if op.curTxHash != newTxHash {
		// do operation
		if !op.storeCurTxHash {
			delete(logStore, op.curTxHash)
		}
		// set fields
		op.curTxHash = newTxHash
		op.storeCurTxHash = false
	}
}

// slice

func insertInSlice(nums []int, block int) []int {
	// search ints return just geaer ind
	ind := sort.SearchInts(nums, block)
	if !(ind != 0 && nums[ind-1] == block) {
		// if cap(nums) > len(nums) {
		nums = append(nums[:ind+1], nums[ind:]...)
		nums[ind] = block
		// } else {
		// 	x := []int,0, cap(nums)
		// 	x = append(x, nums[:ind]...)
		// 	x = append(x, block)
		// 	x = append(x, nums[ind:]...)
		// 	nums = x
		// }
	}
	return nums
}

func deleteInSlice(nums []int, deleteTill int, store map[int]map[string][]Log) []int {
	// search ints return just geaer ind
	ind := sort.SearchInts(nums, int(deleteTill))
	for _, block := range nums[:ind] {
		delete(store, int(block))
	}
	return append(nums[:0], nums[ind:]...)
}
