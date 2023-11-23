package cm_v2

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/configurators/configurator_v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CMv2) setv2AddrIfNotPresent() {
	if mdl.Details != nil && mdl.Details["facade"] != nil {
		return
	}
	opts := &bind.CallOpts{BlockNumber: big.NewInt(mdl.DiscoveredAt)}
	configuratorAddr, err := mdl.cmContractv2.CreditConfigurator(opts)
	log.CheckFatal(err)
	facadeAddr, err := mdl.cmContractv2.CreditFacade(opts)
	log.CheckFatal(err)
	if mdl.Details == nil {
		mdl.Details = core.Json{}
	}
	mdl.Details["facade"] = facadeAddr.Hex()
	mdl.Details["configurator"] = configuratorAddr.Hex()
}

type CMv2Fields struct {
	multicall   ds.MultiCallProcessor
	addrChanged bool
	//
	cmContractv2     *creditManagerv2.CreditManagerv2
	facadeContractv2 *creditFacade.CreditFacade
}

func (mdl *CMv2) setCreditFacadeSyncer(creditFacadeAddr string) {
	if mdl.Details == nil {
		mdl.Details = map[string]interface{}{}
	}

	oldFacade := mdl.GetDetailsByKey("facade")
	if oldFacade != "" && oldFacade == creditFacadeAddr {
		return
	}
	mdl.Details["facade"] = creditFacadeAddr
	mdl.addrChanged = true
}

func (mdl *CMv2) setConfiguratorSyncer(configuratorAddr string) {
	if mdl.Details == nil {
		mdl.Details = map[string]interface{}{}
	}
	oldconfigurator := mdl.GetDetailsByKey("configurator")
	if oldconfigurator != "" && oldconfigurator == configuratorAddr {
		return
	}
	mdl.Details["configurator"] = configuratorAddr
	mdl.addrChanged = true
}

func (mdl *CMv2) IsAddrChanged() bool {
	defer func() { mdl.addrChanged = false }()
	return mdl.addrChanged
}

func (mdl *CMv2) GetCreditFacadeAddr() string {
	return mdl.GetDetailsByKey("facade")
}

func (cm *CMv2) addCreditConfiguratorAdapter(creditConfigurator string) {
	cf := configurator_v2.NewConfiguratorv2(creditConfigurator, cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
}

func (cm CMv2) GetAllAddrsForLogs() []common.Address {
	addrs := []common.Address{common.HexToAddress(cm.GetAddress())}
	if addr := cm.GetDetailsByKey("facade"); addr != "" {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	if addr := cm.GetDetailsByKey("configurator"); addr != "" {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	return addrs
}
