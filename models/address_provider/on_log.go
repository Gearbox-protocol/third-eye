package address_provider

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Gearbox-protocol/third-eye/models/account_factory"
	"github.com/Gearbox-protocol/third-eye/models/acl"
	"github.com/Gearbox-protocol/third-eye/models/contract_register"
	"github.com/Gearbox-protocol/third-eye/models/gear_token"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle"
	"github.com/Gearbox-protocol/third-eye/models/treasury"

	"fmt"
	"strings"
)

func (mdl *AddressProvider) OnLog(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("AddressSet(bytes32,address)"):
		contract := strings.Trim(string(txLog.Topics[1][:]), "\x00")
		address := common.HexToAddress(txLog.Topics[2].Hex()).Hex()
		blockNum := int64(txLog.BlockNumber)

		log.Infof("AddressSet: %s, %s at blockNum %d", contract, address, blockNum)
		switch contract {
		case "ACL":
			obj := acl.NewACL(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(obj)
		case "CONTRACTS_REGISTER":
			cr := contract_register.NewContractRegister(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(cr)
		case "PRICE_ORACLE":
			//
			mdl.addPriceOracle(blockNum, address)
			po := price_oracle.NewPriceOracle(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(po)
		case "ACCOUNT_FACTORY":
			af := account_factory.NewAccountFactory(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(af)
		case "WETH_TOKEN":
			if mdl.Details == nil {
				mdl.Details = make(map[string]interface{})
			}
			mdl.Repo.GetToken(address)
		case "GEAR_TOKEN":
			mdl.Repo.GetToken(address)
			gt := gear_token.NewGearToken(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
			mdl.Repo.AddSyncAdapter(gt)
		case "TREASURY_CONTRACT":
			// NOTE:don't disable the prev treasury as treasury addr is fixed for pool on pool creation
			//
			// addrs := mdl.Repo.GetAdapterAddressByName(ds.Treasury)
			// for _, addr := range addrs {
			// 	adapter := mdl.Repo.GetAdapter(addr)
			// 	if adapter.GetBlockToDisableOn() == math.MaxInt64 { // only disable the treasury adapters that aren't disabled before
			// 		adapter.SetBlockToDisableOn(blockNum)
			// 	}
			// }
			ttf := treasury.NewTreasury(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(ttf)
		case "DATA_COMPRESSOR":
			if mdl.Details == nil {
				mdl.Details = make(map[string]interface{})
			}
			dcObj, ok := mdl.Details["dc"].(map[string]interface{})
			log.Infof("Previous data compressors %#v\n", dcObj)
			if !ok {
				if dcObj == nil {
					dcObj = make(map[string]interface{})
				}
			}
			dcObj[fmt.Sprintf("%d", blockNum)] = address
			mdl.Details["dc"] = dcObj
			mdl.Repo.GetDCWrapper().AddDataCompressor(int64(txLog.BlockNumber), address)
		}
	}
}
