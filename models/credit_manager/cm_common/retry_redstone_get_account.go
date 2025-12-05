package cm_common

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/pkg/redstone"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CommonCMAdapter) priceFeedNeededRedStone(balances core.DBBalanceFormat) (ans []core.RedStonePF) {
	pool := mdl.State.PoolAddress
	priceOracle := mdl.Repo.GetAdapter(pool).(*pool_v3.Poolv3).State.PriceOracle
	feeds := mdl.Repo.GetMainTokenOracles()[priceOracle]
	for token := range balances {
		var con ds.QueryPriceFeedI
		{
			d := feeds[token]
			adapter := mdl.Repo.GetAdapter(d.Feed)
			if adapter.GetName() != ds.QueryPriceFeed {
				continue
			}
			con = adapter.(ds.QueryPriceFeedI)
		}
		// pfType := con.GetPFType()
		// if pfType == ds.CompositeRedStonePF || pfType == ds.RedStonePF || pfType == ds.SingleAssetPF || pfType == ds.Cu {
		reds := con.GetRedstonePF()
		for _, red := range reds {
			ans = append(ans, *red)
		}
		// }
	}
	return
}

func (mdl *CommonCMAdapter) priceFeedPyth(balances core.DBBalanceFormat, blockNum int64) (ans []redstone.PytDataWithFeed) {
	pool := mdl.State.PoolAddress
	priceOracle := mdl.Repo.GetAdapter(pool).(*pool_v3.Poolv3).State.PriceOracle
	feeds := mdl.Repo.GetMainTokenOracles()[priceOracle]
	ts := mdl.Repo.SetAndGetBlock(blockNum).Timestamp
	for token, bal := range balances {
		{
			d := feeds[token]
			adapter := mdl.Repo.GetAdapter(d.Feed)
			if adapter.GetName() == ds.QueryPriceFeed && d.FeedType == ds.PythPF && bal.BI.Convert().Cmp(big.NewInt(10)) > 0 {
				datas, ok := adapter.GetDetails()["underlyings"].([]interface{})
				if !ok {
					log.Fatal("Pyth price feed underlyings not found for token:", token, "adapter:", adapter.GetAddress())
				}
				data, err := pkg.GetPythPrice(datas[0].(string), int64(ts))
				log.CheckFatal(err)
				ans = append(ans, redstone.PytDataWithFeed{
					PythData: data,
					Feed:     common.HexToAddress(adapter.GetAddress()),
				})
			}
		}
	}
	return
}
func (mdl *CommonCMAdapter) retry(oldaccount dc.CreditAccountCallData, blockNum int64) (dc.CreditAccountCallData, error) {
	ts := mdl.Repo.SetAndGetBlock(blockNum).Timestamp
	bal := moreThan1Balance(oldaccount.Balances)
	bal[mdl.GetUnderlyingToken()] = core.DBTokenBalance{BI: (*core.BigInt)(big.NewInt(1))}
	redPFs := mdl.priceFeedNeededRedStone(bal)
	pythPFs := mdl.priceFeedPyth(bal, blockNum)
	v3Pods := mdl.Repo.GetRedStonemgr().GetPodSign(int64(ts), redPFs)
	log.Info(mdl.Address)
	x := core.FetchActualVersion(mdl.Address, 0, mdl.Client)
	if x < 310 {
		x = 300
	}
	v3PodsCalls := redstone.GetpodToCalls(mdl.Client, x, common.HexToAddress(mdl.GetCreditFacadeAddr()), v3Pods, redPFs, pythPFs)
	log.Info("retrying to get credit account data", oldaccount.Addr, blockNum, "pods", len(v3Pods), "calls", len(v3PodsCalls))
	//
	return mdl.Repo.GetDCWrapper().Retry(blockNum, oldaccount.Addr, v3Pods, v3PodsCalls)
}

func moreThan1Balance(oldBal []core.TokenBalanceCallData) core.DBBalanceFormat {
	dbFormat := core.DBBalanceFormat{}
	for _, balance := range oldBal {
		token := balance.Token
		if balance.HasBalanceMoreThanOne() && balance.IsEnabled {
			dbFormat[token] = balance.DBTokenBalance
		}
	}
	return dbFormat
}
