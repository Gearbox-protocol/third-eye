package cm_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacadev3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/models/configurators/configurator_v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CMv3) setv3AddrIfNotPresent() {
	// set credit manager
	cmContract, err := creditManagerv3.NewCreditManagerv3(common.HexToAddress(mdl.Address), mdl.Client)
	log.CheckFatal(err)
	mdl.cmContractv3 = cmContract
	// set credit facade contract
	mdl.facadeContractv3, err = creditFacadev3.NewCreditFacadev3(core.NULL_ADDR, nil)
	log.CheckFatal(err)
	//
	if mdl.Details != nil && mdl.Details["facade"] != nil {
		return
	}
	opts := &bind.CallOpts{BlockNumber: big.NewInt(mdl.DiscoveredAt)}
	configuratorAddr, err := mdl.cmContractv3.CreditConfigurator(opts)
	log.CheckFatal(err)
	facadeAddr, err := mdl.cmContractv3.CreditFacade(opts)
	log.CheckFatal(err)
	if mdl.Details == nil {
		mdl.Details = core.Json{}
	}
	mdl.Details["facade"] = facadeAddr.Hex()
	mdl.Details["configurator"] = configuratorAddr.Hex()
}

type CMv3Fields struct {
	addrChanged bool
	//
	cmContractv3     *creditManagerv3.CreditManagerv3
	facadeContractv3 *creditFacadev3.CreditFacadev3
}

func (mdl *CMv3) setCreditFacadeSyncer(creditFacadeAddr string) {
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

func (mdl *CMv3) setConfiguratorSyncer(configuratorAddr string) {
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

func (mdl CMv3Fields) IsAddrChanged() bool {
	defer func() { mdl.addrChanged = false }()
	return mdl.addrChanged
}

func (cm CMv3) addCreditConfiguratorAdapter(creditConfigurator string) {
	cf := configurator_v3.NewConfiguratorv3(creditConfigurator, cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
}

func (cm CMv3) GetAllAddrsForLogs() []common.Address {
	addrs := []common.Address{common.HexToAddress(cm.GetAddress())}
	if addr := cm.GetDetailsByKey("facade"); addr != "" {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	if addr := cm.GetDetailsByKey("configurator"); addr != "" {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	return addrs
}
