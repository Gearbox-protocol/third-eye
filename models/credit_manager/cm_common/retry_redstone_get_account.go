package cm_common

import (
	"math/big"

	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/pkg/redstone"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CommonCMAdapter) priceFeedNeeded(balances core.DBBalanceFormat) (ans []redstone.TokenAndFeedType) {
	feeds := mdl.Repo.GetTokenOracles()[schemas.V3PF_MAIN]
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
		pfType := con.GetPFType()
		t := 0
		if pfType == ds.CompositeRedStonePF {
			t = core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE
		} else if pfType == ds.RedStonePF {
			t = core.V3_REDSTONE_ORACLE
		}
		{ // ignore LBTC price on mainnet as the pf0 of composite is not updated, so can't provide the pod
			client := mdl.Client
			chainId := core.GetChainId(client)
			addrToSym := core.GetTokenToSymbolByChainId(chainId)
			if addrToSym[common.HexToAddress(token)] == "LBTC" && log.GetBaseNet(chainId) == "MAINNET" {
				continue
			}
		}
		ans = append(ans, redstone.TokenAndFeedType{
			Token:    common.HexToAddress(token),
			Reversed: false,
			PFType:   t,
		})
	}
	return
}
func (mdl *CommonCMAdapter) retry(oldaccount dc.CreditAccountCallData, blockNum int64) (dc.CreditAccountCallData, error) {
	_, addr := mdl.Repo.GetDCWrapper().GetKeyAndAddress(core.NewVersion(300), blockNum)
	dcw, err := dcv3.NewDataCompressorv3(addr, mdl.Client)
	log.CheckFatal(err)
	ts := mdl.Repo.SetAndGetBlock(blockNum).Timestamp
	bal := moreThan1Balance(oldaccount.Balances)
	pod := mdl.Repo.GetRedStonemgr().GetPodSign(int64(ts), mdl.priceFeedNeeded(bal), bal)
	newaccountData, err := dcw.GetCreditAccountData(&bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	},
		oldaccount.Addr,
		pod,
	)
	if err != nil {
		return dc.CreditAccountCallData{}, err
	}
	if !newaccountData.IsSuccessful {
		log.Warn("after retry, getCreditAccoutn data is still not successful", blockNum, oldaccount.Addr)
	}
	return dc.GetAccountDataFromDCCall(mdl.Client, core.NULL_ADDR, blockNum, newaccountData)
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
