package pool_v3

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type Zappers []Zapper

type Zapper struct {
	// RewardToken string
	Farm string
	// Pool        string
	Zapper  string
	TokenIn string
}

func (d *Zappers) Load(in core.Json) {
	if in != nil && in["zappers"] == nil {
		return
	}
	str := utils.ToJson(in["zappers"])
	data := Zappers{}
	err := utils.SetJson([]byte(str), &data)
	log.CheckFatal(err)
	*d = data
}

func (d Zappers) Save(in *core.Json) {
	delete(*in, "USDC-farmedUSDCv3")
	delete(*in, "ETH-farmedETHv3")
	delete(*in, "farmedUSDCv3")
	(*in)["zappers"] = d
}

func (d Zappers) GetFarm() []string {
	var res []string
	for _, x := range d {
		res = append(res, x.Farm)
	}
	return res
}

func (d Zappers) GetZapper() []string {
	var res []string
	for _, x := range d {
		res = append(res, x.Zapper)
	}
	return res
}
