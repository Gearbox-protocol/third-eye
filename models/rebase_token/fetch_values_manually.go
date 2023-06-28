package rebase_token

import (
	"math/big"
	"os"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
)

var validatorsKey string = "0x9f70001d82b6ef54e9d3725b46581c3eb9ee3aa02b941b6aa54d678a9ca35b10"
var totalSharesKey string = "0xe3b4b636e601189b5f4c6742edf2538ac12bb61ed03e6da26949d69838fa447e"

func FetchRebaseDetails() {
	client := ethclient.NewEthClient(&config.Config{EthProvider: ""})

	var start int64 = 13810899
	var end int64 = 17266004

	m := map[int64]*BlockAndValidator{}
	for _, key := range []string{totalSharesKey} {
		startV := getValue(start, client, key)
		endV := getValue(end, client, key)
		log.Info(startV, endV, key)
		check(start, startV, end, endV, client, key, m)
	}
	var arr []*BlockAndValidator
	for _, v := range m {
		arr = append(arr, v)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].Block < arr[j].Block })
	//
	file, err := os.Create("rebase_token.json")
	log.CheckFatal(err)
	file.Write(utils.ToJsonBytes(arr))
	file.Close()
}

func check(start int64, startV *big.Int, end int64, endV *big.Int, client core.ClientI, key string, m map[int64]*BlockAndValidator) {
	if end-start == 1 && startV.Cmp(endV) != 0 {
		if m[end] == nil {
			m[end] = &BlockAndValidator{}
		}
		log.Info(end, endV)
		if key == validatorsKey {
			m[end].Block = end
			m[end].Validator = endV.Int64()
		} else if key == totalSharesKey {
			m[end].Block = end
			m[end].TotalShares = (*core.BigInt)(endV)
		}
		return
	}
	if startV.Cmp(endV) == 0 {
		return
	}
	block := (start + end) / 2
	mid := getValue(block, client, key)
	if mid.Cmp(startV) != 0 && mid.Cmp(endV) != 0 {
		check(start, startV, block, mid, client, key, m)
		check(block, mid, end, endV, client, key, m)
	} else if mid.Cmp(startV) == 0 {
		check(block, startV, end, endV, client, key, m)
	} else {
		check(start, startV, block, endV, client, key, m)
	}
}

func getValue(block int64, client core.ClientI, key string) *big.Int {
	return GetStorageAt(block, key, client, "0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84")
}
