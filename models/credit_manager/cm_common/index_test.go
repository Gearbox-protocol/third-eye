package cm_common

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type RepoWrapper struct {
	*ds.DummyRepo
	prices core.JsonFloatMap
	tokens map[string]*schemas.Token
}

func (repo RepoWrapper) GetToken(addr string) *schemas.Token {
	return repo.tokens[addr]
}
func (repo RepoWrapper) GetPricesInUSD(blockNum int64, tokenAddrs []string) core.JsonFloatMap {
	return repo.prices
}
func TestGetCollateralAmountOnOpen(t *testing.T) {
	usdc, weth := utils.RandomAddr(), utils.RandomAddr()
	repo := RepoWrapper{
		DummyRepo: &ds.DummyRepo{},
		prices: core.JsonFloatMap{
			weth: 1800,
			usdc: 1,
		},
		tokens: map[string]*schemas.Token{
			weth: {Symbol: "WETH", Decimals: 18},
			usdc: {Symbol: "USDC", Decimals: 6},
		},
	}
	common := NewCommonCMAdapter(&ds.SyncAdapter{Repo: repo})
	common.State = &schemas.CreditManagerState{
		UnderlyingToken: weth,
	}
	// account has weth as underlying
	collateral := common.GetCollateralAmount(5, &schemas.AccountOperation{
		MultiCall: []*schemas.AccountOperation{
			{
				Action: "AddCollateral(address,address,uint256)",
				Transfers: &core.Transfers{
					usdc: utils.GetExpInt(6 + 3), // 1000 usdc
				},
			},
			{
				Action: "AddCollateral(address,address,uint256)",
				Transfers: &core.Transfers{
					weth: utils.GetExpInt(18), // 1 weth
				},
			},
		},
	})
	if collateral.String() != "1555555555555555555" { // 1+ 1000/1800
		t.Fatalf("Collateral %d is different.", collateral)
	}
}
