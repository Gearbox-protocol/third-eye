package price_feed

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/Gearbox-protocol/gearscan/core"
	"math/big"
	"github.com/Gearbox-protocol/gearscan/log"
	"fmt"
	// "github.com/Gearbox-protocol/gearscan/models/price_feed"
)

func GetExpFloat(decimals int64) *big.Float {
	if decimals < 0 {
		panic(fmt.Sprintf("GetExpFloat received pow:%d", decimals))
	}
	bigIntDecimal := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(decimals), big.NewInt(0))
	return new(big.Float).SetInt(bigIntDecimal)
}
func IntToFloat(amt *big.Int) *big.Float {
	return new(big.Float).SetInt(amt)
}

func (mdl *PriceFeed) OnLog(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("AnswerUpdated(int256,uint256,uint256)"):
		roundId, ok:=new(big.Int).SetString(txLog.Topics[2].Hex()[2:], 16)
		if ok {
			log.Fatal("roundid failed")
		}
		answerBI, err:=new(big.Int).SetString(txLog.Topics[1].Hex()[2:], 16)
		if err {
			log.Fatal("answer parsing failed")
		}
		answer, _ := new(big.Float).Quo(
			IntToFloat(answerBI),
			GetExpFloat(18),
		).Float64()
		// new(big.Int).SetString(txLog.Data[2:], 16)
		blockNum:=int64(txLog.BlockNumber)
		mdl.Repo.AddPriceFeed(blockNum, &core.PriceFeed{
			BlockNumber: blockNum,
			Token: "",
			Feed: mdl.Address,
			RoundId: roundId,
			PriceETHBI: answerBI,
			PriceETH: answer,
		})
	}
}