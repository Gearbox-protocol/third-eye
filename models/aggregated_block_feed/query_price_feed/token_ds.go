package query_price_feed

import (
	"bytes"
	"encoding/json"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type DetailsDS struct {
	PFType string                                   `json:"pfType"`
	Tokens map[string]map[schemas.PFVersion][]int64 `json:"tokens"` // token enabled and disabled at block numbers
	Logs   [][]interface{}                          `json:"logs"`
	// redunantt
	Reduntant       map[string][]int64       `json:"token,omitempty"` // token enabled and disabled at block numbers
	MergedPFVersion *schemas.MergedPFVersion `json:"mergedPFVersion,omitempty"`
}

//	func NewDetailsDS(pfType string, token string, discoveredAt int64, pfVersion schemas.PFVersion) *DetailsDS {
//		return &DetailsDS{
//			PFType:    pfType,
//			Tokens:    map[string][]int64{token: {discoveredAt}},
//			PFVersion: map[string]schemas.PFVersion{token: pfVersion},
//		}
//	}
func (obj *DetailsDS) Load(in core.Json, version core.VersionType) {
	data, err := json.Marshal(in)
	log.CheckFatal(err)
	err = utils.ReadJsonReaderAndSetInterface(bytes.NewBuffer(data), obj)
	log.CheckFatal(err)

	if obj.Tokens == nil {
		obj.Tokens = map[string]map[schemas.PFVersion][]int64{}
	}
	if len(obj.Reduntant) != 0 && (obj.MergedPFVersion == nil || *obj.MergedPFVersion == 0) {
		log.Fatal("For reduntant, mergedPFVersion should be set", utils.ToJson(obj.Reduntant), obj.MergedPFVersion)
	}
	for token, blockNums := range obj.Reduntant {
		if obj.Tokens[token] == nil {
			obj.Tokens[token] = map[schemas.PFVersion][]int64{}
		}
		if version == core.NewVersion(1) {
			obj.Tokens[token][schemas.V1PF] = blockNums
		} else if version == core.NewVersion(300) {
			obj.Tokens[token][schemas.V3PF_MAIN] = blockNums
		} else {
			for _, pf := range obj.MergedPFVersion.MergedPFVersionToList() {
				obj.Tokens[token][pf] = blockNums
			}
		}
	}
	obj.Reduntant = nil
	obj.MergedPFVersion = nil
	// log.Info(utils.ToJson(obj))
}
func (obj *DetailsDS) Save() core.Json {
	data, err := json.Marshal(obj)
	log.CheckFatal(err)
	out, err := utils.ReadJsonReader(bytes.NewBuffer(data))
	log.CheckFatal(err)
	return core.Json(out)
}
