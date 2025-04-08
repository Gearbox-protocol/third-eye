package main

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	client := ethclient.NewEthClient(cfg)
	data := []*ds.SyncAdapter{}
	err := db.Raw(`select * from sync_adapters where discovered_at>=21264499 and type='QueryPriceFeed' and details->>'pfType'='SingleAssetPF'`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		log.Info(entry.Address)
		underlyingfeeding, err := core.CallFuncGetSingleValue(client, "741bef1a", common.HexToAddress(entry.Address), 0, nil)
		log.CheckFatal(err)
		underlying := common.BytesToAddress(underlyingfeeding).Hex()
		entry.Details["underlyings"] = []string{underlying}

		tokenData := &schemas.TokenOracle{}
		err = db.Raw(`select * from token_oracle where feed=?`, underlying).Find(tokenData).Error
		log.CheckFatal(err)
		if tokenData.Feed == "" {
			log.Info("underlying", underlying)
		}
		err = db.Exec(`update sync_adapters set details=? where address=?`, entry.Details, entry.Address).Error
		log.CheckFatal(err)
	}
}
