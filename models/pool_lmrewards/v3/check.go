package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *LMRewardsv3) check(blockNum int64, currentTs uint64, farm, from, to string) {
	// 4216f972
	if blockNum > 18853342 || !utils.Contains([]int64{1, 7878}, core.GetChainId(mdl.Client)) {
		return
	}
	for _, addr := range []string{from, to} {
		if addr != core.NULL_ADDR.Hex() {
			s := common.HexToHash(addr)
			farmedData, err := core.CallFuncWithExtraBytes(mdl.Client, "4216f972", common.HexToAddress(farm), blockNum, s[:])
			log.CheckFatal(err)
			expected := new(big.Int).SetBytes(farmedData)
			obj := mdl.users[common.HexToAddress(farm)][addr]
			actual := obj.GetPoints(mdl.farms[farm], currentTs)
			if expected.Cmp(actual) != 0 {
				log.Fatal("Expected", expected, "Actual", actual, "farm", utils.ToJson(mdl.farms[farm]), "user", utils.ToJson(obj))
			}
		}
	}
}
