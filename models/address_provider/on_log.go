package address_provider

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Gearbox-protocol/gearscan/models/account_factory"
	"github.com/Gearbox-protocol/gearscan/models/acl"
	"github.com/Gearbox-protocol/gearscan/models/contract_register"
	"github.com/Gearbox-protocol/gearscan/models/price_oracle"
	// "github.com/Gearbox-protocol/gearscan/models/data_compressor"

	"strings"
)

func (mdl *AddressProvider) OnLog(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("AddressSet(bytes32,address)"):
		contract := strings.Trim(string(txLog.Topics[1][:]), "\x00")
		address := common.HexToAddress(txLog.Topics[2].Hex()).Hex()
		blockNum := int64(txLog.BlockNumber)

		log.Infof("AddressSet: %s, %s", contract, address)
		switch contract {
		case "ACL":
			obj := acl.NewACL(address, blockNum, mdl.SyncAdapter.Client, mdl.State.Repo)
			mdl.State.Repo.AddSyncAdapter(obj)
		case "CONTRACTS_REGISTER":
			cr := contract_register.NewContractRegister(address, blockNum, mdl.SyncAdapter.Client, mdl.State.Repo)
			mdl.State.Repo.AddSyncAdapter(cr)
		case "PRICE_ORACLE":
			po := price_oracle.NewPriceOracle(address, blockNum, mdl.SyncAdapter.Client, mdl.State.Repo)
			mdl.State.Repo.AddSyncAdapter(po)
		case "ACCOUNT_FACTORY":
			af := account_factory.NewAccountFactory(address, blockNum, mdl.SyncAdapter.Client, mdl.State.Repo)
			mdl.State.Repo.AddSyncAdapter(af)
			// case "DATA_COMPRESSOR":
			// 	dc := data_compressor.NewDataCompressor(address, mdl.SyncAdapter.Client, mdl.State.Repo, blockNum)
			// 	mdl.State.Repo.AddSyncAdapter(dc)
		}
	}
}
