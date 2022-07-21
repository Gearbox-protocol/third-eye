package credit_filter

import (
	"math/rand"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfigurator"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/utils"
)

type CreditFilter struct {
	*ds.SyncAdapter
	filterContract *creditFilter.CreditFilter
	//
	cfgContract     *creditConfigurator.CreditConfigurator
	underlyingToken *common.Address
}

func NewCreditFilter(addr, contractName, creditManager string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *CreditFilter {
	syncAdapter := ds.NewSyncAdapter(addr, contractName, discoveredAt, client, repo)
	syncAdapter.Details = map[string]interface{}{"creditManager": creditManager}
	mdl := NewCreditFilterFromAdapter(
		syncAdapter,
	)
	// 	DEFAULT_FEE_INTEREST, 1000
	// DEFAULT_FEE_LIQUIDATION 200,
	// PERCENTAGE_FACTOR = 10000 - DEFAULT_LIQUIDATION_PREMIUM = 500
	if mdl.GetVersion() == 2 {
		mdl.addFees(uint(rand.Intn(50000)+50000), discoveredAt, common.Hash{}.Hex(),
			1000, 200, 9500)
	}
	return mdl
}

func NewCreditFilterFromAdapter(adapter *ds.SyncAdapter) *CreditFilter {
	cfContract, err := creditFilter.NewCreditFilter(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &CreditFilter{
		SyncAdapter:    adapter,
		filterContract: cfContract,
	}
	if adapter.ContractName == ds.CreditConfigurator {
		cfgContract, err := creditConfigurator.NewCreditConfigurator(common.HexToAddress(adapter.Address), adapter.Client)
		log.CheckFatal(err)
		obj.cfgContract = cfgContract
		if utils.Contains([]string{"0x3c81C60c2Ca83b7FE5279a10138ea9243f511edb", "0x6144112c74D9a1392b31bd8Ec6d75A0aCa616D1a", "0xbFcbF17b77cc82A685a925A79b3B6E0972e23624"}, obj.Address) {
			underlyingToken, err := cfgContract.Underlying(nil)
			obj.underlyingToken = &underlyingToken
			log.CheckFatal(err)
		}
	}
	return obj
}
