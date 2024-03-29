package address_provider

import (
	"math/big"

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

func (mdl *AddressProvider) commonLogParse(blockNum int64, contract string, address string) {
	switch contract {
	case "ACL":
		obj := acl.NewACL(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
		mdl.Repo.AddSyncAdapter(obj)
	case "CONTRACTS_REGISTER":
		cr := contract_register.NewContractRegister(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
		mdl.Repo.AddSyncAdapter(cr)
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
		// TODO logic for anvil fork, gear token has took many events to be able to sync
		if core.GetChainId(mdl.Client) == 7878 {
			return
		}
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
	}
}

func (mdl *AddressProvider) v2LogParse(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("AddressSet(bytes32,address)"):
		contract := strings.Trim(string(txLog.Topics[1][:]), "\x00")
		address := common.HexToAddress(txLog.Topics[2].Hex()).Hex()
		blockNum := int64(txLog.BlockNumber)

		log.Infof("AddressSet: %s, %s at blockNum %d", contract, address, blockNum)
		switch contract {
		case "PRICE_ORACLE":
			//
			mdl.addPriceOracle(blockNum, address)
			po := price_oracle.NewPriceOracle(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(po)
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
		default:
			mdl.commonLogParse(blockNum, contract, address)
		}
	}
}

func (mdl *AddressProvider) OnLog(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("AddressSet(bytes32,address)"):
		mdl.v2LogParse(txLog)
	case core.Topic("SetAddress(bytes32,address,uint256)"):
		mdl.v3LogParse(txLog)
	}
}

func (mdl *AddressProvider) v3LogParse(txLog types.Log) {
	contract := strings.Trim(string(txLog.Topics[1][:]), "\x00")
	address := common.HexToAddress(txLog.Topics[2].Hex()).Hex()
	version := func() int16 {
		version :=
			new(big.Int).SetBytes(txLog.Topics[3].Bytes()).Int64()
		if version == 0 {
			version = 1
		}
		return int16(version)
	}()
	blockNum := int64(txLog.BlockNumber)
	//
	log.Infof("AddressSet: %s(%d), %s at blockNum %d", contract, version, address, blockNum)
	switch contract {
	case "DATA_COMPRESSOR":
		//
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
		if version < 300 { // don't add dataCompressor with version 2.1
			log.Infof("Don't add %s version %d", address, version)
			return
		}
		dcObj[fmt.Sprintf("%d", blockNum)] = fmt.Sprintf("%s_%d", address, version)
		mdl.Details["dc"] = dcObj
		// v3
		mdl.Repo.GetDCWrapper().AddDataCompressorByVersion(core.NewVersion(version), address, blockNum)
	case "PRICE_ORACLE":
		if version < 300 { // don't except v2,v2.10 or v1 priceOracle , why are already know from v1 addressProvider
			return
		}
		mdl.addPriceOracle(blockNum, address)
		po := price_oracle.NewPriceOracle(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
		mdl.Repo.AddSyncAdapter(po)
	default:
		mdl.commonLogParse(blockNum, contract, address)
	}
}
