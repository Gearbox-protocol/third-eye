package cm_v1

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/configurators/credit_filter"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_mvp"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type CMv1 struct {
	// supplmentary contracts
	contractETHV1 *creditManager.CreditManager
	// tmp storage for find/checking actual adpater operations
	executeParams []ds.ExecuteParams

	//
	cm_mvp.CmMVP
}

func NewCMv1(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *CMv1 {
	// credit manager
	mdl := NewCMv1FromAdapter(
		ds.NewSyncAdapter(addr, ds.CreditManager, discoveredAt, client, repo),
	)
	mdl.CommonInitState(mdl.GetVersion())
	mdl.addCreditFilterAdapter(discoveredAt)
	return mdl
}

func (cm *CMv1) addCreditFilterAdapter(blockNum int64) {
	creditFilter, err := cm.contractETHV1.CreditFilter(&bind.CallOpts{BlockNumber: big.NewInt(blockNum)})
	if err != nil {
		log.Fatal(err)
	}
	cm.Repo.GetDCWrapper().AddCreditManagerToFilter(cm.Address, creditFilter.Hex())
	cf := credit_filter.NewCreditFilter(creditFilter.Hex(), cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
}

func NewCMv1FromAdapter(adapter *ds.SyncAdapter) *CMv1 {
	//
	obj := &CMv1{
		CmMVP: cm_mvp.NewCMCommon(adapter),
		//
	}
	//
	obj.GetAbi()
	obj.SetOnChange()
	// v1 separate logic
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj.contractETHV1 = cmContract
	// v1 separate logic
	return obj
}

func (mdl *CMv1) IsAddrChanged() bool {
	return false
}
