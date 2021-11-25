package repository

import (
	"github.com/Gearbox-protocol/gearscan/ethclient" 
	"github.com/Gearbox-protocol/gearscan/models/address_provider" 
	"github.com/Gearbox-protocol/gearscan/core" 
	"github.com/Gearbox-protocol/gearscan/models/acl" 
	"github.com/Gearbox-protocol/gearscan/models/credit_manager" 
	"github.com/Gearbox-protocol/gearscan/models/price_oracle" 
	"github.com/Gearbox-protocol/gearscan/models/data_compressor" 
	"github.com/Gearbox-protocol/gearscan/models/contract_register" 
	"github.com/Gearbox-protocol/gearscan/models/account_factory" 
	"github.com/Gearbox-protocol/gearscan/models/pool" 
)

func prepareSyncAdapter(adapter *core.SyncAdapter, repo core.RepositoryI, client *ethclient.Client) core.SyncAdapterI {
	adapter.Client = client
	switch adapter.Type {
		case "ACL":
			return &acl.ACL{SyncAdapter: adapter, State: &core.State{Repo: repo}}
		case "AddressProvider":
			return &address_provider.AddressProvider{SyncAdapter: adapter, State: &core.State{Repo: repo}}
		case "AccountFactory":
			return &account_factory.AccountFactory{SyncAdapter: adapter, State: &core.State{Repo: repo}}
		case "Pool":
			return &pool.Pool{SyncAdapter: adapter, State: &core.State{Repo: repo}}
		case "CreditManager":
			return &credit_manager.CreditManager{SyncAdapter: adapter, State: &core.State{Repo: repo}}
		case "PriceOracle":
			return &price_oracle.PriceOracle{SyncAdapter: adapter, State: &core.State{Repo: repo}}
		case "DataCompressor":
			return &data_compressor.DataCompressor{SyncAdapter: adapter, State: &core.State{Repo: repo}}
		case "ContractRegister":
			return &contract_register.ContractRegister{SyncAdapter: adapter, State: &core.State{Repo: repo}}
	}
	return nil
}