package credit_manager

import (
	"math/big"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditManager) setv2AddrIfNotPresent() {
	if mdl.Details != nil && mdl.Details["facade"] != nil {
		return
	}
	opts := &bind.CallOpts{BlockNumber: big.NewInt(mdl.DiscoveredAt)}
	configuratorAddr, err := mdl.contractETHV2.CreditConfigurator(opts)
	log.CheckFatal(err)
	facadeAddr, err := mdl.contractETHV2.CreditFacade(opts)
	log.CheckFatal(err)
	if mdl.Details == nil {
		mdl.Details = core.Json{}
	}
	mdl.Details["facade"] = facadeAddr.Hex()
	mdl.Details["configurator"] = configuratorAddr.Hex()
}

type CMv2Fields struct {
	configuratorSyncer *SubsidiarySyncer
	facadeSyncer       *SubsidiarySyncer
	//
	multicall MultiCallProcessor
	//
	contractETHV2    *creditManagerv2.CreditManagerv2
	facadeContractV2 *creditFacade.CreditFacade
}

func (mdl *CreditManager) setCreditFacadeSyncer(creditFacadeAddr string, syncFrom int64) {
	if mdl.Details == nil {
		mdl.Details = map[string]interface{}{}
	}
	mdl.Details["facade"] = creditFacadeAddr
	if mdl.facadeSyncer != nil && mdl.facadeSyncer.Address.Hex() == creditFacadeAddr {
		return
	}
	mdl.facadeSyncer = NewSubsidiarySyncer(mdl.Client, creditFacadeAddr, nil)
	if syncFrom != 0 {
		mdl.facadeSyncer.FetchLogs(syncFrom, mdl.WillSyncTill)
	}
}

func (mdl *CreditManager) setConfiguratorSyncer(configuratorAddr string, syncFrom int64) {
	if mdl.Details == nil {
		mdl.Details = map[string]interface{}{}
	}
	mdl.Details["configurator"] = configuratorAddr
	if mdl.configuratorSyncer != nil && mdl.configuratorSyncer.Address.Hex() == configuratorAddr {
		return
	}
	mdl.configuratorSyncer = NewSubsidiarySyncer(mdl.Client, configuratorAddr, [][]common.Hash{
		{
			core.Topic("CreditFacadeUpgraded(address)"),
			core.Topic("FeesUpdated(uint16,uint16,uint16,uint16,uint16)"),
		},
	})
	if syncFrom != 0 {
		mdl.configuratorSyncer.FetchLogs(syncFrom, mdl.WillSyncTill)
	}
}

func (mdl *CreditManager) WillBeSyncedTo(blockNum int64) {
	if mdl.configuratorSyncer != nil {
		mdl.configuratorSyncer.FetchLogs(mdl.LastSync+1, blockNum)
	}
	if mdl.facadeSyncer != nil {
		mdl.facadeSyncer.FetchLogs(mdl.LastSync+1, blockNum)
	}
	mdl.SyncAdapter.WillBeSyncedTo(blockNum)
}

func (mdl CreditManager) getv2ExtraLogs(txLog types.Log) (ans []types.Log) {
	if mdl.configuratorSyncer != nil {
		ans = append(ans, mdl.configuratorSyncer.GetLogsBefore(txLog)...)
	}
	if mdl.facadeSyncer != nil {
		ans = append(ans, mdl.facadeSyncer.GetLogsBefore(txLog)...)
	}
	sort.SliceStable(ans, func(i, j int) bool {
		return ans[i].BlockNumber < ans[j].BlockNumber ||
			(ans[i].BlockNumber == ans[j].BlockNumber && ans[i].Index < ans[j].Index)
	})
	return
}

//
//

func (mdl *CreditManager) GetCreditFacadeAddr() string {
	return mdl.GetDetailsByKey("facade")
}

func (cm *CreditManager) addCreditConfiguratorAdapter(creditConfigurator string) {
	// this is need for mask only
	cf := credit_filter.NewCreditFilter(creditConfigurator, ds.CreditConfigurator, cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
}