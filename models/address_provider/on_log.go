package address_provider

import (
	"strconv"

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

		log.Infof("AddressSet: %s, %s", contract, address)
		switch contract {
		case "ACL":
			obj := acl.NewACL(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(obj)
		case "CONTRACTS_REGISTER":
			cr := contract_register.NewContractRegister(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(cr)
		case "PRICE_ORACLE":
			//
			if mdl.Details == nil {
				mdl.Details = make(map[string]interface{})
			}
			// price oracles
			priceOracles, ok := mdl.Details["priceOracles"].(map[string]interface{})
			if !ok {
				if priceOracles == nil {
					priceOracles = map[string]interface{}{}
				}
			}
			var lastVersion int64
			for versionStr := range priceOracles {
				version, err := strconv.ParseInt(versionStr, 10, 64)
				log.CheckFatal(err)
				lastVersion = version
			}
			priceOracles[fmt.Sprintf("%d", lastVersion+1)] = address
			mdl.Details["priceOracles"] = priceOracles
			//
			po := price_oracle.NewPriceOracle(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(po)
		case "ACCOUNT_FACTORY":
			af := account_factory.NewAccountFactory(address, blockNum, mdl.SyncAdapter.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(af)
		case "WETH_TOKEN":
			if mdl.Details == nil {
				mdl.Details = make(map[string]interface{})
			}
			mdl.Repo.SetWETHAddr(address)
			mdl.Repo.AddToken(address)
		case "GEAR_TOKEN":
			gt := gear_token.NewGearToken(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
			mdl.Repo.AddSyncAdapter(gt)
		case "TREASURY_CONTRACT":
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
			mdl.Repo.AddDataCompressor(int64(txLog.BlockNumber), address)
		}
	}
}
