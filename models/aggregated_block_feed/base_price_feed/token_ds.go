package base_price_feed

import (
	"bytes"
	"encoding/json"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type DetailsDS struct {
	PFType      string                      `json:"pfType"`
	Underlyings []string                    `json:"underlyings"` // for pyth it is id, and for singleasset curve it is underlyingFeeds
	Info        map[string]*core.RedStonePF `json:"info"`
	FetchedInfo bool                        `json:"FetchedInfo,omitempty"`
}

func (obj *DetailsDS) Load(in core.Json, version core.VersionType) {
	data, err := json.Marshal(in)
	log.CheckFatal(err)
	err = utils.ReadJsonReaderAndSetInterface(bytes.NewBuffer(data), obj)
	log.CheckFatal(err)

	//
	a := in["info"]
	redstoneMap := map[string]*core.RedStonePF{}
	if a != nil {
		str := utils.ToJson(a)
		err := utils.ReadJsonReaderAndSetInterface(bytes.NewBufferString(str), &redstoneMap)
		log.CheckFatal(err)
	}
	obj.Info = redstoneMap
}
func (obj *DetailsDS) Save() core.Json {
	data, err := json.Marshal(obj)
	log.CheckFatal(err)
	out, err := utils.ReadJsonReader(bytes.NewBuffer(data))
	log.CheckFatal(err)
	return core.Json(out)
}
