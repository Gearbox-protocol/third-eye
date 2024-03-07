package ds

import (
	"fmt"
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

// for use in chainlinkPriceFeed and compositePriceFeed

type entry struct {
	BlockNumber     int64                   `json:"blockNum"`
	MergedPFVersion schemas.MergedPFVersion `json:"mergedPFVersion"`
}
type MergedPFManager map[string][]entry

func (mdl MergedPFManager) add(v int64, details core.Json, discoveredAt int64) {
	token := details["token"].(string)
	mdl[token] = append(mdl[token], entry{
		MergedPFVersion: schemas.MergedPFVersion(v),
		BlockNumber:     discoveredAt,
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
			var tokens []string
			switch _tokens := details["token"].(type) {
			case string:
				tokens = []string{_tokens}
			case []string:
				tokens = _tokens
			case []interface{}:
				for _, x := range _tokens {
					tokens = append(tokens, x.(string))
				}
			default:
				log.Fatal("can't get mergedPFVersion", details, reflect.TypeOf(details["token"]))
			}
			for _, det := range v {
				snapDetails := det.(map[string]interface{})
				for _, token := range tokens {
					(*mdl)[token] = append((*mdl)[token], entry{
						MergedPFVersion: schemas.MergedPFVersion(snapDetails["mergedPFVersion"].(float64)),
						BlockNumber:     int64(snapDetails["blockNum"].(float64)),
					})
				}
			}
		case map[string]interface{}:
			for token, det := range v {
				snaps := det.([]interface{})
				for _, snap := range snaps {
					snapDetails := snap.(map[string]interface{})
					(*mdl)[token] = append((*mdl)[token], entry{
						MergedPFVersion: schemas.MergedPFVersion(snapDetails["mergedPFVersion"].(float64)),
						BlockNumber:     int64(snapDetails["blockNum"].(float64)),
					})
				}
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

func (mdl MergedPFManager) GetMergedPFVersion(token string, blockNum int64, syncAdapterAddr string) schemas.MergedPFVersion {
	version, err := mdl.getMergedPFVersion(token, blockNum, syncAdapterAddr)
	log.CheckFatal(err)
	return version
}

func (mdl MergedPFManager) getMergedPFVersion(token string, blockNum int64, syncAdapterAddr string) (schemas.MergedPFVersion, error) {
	for ind := len(mdl[token]) - 1; ind >= 0; ind-- {
		if mdl[token][ind].BlockNumber <= blockNum {
			return mdl[token][ind].MergedPFVersion, nil
		}
	}
	return schemas.MergedPFVersion(0), fmt.Errorf("can't get mergedPFVersion %v at %d for adapter: %s", mdl, blockNum, syncAdapterAddr)
}

func (mdl MergedPFManager) AddToken(token string, blockNum int64, pfVersion schemas.PFVersion) {
	var last schemas.MergedPFVersion
	if len(mdl[token]) != 0 {
		obj := mdl[token][len(mdl[token])-1]
		last = obj.MergedPFVersion
	}
	mdl[token] = append(mdl[token], entry{
		MergedPFVersion: schemas.MergedPFVersion(pfVersion) | last,
		BlockNumber:     blockNum,
	})
}
func (mdl MergedPFManager) GetTokens(blockNum int64) (tokens []string) {
	for token := range mdl {
		version, err := mdl.getMergedPFVersion(token, blockNum, "")
		if version != 0 && err == nil {
			tokens = append(tokens, token)
		}
	}
	return
}

func (mdl MergedPFManager) DisableToken(blockNum int64, token string, pfVersion schemas.PFVersion) {
	var last schemas.MergedPFVersion
	if len(mdl[token]) != 0 {
		obj := mdl[token][len(mdl[token])-1]
		last = obj.MergedPFVersion
	}
	final := last ^ schemas.MergedPFVersion(pfVersion)
	mdl[token] = append(mdl[token], entry{
		MergedPFVersion: final,
		BlockNumber:     blockNum,
	})
}
