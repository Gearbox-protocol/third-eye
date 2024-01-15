package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

var _SCALE = utils.GetExpInt(18)

type Farmv3 struct {
	Pool        string `gorm:"column:pool"`
	Farm        string `gorm:"column:farm;primaryKey"`
	DieselToken string `gorm:"column:diesel_token"`
	//
	Checkpoint uint64       `gorm:"column:checkpoint"`
	Fpt        *core.BigInt `gorm:"column:farmed_per_token"`
	//
	Reward *core.BigInt `gorm:"column:reward"`
	Period uint64       `gorm:"column:period"`
	EndTs  uint64       `gorm:"column:end_ts"`
	//
	TotalSupply *core.BigInt `gorm:"column:total_supply"`
}

func (Farmv3) TableName() string {
	return "farm_v3"
}

// farmAccounting
func (farm *Farmv3) startFarming(reward *big.Int, newPeriod, currentTs uint64) {
	lastEndTs := farm.EndTs
	if lastEndTs > currentTs {
		finishedRewards := new(big.Int).Quo(farm.farmedSinceCheckpointScaled(currentTs), _SCALE)
		remainingFunds := new(big.Int).Sub(farm.Reward.Convert(), finishedRewards)
		reward = new(big.Int).Add(reward, remainingFunds)
	}
	//
	farm.EndTs = currentTs + newPeriod
	farm.Period = newPeriod
	farm.Reward = (*core.BigInt)(reward)
}
func (farm *Farmv3) stopFarming(reward *big.Int, currentTs uint64) {
	farm.EndTs = currentTs
	farm.Period = 0
	farm.Reward = (*core.BigInt)(new(big.Int))
}

func (farm *Farmv3) farmedSinceCheckpointScaled(currentTs uint64) *big.Int {
	if farm.Period == 0 {
		return big.NewInt(0)
	}
	elapsed := utils.Min(currentTs, farm.EndTs) - (farm.EndTs - farm.Period)
	num := new(big.Int).Mul(
		new(big.Int).Mul(big.NewInt(int64(elapsed)), farm.Reward.Convert()),
		_SCALE,
	)
	return new(big.Int).Quo(
		num,
		big.NewInt(int64(farm.Period)),
	)
}

// userAccounting
func (farm *Farmv3) calcFarmedPerToken(currentTs uint64) *big.Int {
	fpt := farm.Fpt.Convert()
	if farm.TotalSupply.Convert().Sign() != 0 {
		_fpt := new(big.Int).Quo(
			farm.farmedSinceCheckpointScaled(currentTs),
			farm.TotalSupply.Convert(),
		)
		fpt = new(big.Int).Add(fpt, _fpt)
	}
	log.Info(farm.TotalSupply.Convert(), "total_supply")
	return fpt
}
