package ds

import (
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

// for use in chainlinkPriceFeed and compositePriceFeed

type entry struct {
	BlockNumber     int64                   `json:"blockNum"`
	MergedPFVersion schemas.MergedPFVersion `json:"mergedPFVersion"`
	Token           string                  `json:"-"`
}
type MergedPFManager []entry

func (mdl *MergedPFManager) add(v int64, details core.Json, discoveredAt int64) {
	*mdl = append(*mdl, entry{
		MergedPFVersion: schemas.MergedPFVersion(v),
		BlockNumber:     discoveredAt,
		Token:           details["token"].(string),
	})

	delete(details, "mergedPFVersion")
}
func (mdl *MergedPFManager) Load(details core.Json, discoveredAt int64) {
	if details == nil {
		return
	}
	if details["mergedPFVersion"] != nil {

		switch v := details["mergedPFVersion"].(type) {
		case float64:
			mdl.add(int64(v), details, discoveredAt)
			return
		case int64:
			mdl.add(int64(v), details, discoveredAt)
			return
		case schemas.MergedPFVersion:
			mdl.add(int64(v), details, discoveredAt)
			return
		case []interface{}:
			for _, val := range v {
				x := val.(map[string]interface{})
				*mdl = append(*mdl, entry{
					MergedPFVersion: schemas.MergedPFVersion(x["mergedPFVersion"].(float64)),
					BlockNumber:     int64(x["blockNum"].(float64)),
					Token:           details["token"].(string),
				})
			}
		default:
			log.Fatal("can't get mergedPFVersion", details, reflect.TypeOf(details["mergedPFVersion"]))
		}

	}
	// if it is nil
}

func (mdl MergedPFManager) Save(details *core.Json) {
	(*details)["mergedPFVersion"] = mdl
	// log.Info(utils.ToJson((*details)["mergedPFVersion"]))
}

func (mdl MergedPFManager) GetMergedPFVersion(blockNum int64, syncAdapterAddr string) schemas.MergedPFVersion {
	for ind := len(mdl) - 1; ind >= 0; ind-- {
		if mdl[ind].BlockNumber <= blockNum {
			return mdl[ind].MergedPFVersion
		}
	}
	log.Fatal("Can't get mergedPFVersion", mdl, blockNum, syncAdapterAddr)
	return schemas.MergedPFVersion(0)
}
func (mdl *MergedPFManager) AddToken(token string, blockNum int64, pfVersion schemas.PFVersion) {
	var last schemas.MergedPFVersion
	if len(*mdl) != 0 {
		obj := (*mdl)[len(*mdl)-1]
		if obj.Token != token {
			log.Fatal("stored token for chainlink is different from new added token", obj.Token, token)
		}
		last = obj.MergedPFVersion
	}
	*mdl = append(*mdl, entry{
		Token:           token,
		MergedPFVersion: schemas.MergedPFVersion(pfVersion) | last,
		BlockNumber:     blockNum,
	})
}

func (mdl *MergedPFManager) DisableToken(token string, blockNum int64, pfVersion schemas.PFVersion) {
	var last schemas.MergedPFVersion
	if len(*mdl) != 0 {
		obj := (*mdl)[len(*mdl)-1]
		if obj.Token != token {
			log.Fatal("stored token for chainlink is different from new added token", obj.Token, token)
		}
		last = obj.MergedPFVersion
	}
	final := last ^ schemas.MergedPFVersion(pfVersion)
	*mdl = append(*mdl, entry{
		Token:           token,
		MergedPFVersion: final,
		BlockNumber:     blockNum,
	})
}
