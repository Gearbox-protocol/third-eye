package framework 

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type TxLogList []types.Log


func (ts TxLogList) Len() int {
	return len(ts)
}
func (ts TxLogList) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts TxLogList) Less(i, j int) bool {
	return ts[i].BlockNumber < ts[j].BlockNumber || 
	(ts[i].BlockNumber == ts[j].BlockNumber && ts[i].Index < ts[j].Index)
}