package rebase_token

import (
	"encoding/hex"
	"math/big"
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *RebaseToken) previousDetails() stETHValues {
	totalETH := mdl.getBigIntFromDetails("totalETH")
	totalShares := mdl.getBigIntFromDetails("totalShares")
	blockNum := func() int64 {
		i, ok := mdl.Details["blockNum"]
		if !ok {
			return 0
		}
		switch i.(type) {
		case float64:
			return int64(i.(float64))
		case int64:
			return int64(i.(int64))
		default:
			log.Fatal("", reflect.TypeOf(i))
			return 0
		}
	}()
	if blockNum == 0 {
		return mdl.GetstETHDetails(mdl.DiscoveredAt)
	}
	return stETHValues{
		totalETH:    totalETH,
		totalShares: totalShares,
		blockNum:    blockNum,
	}
}

type stETHValues struct {
	totalETH    *big.Int
	totalShares *big.Int
	blockNum    int64
}

func (mdl stETHValues) ratio() *big.Int {
	if mdl.totalShares.Cmp(new(big.Int)) == 0 {
		return new(big.Int)
	}
	return new(big.Int).Quo(new(big.Int).Mul(mdl.totalETH, utils.GetExpInt(8)), mdl.totalShares)
}

func (mdl *RebaseToken) GetstETHDetails(blockNum int64) stETHValues {
	totalSharesData, err := hex.DecodeString("d5002f2e") // total shares
	log.CheckFatal(err)
	totalETHData, err := hex.DecodeString("37cfdaca") // total pooled eth
	log.CheckFatal(err)
	token := common.HexToAddress(mdl.GetAddress())
	calls := []multicall.Multicall2Call{
		{
			Target:   token,
			CallData: totalSharesData,
		},
		{
			Target:   token,
			CallData: totalETHData,
		},
	}
	results := core.MakeMultiCall(mdl.Client, blockNum, false, calls)
	totalShares, _ := core.MulticallAnsBigInt(results[0])
	totalETH, _ := core.MulticallAnsBigInt(results[1])
	return stETHValues{
		totalETH:    totalETH,
		totalShares: totalShares,
		blockNum:    blockNum,
	}
}

func (mdl stETHValues) ToDB() *schemas.RebaseDetailsForDB {
	return &schemas.RebaseDetailsForDB{
		TotalShares: (*core.BigInt)(mdl.totalShares),
		TotalETH:    (*core.BigInt)(mdl.totalETH),
		BlockNum:    mdl.blockNum,
	}
}

func (mdl RebaseToken) save(value stETHValues) {
	mdl.Details["totalETH"] = value.totalETH.String()
	mdl.Details["totalShares"] = value.totalShares.String()
	mdl.Details["blockNum"] = value.blockNum
}
