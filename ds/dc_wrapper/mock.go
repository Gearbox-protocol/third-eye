package dc_wrapper

import (
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
)

type DCTesting struct {
	calls  map[int64]*dc.DCCalls
	client core.ClientI
}

func (t DCTesting) getPoolData(blockNum int64, addr string) (dc.PoolCallData, error) {
	obj := t.calls[blockNum].Pools[addr]
	return obj, nil
}
func (t *DCTesting) getCMData(blockNum int64, key string) (dc.CMCallData, error) {
	if t.calls == nil || t.calls[blockNum] == nil {
		return dc.CMCallData{}, nil
	}
	obj := t.calls[blockNum].CMs[key]
	return obj, nil
}

func (t *DCTesting) getAccountData(blockNum int64, key string) (dc.CreditAccountCallData, error) {
	obj := t.calls[blockNum].Accounts[key]
	//
	var maskInBits string
	if obj.Version.IsGBv1() {
		mask := dc.GetMask(t.client, blockNum, core.NULL_ADDR, obj.Addr)
		maskInBits = fmt.Sprintf("%b", mask)
	}
	maskLen := len(maskInBits)
	//
	for ind, entry := range obj.Balances {
		if obj.Version.IsGBv1() {
			if maskLen-ind-1 >= 0 {
				entry.IsEnabled = maskInBits[maskLen-ind-1] == '1'
			}
		}
		entry.Ind = ind
		obj.Balances[ind] = entry
	}
	return obj, nil
}
