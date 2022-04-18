package gear_token

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

func (mdl *GearToken) GetUnderlyingState() interface{} {
	arr := mdl.arrayOfGearBalanceUpdates
	mdl.arrayOfGearBalanceUpdates = []*schemas.GearBalance{}
	return arr
}

func (mdl *GearToken) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStatePresent = true
	gb, ok := obj.([]*schemas.GearBalance)
	if !ok {
		log.Fatal("Type assertion for gear token state failed")
	}
	state := map[string]*schemas.GearBalance{}
	for _, entry := range gb {
		state[entry.User] = entry
	}
	mdl.State = state
}

func (mdl *GearToken) HasUnderlyingState() bool {
	gb := []*schemas.GearBalance{}
	for _, entry := range mdl.State {
		if entry.Updated {
			gb = append(gb, entry)
			entry.Updated = false
		}
	}
	mdl.arrayOfGearBalanceUpdates = gb
	return len(gb) > 0
}
