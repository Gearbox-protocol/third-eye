package helper

import (
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"gorm.io/gorm"
)

func GetDC(client core.ClientI, db *gorm.DB) *dc_wrapper.DataCompressorWrapper {
	dc := dc_wrapper.NewDataCompressorWrapper(client)

	{ // add dc addrs
		s := &ds.SyncAdapter{}
		err := db.Raw(`select details from sync_adapters where type='AddressProvider'`).Find(s).Error
		log.CheckFatal(err)
		for block, dcAddr := range s.Details["dc"].(map[string]interface{}) {
			i, err := strconv.ParseInt(block, 10, 64)
			log.CheckFatal(err)
			splits := strings.Split(dcAddr.(string), "_")
			if len(splits) == 2 {
				dc.AddDataCompressorByVersion(core.NewVersion(300), splits[0], i)
			}
		}
	}
	return dc
}
