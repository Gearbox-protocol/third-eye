package gear_token

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (mdl *GearToken) GetUnderlyingState() interface{} {
	arr := mdl.arrayOfGearBalanceUpdates
	mdl.arrayOfGearBalanceUpdates = []*core.GearBalance{}
	return arr
}

func (mdl *GearToken) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStatePresent = true
	gb, ok := obj.([]*core.GearBalance)
	if !ok {
		log.Fatal("Type assertion for gear token state failed")
	}
	state := map[string]*core.GearBalance{}
	for _, entry := range gb {
		state[entry.User] = entry
	}
	mdl.State = state
}

func (mdl *GearToken) HasUnderlyingState() bool {
	gb := []*core.GearBalance{}
	for _, entry := range mdl.State {
		if entry.Updated {
			gb = append(gb, entry)
			entry.Updated = false
		}
	}
	mdl.arrayOfGearBalanceUpdates = gb
	return len(gb) > 0
}
